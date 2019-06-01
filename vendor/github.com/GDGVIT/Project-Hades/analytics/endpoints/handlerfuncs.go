package endpoints

import (
	"encoding/json"
	"net/http"

	db "github.com/GDGVIT/Project-Hades/analytics/modelfuncs"
)

/**
*@api {get} /api/v1/analytics read specific logs
*@apiGroup analytics
*@apiName read specific logs
*@apiPermission admin
*@apiParam {string} subject event subject
*@apiParam {string} timestamp event timestamp
*@apiParam {String-Object} body event body
*
*@apiParamExample {json} request-url-example
*http://localhost:8085/api/v1/analytics?subject=hades.event.CreateEvent
*
*@apiParamExample {json} response-example
*{
*    "logs": [
*        [
*            {
*                "ID": 1,
*                "CreatedAt": "2019-02-21T20:45:21.579518607Z",
*                "UpdatedAt": "2019-02-21T20:45:21.579518607Z",
*                "DeletedAt": null,
*                "Subject": "hades.event.CreateEvent",
*                "Timestamp": "2019-02-21 20:45:21.579267685 +0000 UTC m=+943.125297422",
*                "Body": "{\"clubName\":\"GDG\",\"name\":\"DEVRELCONF\",\"toDate\":\"10TH OCTOBER\",\"fromDate\":\"8TH OCTOBER\",\"toTime\":\"10 PM\",\"fromTime\":\"11 AM\",\"budget\":\"200000\",\"description\":\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\",\"category\":\"TECHNICAL\",\"venue\":\"ANNA AUDI\",\"attendance\":\"4000\",\"expectedParticipants\":\"4000\",\"facultyCoordinator\":{\"name\":\"NOORU MAA\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"studentCoordinator\":{\"name\":\"NOOR\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"PROrequest\":\"SAJDOOSIJANDFSAKFDSAFD\",\"campusEngineerRequest\":\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\",\"duration\":\"16 hours\",\"status\":\"\"}"
*            },
*            {
*                "ID": 2,
*                "CreatedAt": "2019-02-21T20:47:28.002589821Z",
*                "UpdatedAt": "2019-02-21T20:47:28.002589821Z",
*                "DeletedAt": null,
*                "Subject": "hades.event.CreateEvent",
*                "Timestamp": "2019-02-21 20:47:28.002438401 +0000 UTC m=+1069.548468003",
*                "Body": "{\"clubName\":\"GDG\",\"name\":\"DEVRELCONF\",\"toDate\":\"10TH OCTOBER\",\"fromDate\":\"8TH OCTOBER\",\"toTime\":\"10 PM\",\"fromTime\":\"11 AM\",\"budget\":\"200000\",\"description\":\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\",\"category\":\"TECHNICAL\",\"venue\":\"ANNA AUDI\",\"attendance\":\"4000\",\"expectedParticipants\":\"4000\",\"facultyCoordinator\":{\"name\":\"NOORU MAA\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"studentCoordinator\":{\"name\":\"NOOR\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"PROrequest\":\"SAJDOOSIJANDFSAKFDSAFD\",\"campusEngineerRequest\":\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\",\"duration\":\"16 hours\",\"status\":\"\"}"
*            },
*            {
*                "ID": 3,
*                "CreatedAt": "2019-03-09T10:29:40.344850552Z",
*                "UpdatedAt": "2019-03-09T10:29:40.344850552Z",
*                "DeletedAt": null,
*                "Subject": "hades.event.CreateEvent",
*                "Timestamp": "2019-03-09 10:29:40.344694322 +0000 UTC m=+327.442605414",
*                "Body": "{\"clubName\":\"\",\"name\":\"\",\"toDate\":\"\",\"fromDate\":\"\",\"toTime\":\"\",\"fromTime\":\"\",\"budget\":\"\",\"description\":\"\",\"category\":\"\",\"venue\":\"\",\"attendance\":\"\",\"expectedParticipants\":\"\",\"facultyCoordinator\":{\"name\":\"\",\"registrationNumber\":\"\",\"email\":\"\",\"phoneNumber\":\"\",\"gender\":\"\"},\"studentCoordinator\":{\"name\":\"\",\"registrationNumber\":\"\",\"email\":\"\",\"phoneNumber\":\"\",\"gender\":\"\"},\"PROrequest\":\"\",\"campusEngineerRequest\":\"\",\"duration\":\"\",\"status\":\"\"}"
*            },
*            {
*                "ID": 4,
*                "CreatedAt": "2019-03-09T10:39:59.807360183Z",
*                "UpdatedAt": "2019-03-09T10:39:59.807360183Z",
*                "DeletedAt": null,
*                "Subject": "hades.event.CreateEvent",
*                "Timestamp": "2019-03-09 10:39:59.807139694 +0000 UTC m=+4.791159116",
*                "Body": "{\"clubName\":\"\",\"name\":\"\",\"toDate\":\"\",\"fromDate\":\"\",\"toTime\":\"\",\"fromTime\":\"\",\"budget\":\"\",\"description\":\"\",\"category\":\"\",\"venue\":\"\",\"attendance\":\"\",\"expectedParticipants\":\"\",\"facultyCoordinator\":{\"name\":\"\",\"registrationNumber\":\"\",\"email\":\"\",\"phoneNumber\":\"\",\"gender\":\"\"},\"studentCoordinator\":{\"name\":\"\",\"registrationNumber\":\"\",\"email\":\"\",\"phoneNumber\":\"\",\"gender\":\"\"},\"PROrequest\":\"\",\"campusEngineerRequest\":\"\",\"duration\":\"\",\"status\":\"\"}"
*            },
*            {
*                "ID": 5,
*                "CreatedAt": "2019-03-09T10:42:27.834845067Z",
*                "UpdatedAt": "2019-03-09T10:42:27.834845067Z",
*                "DeletedAt": null,
*                "Subject": "hades.event.CreateEvent",
*                "Timestamp": "2019-03-09 10:42:27.834629395 +0000 UTC m=+18.339562524",
*                "Body": "{\"clubName\":\"GDG\",\"name\":\"DEVFEST\",\"toDate\":\"10TH OCTOBER\",\"fromDate\":\"8TH OCTOBER\",\"toTime\":\"10 PM\",\"fromTime\":\"11 AM\",\"budget\":\"200000\",\"description\":\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\",\"category\":\"TECHNICAL\",\"venue\":\"ANNA AUDI\",\"attendance\":\"4000\",\"expectedParticipants\":\"4000\",\"facultyCoordinator\":{\"name\":\"NOORU MAA\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"studentCoordinator\":{\"name\":\"NOORU BAAP\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"PROrequest\":\"SAJDOOSIJANDFSAKFDSAFD\",\"campusEngineerRequest\":\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\",\"duration\":\"16 hours\",\"status\":\"\"}"
*            },
*            {
*                "ID": 6,
*                "CreatedAt": "2019-03-09T10:43:26.780351991Z",
*                "UpdatedAt": "2019-03-09T10:43:26.780351991Z",
*                "DeletedAt": null,
*                "Subject": "hades.event.CreateEvent",
*                "Timestamp": "2019-03-09 10:43:26.77995934 +0000 UTC m=+77.284892489",
*                "Body": "{\"clubName\":\"CC\",\"name\":\"DEVSOC\",\"toDate\":\"10TH OCTOBER\",\"fromDate\":\"8TH OCTOBER\",\"toTime\":\"10 PM\",\"fromTime\":\"11 AM\",\"budget\":\"200000\",\"description\":\"TECHNICAL EVENT AT CC VIT. ITS GONNA BE AMAZING\",\"category\":\"TECHNICAL\",\"venue\":\"ANNA AUDI\",\"attendance\":\"4000\",\"expectedParticipants\":\"4000\",\"facultyCoordinator\":{\"name\":\"NOOR\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"studentCoordinator\":{\"name\":\"NOORU\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"PROrequest\":\"SAJDOOSIJANDFSAKFDSAFD\",\"campusEngineerRequest\":\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\",\"duration\":\"16 hours\",\"status\":\"\"}"
*            },
*            {
*                "ID": 7,
*                "CreatedAt": "2019-03-09T10:52:49.264936787Z",
*                "UpdatedAt": "2019-03-09T10:52:49.264936787Z",
*                "DeletedAt": null,
*                "Subject": "hades.event.CreateEvent",
*                "Timestamp": "2019-03-09 10:52:49.264770904 +0000 UTC m=+10.214397568",
*                "Body": "{\"clubName\":\"GDG\",\"name\":\"WOMENTECHIxES\",\"toDate\":\"10TH OCTOBER\",\"fromDate\":\"8TH OCTOBER\",\"toTime\":\"10 PM\",\"fromTime\":\"11 AM\",\"budget\":\"200000\",\"description\":\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\",\"category\":\"TECHNICAL\",\"venue\":\"ANNA AUDI\",\"attendance\":\"4000\",\"expectedParticipants\":\"4000\",\"facultyCoordinator\":{\"name\":\"NOOR\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"studentCoordinator\":{\"name\":\"NOORU BAAP\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"PROrequest\":\"SAJDOOSIJANDFSAKFDSAFD\",\"campusEngineerRequest\":\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\",\"duration\":\"16 hours\",\"status\":\"\"}"
*            }
*        ]
*    ],
*    "errors": null
*}

**/
func readFromDB() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// make a map of all GET queries
		queries := r.URL.Query()
		var (
			arr    [][]db.Logs
			errs   []error
			buflen = len(queries)
			count  int
		)

		// make a channel with buffer length same as the number of queries
		ch := make(chan db.LogsReturn, buflen)
		for key, value := range queries {
			go db.ReadLogs(key, value[0], ch)
		}

		// query channel buffers, then close channel
		for i := range ch {
			count++
			if i.Err != nil {
				errs = append(errs, i.Err)
			}
			arr = append(arr, i.Logs)
			if count == buflen {
				break
			}
		}
		close(ch)

		// send JSON response
		json.NewEncoder(w).Encode(struct {
			Logs   [][]db.Logs `json:"logs"`
			Errors []error     `json:"errors"`
		}{
			arr,
			errs,
		})

	}
}

/**
*@api {get} /api/v1/analytics/all read all logs
*@apiGroup analytics
*@apiName read all logs
*@apiPermission admin
*
*@apiParamExample {json} request-url-example
*http://localhost:8085/api/v1/analytics/all
*
*@apiParamExample {json} response-example
*{
*    "logs": [
*        [
*            {
*                "ID": 1,
*                "CreatedAt": "2019-02-21T20:45:21.579518607Z",
*                "UpdatedAt": "2019-02-21T20:45:21.579518607Z",
*                "DeletedAt": null,
*                "Subject": "hades.event.CreateEvent",
*                "Timestamp": "2019-02-21 20:45:21.579267685 +0000 UTC m=+943.125297422",
*                "Body": "{\"clubName\":\"GDG\",\"name\":\"DEVRELCONF\",\"toDate\":\"10TH OCTOBER\",\"fromDate\":\"8TH OCTOBER\",\"toTime\":\"10 PM\",\"fromTime\":\"11 AM\",\"budget\":\"200000\",\"description\":\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\",\"category\":\"TECHNICAL\",\"venue\":\"ANNA AUDI\",\"attendance\":\"4000\",\"expectedParticipants\":\"4000\",\"facultyCoordinator\":{\"name\":\"NOORU MAA\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"studentCoordinator\":{\"name\":\"NOOR\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"PROrequest\":\"SAJDOOSIJANDFSAKFDSAFD\",\"campusEngineerRequest\":\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\",\"duration\":\"16 hours\",\"status\":\"\"}"
*            },
*            {
*                "ID": 2,
*                "CreatedAt": "2019-02-21T20:47:28.002589821Z",
*                "UpdatedAt": "2019-02-21T20:47:28.002589821Z",
*                "DeletedAt": null,
*                "Subject": "hades.event.CreateEvent",
*                "Timestamp": "2019-02-21 20:47:28.002438401 +0000 UTC m=+1069.548468003",
*                "Body": "{\"clubName\":\"GDG\",\"name\":\"DEVRELCONF\",\"toDate\":\"10TH OCTOBER\",\"fromDate\":\"8TH OCTOBER\",\"toTime\":\"10 PM\",\"fromTime\":\"11 AM\",\"budget\":\"200000\",\"description\":\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\",\"category\":\"TECHNICAL\",\"venue\":\"ANNA AUDI\",\"attendance\":\"4000\",\"expectedParticipants\":\"4000\",\"facultyCoordinator\":{\"name\":\"NOORU MAA\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"studentCoordinator\":{\"name\":\"NOOR\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"PROrequest\":\"SAJDOOSIJANDFSAKFDSAFD\",\"campusEngineerRequest\":\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\",\"duration\":\"16 hours\",\"status\":\"\"}"
*            },
*            {
*                "ID": 3,
*                "CreatedAt": "2019-03-09T10:29:40.344850552Z",
*                "UpdatedAt": "2019-03-09T10:29:40.344850552Z",
*                "DeletedAt": null,
*                "Subject": "hades.event.CreateEvent",
*                "Timestamp": "2019-03-09 10:29:40.344694322 +0000 UTC m=+327.442605414",
*                "Body": "{\"clubName\":\"\",\"name\":\"\",\"toDate\":\"\",\"fromDate\":\"\",\"toTime\":\"\",\"fromTime\":\"\",\"budget\":\"\",\"description\":\"\",\"category\":\"\",\"venue\":\"\",\"attendance\":\"\",\"expectedParticipants\":\"\",\"facultyCoordinator\":{\"name\":\"\",\"registrationNumber\":\"\",\"email\":\"\",\"phoneNumber\":\"\",\"gender\":\"\"},\"studentCoordinator\":{\"name\":\"\",\"registrationNumber\":\"\",\"email\":\"\",\"phoneNumber\":\"\",\"gender\":\"\"},\"PROrequest\":\"\",\"campusEngineerRequest\":\"\",\"duration\":\"\",\"status\":\"\"}"
*            },
*            {
*                "ID": 4,
*                "CreatedAt": "2019-03-09T10:39:59.807360183Z",
*                "UpdatedAt": "2019-03-09T10:39:59.807360183Z",
*                "DeletedAt": null,
*                "Subject": "hades.event.CreateEvent",
*                "Timestamp": "2019-03-09 10:39:59.807139694 +0000 UTC m=+4.791159116",
*                "Body": "{\"clubName\":\"\",\"name\":\"\",\"toDate\":\"\",\"fromDate\":\"\",\"toTime\":\"\",\"fromTime\":\"\",\"budget\":\"\",\"description\":\"\",\"category\":\"\",\"venue\":\"\",\"attendance\":\"\",\"expectedParticipants\":\"\",\"facultyCoordinator\":{\"name\":\"\",\"registrationNumber\":\"\",\"email\":\"\",\"phoneNumber\":\"\",\"gender\":\"\"},\"studentCoordinator\":{\"name\":\"\",\"registrationNumber\":\"\",\"email\":\"\",\"phoneNumber\":\"\",\"gender\":\"\"},\"PROrequest\":\"\",\"campusEngineerRequest\":\"\",\"duration\":\"\",\"status\":\"\"}"
*            },
*            {
*                "ID": 5,
*                "CreatedAt": "2019-03-09T10:42:27.834845067Z",
*                "UpdatedAt": "2019-03-09T10:42:27.834845067Z",
*                "DeletedAt": null,
*                "Subject": "hades.event.CreateEvent",
*                "Timestamp": "2019-03-09 10:42:27.834629395 +0000 UTC m=+18.339562524",
*                "Body": "{\"clubName\":\"GDG\",\"name\":\"DEVFEST\",\"toDate\":\"10TH OCTOBER\",\"fromDate\":\"8TH OCTOBER\",\"toTime\":\"10 PM\",\"fromTime\":\"11 AM\",\"budget\":\"200000\",\"description\":\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\",\"category\":\"TECHNICAL\",\"venue\":\"ANNA AUDI\",\"attendance\":\"4000\",\"expectedParticipants\":\"4000\",\"facultyCoordinator\":{\"name\":\"NOORU MAA\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"studentCoordinator\":{\"name\":\"NOORU BAAP\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"PROrequest\":\"SAJDOOSIJANDFSAKFDSAFD\",\"campusEngineerRequest\":\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\",\"duration\":\"16 hours\",\"status\":\"\"}"
*            },
*            {
*                "ID": 6,
*                "CreatedAt": "2019-03-09T10:43:26.780351991Z",
*                "UpdatedAt": "2019-03-09T10:43:26.780351991Z",
*                "DeletedAt": null,
*                "Subject": "hades.event.CreateEvent",
*                "Timestamp": "2019-03-09 10:43:26.77995934 +0000 UTC m=+77.284892489",
*                "Body": "{\"clubName\":\"CC\",\"name\":\"DEVSOC\",\"toDate\":\"10TH OCTOBER\",\"fromDate\":\"8TH OCTOBER\",\"toTime\":\"10 PM\",\"fromTime\":\"11 AM\",\"budget\":\"200000\",\"description\":\"TECHNICAL EVENT AT CC VIT. ITS GONNA BE AMAZING\",\"category\":\"TECHNICAL\",\"venue\":\"ANNA AUDI\",\"attendance\":\"4000\",\"expectedParticipants\":\"4000\",\"facultyCoordinator\":{\"name\":\"NOOR\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"studentCoordinator\":{\"name\":\"NOORU\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"PROrequest\":\"SAJDOOSIJANDFSAKFDSAFD\",\"campusEngineerRequest\":\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\",\"duration\":\"16 hours\",\"status\":\"\"}"
*            },
*            {
*                "ID": 7,
*                "CreatedAt": "2019-03-09T10:52:49.264936787Z",
*                "UpdatedAt": "2019-03-09T10:52:49.264936787Z",
*                "DeletedAt": null,
*                "Subject": "hades.event.CreateEvent",
*                "Timestamp": "2019-03-09 10:52:49.264770904 +0000 UTC m=+10.214397568",
*                "Body": "{\"clubName\":\"GDG\",\"name\":\"WOMENTECHIxES\",\"toDate\":\"10TH OCTOBER\",\"fromDate\":\"8TH OCTOBER\",\"toTime\":\"10 PM\",\"fromTime\":\"11 AM\",\"budget\":\"200000\",\"description\":\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\",\"category\":\"TECHNICAL\",\"venue\":\"ANNA AUDI\",\"attendance\":\"4000\",\"expectedParticipants\":\"4000\",\"facultyCoordinator\":{\"name\":\"NOOR\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"studentCoordinator\":{\"name\":\"NOORU BAAP\",\"registrationNumber\":\"17BBE1010\",\"email\":\"SDADAS@A.COM\",\"phoneNumber\":\"919191991911\",\"gender\":\"M\"},\"PROrequest\":\"SAJDOOSIJANDFSAKFDSAFD\",\"campusEngineerRequest\":\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\",\"duration\":\"16 hours\",\"status\":\"\"}"
*            }
*        ]
*    ],
*    "errors": null
*}

**/
func readAllFromDB() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// make a channel with buffer length same as the number of queries
		ch := make(chan db.LogsReturn)
		go db.ReadAllLogs(ch)

		msg := <-ch
		// send JSON response
		json.NewEncoder(w).Encode(db.LogsReturn{
			msg.Logs,
			msg.Err,
		})

	}
}
