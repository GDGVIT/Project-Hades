package model

import (
	"fmt"
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
		campusEngineerRequest:$campusEngineerRequest, duration:$duration, status:$status}) 
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
		"status":                "true",
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

func ShowEventData(q Query, c chan EventReturn) {
	data, _, _, err := con.QueryNeoAll(`
	MATCH (n:EVENT)-[:StudentCoordinator]->(a)
	MATCH (n:EVENT)-[:FacultyCoordinator]->(b)
	MATCH (n:EVENT)-[:GUEST]->(c)
	WHERE n.`+q.Key+`=$val 
	RETURN n.clubName, n.name, n.toDate, n.fromDate, n.toTime, n.fromTime, n.budget, n.description, n.category,
	n.venue, n.attendance, n.expectedParticipants, n.PROrequest, n.campusEngineerRequest, n.duration, a.name, 
	a.registrationNumber, a.email, a.phoneNumber, a.gender, b.name, b.registrationNumber, b.email, 
	b.phoneNumber, b.gender, c.name, c.email, c.phoneNumber, c.gender, c.stake, c.locationOfStay, n.status
	`, map[string]interface{}{
		"val": q.Value,
	})

	if err != nil {
		c <- EventReturn{Event{}, err}
		return
	}

	var ev Event

	if len(data) < 1 {
		c <- EventReturn{ev, fmt.Errorf("No Event found")}
		return
	}

	ev = Event{
		ClubName:              data[0][0].(string),
		Name:                  data[0][1].(string),
		ToDate:                data[0][2].(string),
		FromDate:              data[0][3].(string),
		ToTime:                data[0][4].(string),
		FromTime:              data[0][5].(string),
		Budget:                data[0][6].(string),
		Description:           data[0][7].(string),
		Category:              data[0][8].(string),
		Venue:                 data[0][9].(string),
		Attendance:            data[0][10].(string),
		ExpectedParticipants:  data[0][11].(string),
		PROrequest:            data[0][12].(string),
		CampusEngineerRequest: data[0][13].(string),
		Duration:              data[0][14].(string),
		StudentCoordinator: Participant{
			data[0][15].(string),
			data[0][16].(string),
			data[0][17].(string),
			data[0][18].(string),
			data[0][19].(string),
		},
		FacultyCoordinator: Participant{
			data[0][20].(string),
			data[0][21].(string),
			data[0][22].(string),
			data[0][23].(string),
			data[0][24].(string),
		},
		GuestDetails: Guest{
			data[0][25].(string),
			data[0][26].(string),
			data[0][27].(string),
			data[0][28].(string),
			data[0][29].(string),
			data[0][30].(string),
		},
		Status: data[0][31].(string),
	}

	c <- EventReturn{ev, nil}

	return
}

// delete event with given query
func DeleteEvent(q Query, c chan error) {
	result, err := con.ExecNeo(`
		MATCH(n:EVENT)-[r]->(a)
		WHERE n.`+q.Key+`=$val
		DETACH DELETE n, a
	`, map[string]interface{}{
		"val": q.Value,
	})
	if err != nil {
		c <- err
	}

	log.Println(result)
	log.Println("Event deleted")
	c <- nil
	return
}

// update event with given query and new value
func UpdateEvent(q Query, c chan error) {
	result, err := con.ExecNeo(`
		MATCH(n:EVENT)
		WHERE n.`+q.Key+`=$val
		SET n.`+q.ChangeKey+`=$val1
		RETURN n.`+q.ChangeKey+`
	`, map[string]interface{}{
		"val":  q.Value,
		"val1": q.ChangeValue,
	})

	if err != nil {
		c <- err
		return
	}
	c <- nil

	log.Println(result)
	log.Printf("Updated")

}

// create a new node with given label and participant data struct (FOR COORDINATORS)
func CreateParticipant(e Event, label string, c chan error, mutex *sync.Mutex) {
	if e.GetField(label, "Email") == "" {
		c <- nil
		return
	}
	mutex.Lock()
	result, err := con.ExecNeo(`MATCH(a:EVENT) WHERE a.name=$EventName
	CREATE (n:INCHARGE {name:$name, registrationNumber:$registrationNumber,
		email:$email, phoneNumber:$phoneNumber, gender: $gender})<-[:`+label+`]-(a) `, map[string]interface{}{
		"EventName":          e.Name,
		"name":               e.GetField(label, "Name"),
		"registrationNumber": e.GetField(label, "RegistrationNumber"),
		"email":              e.GetField(label, "Email"),
		"phoneNumber":        e.GetField(label, "PhoneNumber"),
		"gender":             e.GetField(label, "Gender"),
	})
	if err != nil {
		c <- err
		return
	}
	mutex.Unlock()
	log.Println(result)
	log.Printf("Created %s node", label)
	c <- nil
	return
}

// create a new guest node with relationship to the event
func CreateGuest(e Event, c chan error, mutex *sync.Mutex) {
	if e.GuestDetails.Email == "" {
		c <- nil
		return
	}
	mutex.Lock()
	result, err := con.ExecNeo(`MATCH(a:EVENT) WHERE a.name=$EventName
	CREATE (n:GUEST {name:$name, stake:$stake,
	email:$email, phoneNumber:$phoneNumber, gender: $gender, locationOfStay:$locationOfStay
	})<-[:GUEST]-(a) `, map[string]interface{}{
		"EventName":      e.Name,
		"name":           e.GetField("GuestDetails", "Name"),
		"stake":          e.GetField("GuestDetails", "Stake"),
		"email":          e.GetField("GuestDetails", "Email"),
		"phoneNumber":    e.GetField("GuestDetails", "PhoneNumber"),
		"gender":         e.GetField("GuestDetails", "Gender"),
		"locationOfStay": e.GetField("GuestDetails", "LocationOfStay"),
	})
	if err != nil {
		c <- err
		return
	}
	mutex.Unlock()

	log.Println(result)
	log.Println("Created GUEST node")
	c <- nil
	return
}
