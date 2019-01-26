package model

import (
	"log"
	"sync"
)

func CreateEvent(e Event, ce chan error) {
	c := make(chan error)
	//go createParticipant(e, "StudentCoordinator", c)
	result, err := con.ExecNeo(`CREATE (n:EVENT {name:$name, clubName:$clubName, toDate:$toDate, 
		fromDate: $fromDate, toTime:$toTime, fromTime:$fromTime, budget:$budget, 
		description:$description, category:$category, venue:$venue, attendance:$attendance, 
		expectedParticipants:$expectedParticipants, PROrequest:$PROrequest, 
		campusEngineerRequest:$campusEngineerRequest, duration:$duration}) 
		RETURN n.name`, map[string]interface{}{

		"name":                  e.Name,
		"clubName":              e.ClubName,
		"toDate":                e.ToDate,
		"fromDate":              e.FromDate,
		"toTime":                e.ToTime,
		"fromTime":              e.FromTime,
		"budget":                e.Budget,
		"description":           e.Description,
		"category":              e.Category,
		"venue":                 e.Venue,
		"PROrequest":            e.PROrequest,
		"campusEngineerRequest": e.CampusEngineerRequest,
		"duration":              e.Duration,
		"attendance":            e.Attendance,
		"expectedParticipants":  e.ExpectedParticipants,
	})
	if err != nil {
		ce <- err
		return
	}
	log.Println(result)

	// CREATE STUDENT COORDINATOR, FACULTY COORDINATOR, SPONSOR AND GUEST NODES
	var mutex = &sync.Mutex{}
	go CreateParticipant(e, "StudentCoordinator", c, mutex)
	go CreateParticipant(e, "FacultyCoordinator", c, mutex)
	go CreateParticipant(e, "MainSponsor", c, mutex)
	go CreateGuest(e, c, mutex)

	err1, err2, err3, err4 := <-c, <-c, <-c, <-c

	switch {
	case err1 != nil:
		ce <- err1
		return
	case err2 != nil:
		ce <- err2
		return
	case err3 != nil:
		ce <- err3
		return
	case err4 != nil:
		ce <- err4
		return
	}

	log.Println("Created Event node")
	ce <- nil
	return
}
