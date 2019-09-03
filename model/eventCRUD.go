package model

import (
	"fmt"
	"log"
	"sync"
)

func CreateEvent(e Event, ce chan error) {
	c := make(chan error)

	// check if event with same name exists
	data, _, _, err := con.QueryNeoAll(
		`MATCH(n:EVENT) WHERE n.name=$name RETURN n.clubName`,
		map[string]interface{}{
			"name": e.Name,
		})
	str := fmt.Sprintf("%v", data)
	if str != "[[]]" && str != "[]" {
		ce <- fmt.Errorf("An event with this name already exists")
		return
	}

	result, err := con.ExecNeo(`
MATCH (b:ORG)
WHERE b.name = $clubName
CREATE (n:EVENT {name:$name, clubName:$clubName, toDate:$toDate, 
		fromDate: $fromDate, toTime:$toTime, fromTime:$fromTime, budget:$budget, 
		description:$description, category:$category, venue:$venue, attendance:$attendance, 
		expectedParticipants:$expectedParticipants, PROrequest:$PROrequest, 
		campusEngineerRequest:$campusEngineerRequest, duration:$duration, status:$status})-[:EVENT]->(b) 
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
	// go CreateParticipant(e, "MainSponsor", c, mutex)
	// go CreateGuest(e, c, mutex)

	err1, err2 := <-c, <-c

	switch {
	case err1 != nil:
		ce <- err1
		return
	case err2 != nil:
		ce <- err2
		return
	}

	log.Println("Created Event node")
	ce <- nil
	return
}

func ShowEventData(q Query, c chan EventReturn) {
	var (
		ev  []Event
		str string
	)

	if q.Specific == "" {
		str = `
		MATCH (n:EVENT)-[:StudentCoordinator]->(a)
		MATCH (n:EVENT)-[:FacultyCoordinator]->(b)
		WHERE n.` + q.Key + `=$val 
		RETURN n.clubName, n.name, n.toDate, n.fromDate, n.toTime, n.fromTime, n.budget, n.description, n.category,
		n.venue, n.attendance, n.expectedParticipants, n.PROrequest, n.campusEngineerRequest, n.duration, a.name, 
		a.registrationNumber, a.email, a.phoneNumber, a.gender, b.name, b.registrationNumber, b.email, 
		b.phoneNumber, b.gender, n.status
		`
	} else {
		str = `
		MATCH (n:EVENT)-[:StudentCoordinator]->(a)
		MATCH (n:EVENT)-[:FacultyCoordinator]->(b)
		WHERE n.` + q.Key + `=$val  AND n.name=$specific
		RETURN n.clubName, n.name, n.toDate, n.fromDate, n.toTime, n.fromTime, n.budget, n.description, n.category,
		n.venue, n.attendance, n.expectedParticipants, n.PROrequest, n.campusEngineerRequest, n.duration, a.name, 
		a.registrationNumber, a.email, a.phoneNumber, a.gender, b.name, b.registrationNumber, b.email, 
		b.phoneNumber, b.gender, n.status
		`
	}

	data, _, _, err := con.QueryNeoAll(str, map[string]interface{}{
		"val":      q.Value,
		"specific": q.Specific,
	})

	if err != nil {
		c <- EventReturn{[]Event{}, err}
		return
	}

	fmt.Println(data)
	if len(data) < 1 {
		c <- EventReturn{ev, fmt.Errorf("No Event found")}
		return
	}

	for i, _ := range data {
		ev = append(ev, Event{
			ClubName:              data[i][i].(string),
			Name:                  data[i][1].(string),
			ToDate:                data[i][2].(string),
			FromDate:              data[i][3].(string),
			ToTime:                data[i][4].(string),
			FromTime:              data[i][5].(string),
			Budget:                data[i][6].(string),
			Description:           data[i][7].(string),
			Category:              data[i][8].(string),
			Venue:                 data[i][9].(string),
			Attendance:            data[i][10].(string),
			ExpectedParticipants:  data[i][11].(string),
			PROrequest:            data[i][12].(string),
			CampusEngineerRequest: data[i][13].(string),
			Duration:              data[i][14].(string),
			StudentCoordinator: Participant{
				data[i][15].(string),
				data[i][16].(string),
				data[i][17].(string),
				data[i][18].(string),
				data[i][19].(string),
			},
			FacultyCoordinator: Participant{
				data[i][20].(string),
				data[i][21].(string),
				data[i][22].(string),
				data[i][23].(string),
				data[i][24].(string),
			},
			Status: data[i][25].(string),
		})
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

	mutex.Lock()
	_, err := con.ExecNeo(`MATCH(a:EVENT) WHERE a.name=$EventName
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

	log.Printf("Created %s node", label)
	c <- nil
	return
}

func ReadEventsByOrg(org string) (events []Event, err error) {
	data, _, _, err := con.QueryNeoAll(`
						OPTIONAL MATCH (n:EVENT)-[:EVENT]->(o:ORG)
						WHERE o.name = $org
						RETURN n.clubName, n.name, n.toDate, n.fromDate, n.toTime, n.fromTime, n.budget, n.description, n.category,
		n.venue, n.attendance, n.expectedParticipants, n.PROrequest, n.campusEngineerRequest, n.duration
				`, map[string]interface{}{
		"org": org,
	})
	if err != nil {
		return nil, err
	}

	fmt.Println(data)
	fmt.Println(len(data))
	if len(data) < 1 {
		return events, nil
	}

	for i, _ := range data {
		if data[i][0] == nil {
			return events, nil
		}
		events = append(events, Event{
			ClubName:              data[i][0].(string),
			Name:                  data[i][1].(string),
			ToDate:                data[i][2].(string),
			FromDate:              data[i][3].(string),
			ToTime:                data[i][4].(string),
			FromTime:              data[i][5].(string),
			Budget:                data[i][6].(string),
			Description:           data[i][7].(string),
			Category:              data[i][8].(string),
			Venue:                 data[i][9].(string),
			Attendance:            data[i][10].(string),
			ExpectedParticipants:  data[i][11].(string),
			PROrequest:            data[i][12].(string),
			CampusEngineerRequest: data[i][13].(string),
			Duration:              data[i][14].(string),
		})
	}

	return events, nil
}
