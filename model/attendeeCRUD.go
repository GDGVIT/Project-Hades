package model

import (
	"fmt"
	"log"
	"sync"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func CreateAttendee(eventName string, p Participant, c chan error, mutex *sync.Mutex, conn bolt.Conn) {

	mutex.Lock()
	_, err := conn.ExecNeo(`MATCH(a:EVENT) WHERE a.name=$EventName
	CREATE (n:ATTENDEE {name:$name, registrationNumber:$registrationNumber,
		email:$email, phoneNumber:$phoneNumber, gender: $gender})<-[:ATTENDS]-(a) `, map[string]interface{}{
		"EventName":          eventName,
		"name":               p.Name,
		"registrationNumber": p.RegistrationNumber,
		"email":              p.Email,
		"phoneNumber":        p.PhoneNumber,
		"gender":             p.Gender,
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

func ReadAttendee(q Query, c chan ParticipantReturn, mutex *sync.Mutex, conn bolt.Conn) {

	mutex.Lock()
	data, _, _, err := conn.QueryNeoAll(`MATCH(a:ATTENDEE) WHERE a.`+q.Key+`=$val
	RETURN a.name, a.registrationNumber,a.email, a.phoneNumber, a.gender`, map[string]interface{}{
		"val": q.Value,
	})

	mutex.Unlock()

	var pt []Participant

	if err != nil {
		c <- ParticipantReturn{pt, err}
		return
	}

	if len(data) < 1 {
		c <- ParticipantReturn{pt, fmt.Errorf("No attendee found")}
		return
	}

	for i, _ := range data {
		pt = append(pt, Participant{
			Name:               data[i][0].(string),
			RegistrationNumber: data[i][1].(string),
			Email:              data[i][2].(string),
			PhoneNumber:        data[i][3].(string),
			Gender:             data[i][4].(string),
		})
	}
	log.Printf("Found attendee node")
	c <- ParticipantReturn{pt, nil}

	return
}

// update Attendee with given query and new value
func UpdateAttendee(q Query, c chan error, conn bolt.Conn) {
	result, err := conn.ExecNeo(`
		MATCH(n:ATTENDEE)
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

// delete attendee with given query
func DeleteAttendee(q Query, c chan error, conn bolt.Conn) {
	result, err := conn.ExecNeo(`
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
