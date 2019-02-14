package test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

func TestCreateEvent(t *testing.T) {

	resp, err := http.Post("http://206.189.133.125/api/v1/event/create-event", "application/json",
		bytes.NewBuffer([]byte(`
			{"event":{
			 "clubName":"GDG",
			 "name":"DEVFEST 2019",
			 "toDate":"10TH OCTOBER",
			 "fromDate":"8TH OCTOBER",
			 "toTime":"10 PM",
			 "fromTime":"11 AM",
			 "budget":"200000",
			 "description":"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING",
			 "category":"TECHNICAL",
			 "venue":"ANNA AUDI",
			 "attendance":"4000",
			 "expectedParticipants":"4000",
			 "facultyCoordinator":{
			    "name":"NOORU MAA",
			    "registrationNumber":"17BBE1010",
			    "email":"SDADAS@A.COM",
			    "phoneNumber":"919191991911",
			    "gender":"M",
			    "eventsAttended":"ALL"
			 },
			 "studentCoordinator":{
			    "name":"NOOR",
			    "registrationNumber":"17BBE1010",
			    "email":"SDADAS@A.COM",
			    "phoneNumber":"919191991911",
			    "gender":"M",
			    "eventsAttended":"ALL"
			 },
			 "guest":{
			    "name":"DAAS",
			    "email":"ASDSAD#ASD.COM",
			    "phoneNumber":"11111111111",
			    "gender":"F",
			    "stake":"SOME MONAYYYY",
			    "locationOfStay":"GHAR"
			 },
			 "PROrequest":"SAJDOOSIJANDFSAKFDSAFD",
			 "campusEngineerRequest":"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD",
			 "duration":"16 hours",
			 "mainSponsor":{
			    "name":"DAASA",
			    "email":"ASDSAD#ASD.COM",
			    "phoneNumber":"11111111111",
			    "gender":"F",
			    "stake":"SOME MONAYYYY",
			    "locationOfStay":"GHAR2"
			 }
			}}`)))

	if err != nil {
		t.Errorf("Error occurred while fetching request: %t", err)
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	log.Println(resp.StatusCode)

	if err != nil {
		t.Errorf("Error unmarshaling JSON response %t", err)
	}

	if result["err"] != nil {
		t.Errorf("Error creating event: %t", result["err"])
	}

	if result["error"] != nil {
		t.Errorf("Error creating event: %t", result["error"])
	}

	log.Printf("Response from server: %v", result["rs"])

}

func TestReadEvent(t *testing.T) {

	if resp, err := http.Post("http://206.189.133.125/api/v1/event/read-event", "application/json",
		bytes.NewBuffer([]byte(`{
			"query": {
				"key": "clubName",
				"value": "GDG",
				"specific": "DEVFEST 2019"
			}
		}`))); err != nil {
		t.Errorf("Error sending request, %t", err)
	} else {
		var result map[string]interface{}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			t.Errorf("Error unmarshaling JSON response, %t", err)
		}

		if result["err"] != nil {
			t.Errorf("Error from server: %t", result["err"])
		}
	}

}
