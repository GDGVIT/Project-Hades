package projectionModule

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

func ViewAll(event string, query Query, c chan SafeParticipantReturn, mutex *sync.Mutex) {
	var str string
	if query.Key == "" {

		str = `MATCH (n:EVENT)-[:ATTENDS]->(a)
		WHERE n.name=$ev
		RETURN a.name, a.registrationNumber,a.email, a.phoneNumber, a.gender`

	} else {
		str = fmt.Sprintf(`MATCH (n:EVENT)-[:ATTENDS]->(a)
			WHERE n.name=$ev AND a.%s = "%s"
			RETURN a.name, a.registrationNumber,a.email, a.phoneNumber, a.gender`,
			query.Key,
			query.Value,
		)
	}

	mutex.Lock()
	data, _, _, err := con.QueryNeoAll(str, map[string]interface{}{
		"ev": event,
	})
	mutex.Unlock()

	var pt []Participant

	if err != nil {
		c <- SafeParticipantReturn{pt, err}
		return
	}

	if len(data) < 1 {
		c <- SafeParticipantReturn{pt, fmt.Errorf("No attendee found")}
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
	log.Printf("Found attendee nodes")
	c <- SafeParticipantReturn{pt, nil}

	return

}

func ViewPresent(event string, query Query, day int, c chan SafeParticipantReturn, mutex *sync.Mutex) {
	var str string

	if query.Key == "" {

		str = `MATCH (n:EVENT)-[:PRESENT` + strconv.Itoa(day) + `]->(a)
			WHERE n.name=$ev
			RETURN a.name, a.registrationNumber,a.email, a.phoneNumber, a.gender`

	} else {
		str = fmt.Sprintf(`MATCH (n:EVENT)-[:PRESENT%d]->(a)
				WHERE n.name=$ev AND a.%s = "%s"
				RETURN a.name, a.registrationNumber,a.email, a.phoneNumber, a.gender`,
			day,
			query.Key,
			query.Value,
		)
	}

	mutex.Lock()

	data, _, _, err := con.QueryNeoAll(str, map[string]interface{}{
		"ev": event,
	})

	mutex.Unlock()

	var pt []Participant

	if err != nil {
		c <- SafeParticipantReturn{pt, err}
		return
	}

	if len(data) < 1 {
		c <- SafeParticipantReturn{pt, fmt.Errorf("No attendee found")}
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
	log.Printf("Found present nodes")
	c <- SafeParticipantReturn{pt, nil}

	return

}
func ViewAbsent(event string, query Query, day int, c chan SafeParticipantReturn) {

	var pt []Participant
	c1 := make(chan SafeParticipantReturn)
	c2 := make(chan SafeParticipantReturn)
	mutex := &sync.Mutex{}
	go ViewAll(event, query, c1, mutex)
	go ViewPresent(event, query, day, c2, mutex)

	ms1 := <-c1
	ms2 := <-c2

	if err1 := ms1.Err; err1 != nil {
		c <- SafeParticipantReturn{pt, err1}
		return
	}

	if err2 := ms2.Err; err2 != nil {
		c <- SafeParticipantReturn{pt, err2}
		return
	}

	l := len(ms1.Participants) - len(ms2.Participants)

	c <- SafeParticipantReturn{ms1.Participants[:l], nil}

}
