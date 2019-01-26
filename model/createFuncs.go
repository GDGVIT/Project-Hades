package model

import (
	"log"
	"sync"
)

// create a new node with given label and participant data struct
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
