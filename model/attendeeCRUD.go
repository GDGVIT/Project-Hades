package model

import (
	"fmt"
	"log"
	"sync"
)

func CreateAttendee(eventName string, p Participant, c chan error, mutex *sync.Mutex) {

	mutex.Lock()
	_, err := con.ExecNeo(`MATCH(a:EVENT) WHERE a.name=$EventName
	CREATE (n:ATTENDEE {name:$name, registrationNumber:$registrationNumber,
		email:$email, phoneNumber:$phoneNumber, gender: $gender, attended:$attended})<-[:ATTENDS]-(a) `, map[string]interface{}{
		"EventName":          eventName,
		"name":               p.Name,
		"registrationNumber": p.RegistrationNumber,
		"email":              p.Email,
		"phoneNumber":        p.PhoneNumber,
		"gender":             p.Gender,
		"attended":           "absent",
	})
	if err != nil {
		c <- err
		return
	}
	mutex.Unlock()
	log.Printf("Created attendee node")
	c <- nil
	return
}

func UpdateAttendee(q Query, c chan error) {

	result, err := con.ExecNeo(`
		MATCH(n:ATTENDEE)
		WHERE n.`+q.Key+`=$val
		SET n.`+q.ChangeKey+`=$val1
		RETURN n.`+q.ChangeKey+`
	`, map[string]interface{}{
		"val":  q.Value,
		"val1": q.ChangeKey,
	})

	if err != nil {
		c <- err
		return
	}
	c <- nil

	log.Println(result)
	log.Printf("Updated")

}

func DeleteAttendee(q Query, c chan error) {
	result, err := con.ExecNeo(`
		MATCH(n:ATTENDEE)<-[r]-(a)
		WHERE n.`+q.Key+`=$val
		DETACH DELETE n
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

func ReadAttendee(q Query, c chan ParticipantReturn, mutex *sync.Mutex) {

	mutex.Lock()
	data, _, _, err := con.QueryNeoAll(`MATCH(a:ATTENDEE) WHERE a.`+q.Key+`=$val
	RETURN a.name, a.registrationNumber,a.email, a.phoneNumber, a.gender, a.attended`, map[string]interface{}{
		"val": q.Value,
	})

	mutex.Unlock()

	var pt []Attendee

	if err != nil {
		c <- ParticipantReturn{pt, err}
		return
	}

	if len(data) < 1 {
		c <- ParticipantReturn{pt, fmt.Errorf("No attendee found")}
		return
	}

	for i, _ := range data {
		pt = append(pt, Attendee{
			Name:               data[i][0].(string),
			RegistrationNumber: data[i][1].(string),
			Email:              data[i][2].(string),
			PhoneNumber:        data[i][3].(string),
			Gender:             data[i][4].(string),
			Attended:           data[i][5].(string),
		})
	}
	log.Printf("Found attendee node")
	c <- ParticipantReturn{pt, nil}

	return
}