package model

import (
	"fmt"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func ShowEventData(q Query, c chan EventReturn, conn bolt.Conn) {
	data, _, _, err := conn.QueryNeoAll(`
	MATCH (n:EVENT)-[:StudentCoordinator]->(a)
	MATCH (n:EVENT)-[:FacultyCoordinator]->(b)
	MATCH (n:EVENT)-[:GUEST]->(c)
	WHERE n.`+q.Key+`=$val 
	RETURN n.clubName, n.name, n.toDate, n.fromDate, n.toTime, n.fromTime, n.budget, n.description, n.category,
	n.venue, n.attendance, n.expectedParticipants, n.PROrequest, n.campusEngineerRequest, n.duration, a.name, 
	a.registrationNumber, a.email, a.phoneNumber, a.gender, b.name, b.registrationNumber, b.email, 
	b.phoneNumber, b.gender, c.name, c.email, c.phoneNumber, c.gender, c.stake, c.locationOfStay
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
	}

	c <- EventReturn{ev, nil}

	return
}
