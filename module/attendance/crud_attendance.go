package attendanceModule

import (
	"fmt"
	"log"
	"sync"
)

func CreateAttendee(eventName string, p Participant, c chan error, mutex *sync.Mutex) {

	// check if user exists
	data, _, _, err := con.QueryNeoAll(`
	MATCH(n:ATTENDEE)
	WHERE n.email=$rn
	RETURN n.email
		`, map[string]interface{}{
		"rn": p.Email,
	})

	if err != nil {
		c <- err
		return
	}

	// if not, rceate user
	if len(data) < 1 {
		mutex.Lock()

		rss, err := con.ExecNeo(`MATCH(a:EVENT) WHERE a.name=$EventName
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
		log.Println(rss)
		mutex.Unlock()
	} else {

		// if yes, check if relation exists

		data, _, _, err := con.QueryNeoAll(`
		MATCH(n:EVENT)-[r:ATTENDS]->(b)
		WHERE b.email=$rn and n.name=$ev
		RETURN r
			`, map[string]interface{}{
			"rn": p.Email,
			"ev": eventName,
		})

		rel := fmt.Sprintf("%v", data)
		log.Println("HAGGA\n\n\n\n")
		log.Println(rel)
		if err != nil {
			c <- err
			return
		}

		if rel == "[]" || rel == "" {

			// if doesnt exist then create relation
			mutex.Lock()

			rss, err := con.ExecNeo(`MATCH(a:EVENT) WHERE a.name=$EventName
			MATCH(b:ATTENDEE) WHERE b.email=$rn
			CREATE (b)<-[:ATTENDS]-(a) `, map[string]interface{}{
				"EventName": eventName,
				"rn":        p.Email,
			})
			if err != nil {
				c <- err
				return
			}
			log.Println(rss)

			mutex.Unlock()
		} else {
			c <- fmt.Errorf("User has already registered")
			return
		}
	}

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
		})
	}
	log.Printf("Found attendee node")
	c <- ParticipantReturn{pt, nil}

	return
}

func RmAttendee(q Query, c chan error) {
	result, err := con.ExecNeo(`
		MATCH(n:ATTENDEE)<-[r:ATTENDS]-(a:EVENT)
		WHERE n.`+q.Key+`=$val AND a.clubName=$club AND a.name=$event
		DELETE r
	`, map[string]interface{}{
		"val":   q.Value,
		"club":  q.ChangeKey,
		"event": q.ChangeValue,
	})
	if err != nil {
		c <- err
	}

	log.Println(result)

	// TDO check if zombie attendee,, and remove them

	log.Println("Relation to event removed")
	c <- nil
	return
}
