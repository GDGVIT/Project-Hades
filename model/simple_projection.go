package model

import (
	"fmt"
	"log"
)

func ViewAll(event string, query Query, c chan SafeParticipantReturn) {
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

	data, _, _, err := con.QueryNeoAll(str, map[string]interface{}{
		"ev": event,
	})

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
