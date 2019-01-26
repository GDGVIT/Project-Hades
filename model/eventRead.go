package model

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"log"
)

type Query struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	ChangeKey   string `json:"changeKey"`
	ChangeValue string `json:"changeValue"`
}

type EventReturn struct {
	Event Event
	Err   error
}

func ShowEventData(q Query, c chan EventReturn, conn bolt.Conn) {
	result, err := conn.ExecNeo(`
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
	// for result.Next() {
	// 	log.Println(result.Record())
	// 	ev = Event{
	// 		ClubName:              result.Record().GetByIndex(0).(string),
	// 		Name:                  result.Record().GetByIndex(1).(string),
	// 		ToDate:                result.Record().GetByIndex(2).(string),
	// 		FromDate:              result.Record().GetByIndex(3).(string),
	// 		ToTime:                result.Record().GetByIndex(4).(string),
	// 		FromTime:              result.Record().GetByIndex(5).(string),
	// 		Budget:                result.Record().GetByIndex(6).(string),
	// 		Description:           result.Record().GetByIndex(7).(string),
	// 		Category:              result.Record().GetByIndex(8).(string),
	// 		Venue:                 result.Record().GetByIndex(9).(string),
	// 		Attendance:            result.Record().GetByIndex(10).(string),
	// 		ExpectedParticipants:  result.Record().GetByIndex(11).(string),
	// 		PROrequest:            result.Record().GetByIndex(12).(string),
	// 		CampusEngineerRequest: result.Record().GetByIndex(13).(string),
	// 		Duration:              result.Record().GetByIndex(14).(string),
	// 		StudentCoordinator: Participant{
	// 			result.Record().GetByIndex(15).(string),
	// 			result.Record().GetByIndex(16).(string),
	// 			result.Record().GetByIndex(17).(string),
	// 			result.Record().GetByIndex(18).(string),
	// 			result.Record().GetByIndex(19).(string),
	// 		},
	// 		FacultyCoordinator: Participant{
	// 			result.Record().GetByIndex(20).(string),
	// 			result.Record().GetByIndex(21).(string),
	// 			result.Record().GetByIndex(22).(string),
	// 			result.Record().GetByIndex(23).(string),
	// 			result.Record().GetByIndex(24).(string),
	// 		},
	// 		GuestDetails: Guest{
	// 			result.Record().GetByIndex(25).(string),
	// 			result.Record().GetByIndex(26).(string),
	// 			result.Record().GetByIndex(27).(string),
	// 			result.Record().GetByIndex(28).(string),
	// 			result.Record().GetByIndex(29).(string),
	// 			result.Record().GetByIndex(30).(string),
	// 		},
	// 	}
	// }

	log.Println(result)
	c <- EventReturn{ev, nil}
	return
}
