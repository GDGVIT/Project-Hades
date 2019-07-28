package guestModule

import (
	"fmt"
	"log"
	"sync"
)

// func DeleteGuest(q Query, c chan error) {
// 	result, err := con.ExecNeo(`
// 		MATCH(n:Guest)<-[r]-(a)
// 		WHERE n.`+q.Key+`=$val
// 		DETACH DELETE n
// 	`, map[string]interface{}{
// 		"val": q.Value,
// 	})
// 	if err != nil {
// 		c <- err
// 	}

// 	log.Println(result)
// 	log.Println("Event deleted")
// 	c <- nil
// 	return
// }

func ReadGuest(q Query, c chan GuestReturn) {

	data, _, _, err := con.QueryNeoAll(`MATCH(n:EVENT)-[:GUEST]->(a:GUEST) WHERE n.`+q.Key+`=$val
	RETURN a.name,a.email, a.phoneNumber, a.gender, a.stake, a.locationOfStay`, map[string]interface{}{
		"val": q.Value,
	})

	var pt []Guest

	if err != nil {
		c <- GuestReturn{pt, err}
		return
	}

	if len(data) < 1 {
		c <- GuestReturn{pt, fmt.Errorf("No Guest found")}
		return
	}

	for i, _ := range data {
		pt = append(pt, Guest{
			Name:           data[i][0].(string),
			Email:          data[i][1].(string),
			PhoneNumber:    data[i][2].(string),
			Gender:         data[i][3].(string),
			Stake:          data[i][4].(string),
			LocationOfStay: data[i][5].(string),
		})
	}
	log.Printf("Found Guest node")
	c <- GuestReturn{pt, nil}

	return
}

func RmGuest(q Query, c chan error) {
	result, err := con.ExecNeo(`
		MATCH(n:EVENT)-[r:GUEST]->(a:GUEST)
		WHERE n.`+q.Key+`=$val AND a.name=$guestName
		DETACH DELETE a
	`, map[string]interface{}{
		"val":       q.Value,
		"guestName": q.ChangeKey,
	})
	if err != nil {
		c <- err
	}

	log.Println(result)

	// TDO check if zombie Guest, and remove them

	log.Println("Relation to event removed")
	c <- nil
	return
}

// create a new guest node with relationship to the event
func CreateGuest(event string, g Guest, c chan error, mutex *sync.Mutex) {

	mutex.Lock()
	_, err := con.ExecNeo(`MATCH(a:EVENT) WHERE a.name=$EventName
	CREATE (n:GUEST {name:$name, stake:$stake,
	email:$email, phoneNumber:$phoneNumber, gender: $gender, locationOfStay:$locationOfStay
	})<-[:GUEST]-(a) `, map[string]interface{}{
		"EventName":      event,
		"name":           g.Name,
		"stake":          g.Stake,
		"email":          g.Email,
		"phoneNumber":    g.PhoneNumber,
		"gender":         g.Gender,
		"locationOfStay": g.LocationOfStay,
	})
	if err != nil {
		c <- err
		return
	}
	mutex.Unlock()
	log.Println("Created GUEST node")
	c <- nil
	return
}
