define({ "api": [
  {
    "type": "get",
    "url": "/api/v1/analytics/all",
    "title": "read all logs",
    "group": "analytics",
    "name": "read_all_logs",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "examples": [
        {
          "title": "request-url-example",
          "content": "http://localhost:8085/api/v1/analytics/all",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"logs\": [\n       [\n           {\n               \"ID\": 1,\n               \"CreatedAt\": \"2019-02-21T20:45:21.579518607Z\",\n               \"UpdatedAt\": \"2019-02-21T20:45:21.579518607Z\",\n               \"DeletedAt\": null,\n               \"Subject\": \"hades.event.CreateEvent\",\n               \"Timestamp\": \"2019-02-21 20:45:21.579267685 +0000 UTC m=+943.125297422\",\n               \"Body\": \"{\\\"clubName\\\":\\\"GDG\\\",\\\"name\\\":\\\"DEVRELCONF\\\",\\\"toDate\\\":\\\"10TH OCTOBER\\\",\\\"fromDate\\\":\\\"8TH OCTOBER\\\",\\\"toTime\\\":\\\"10 PM\\\",\\\"fromTime\\\":\\\"11 AM\\\",\\\"budget\\\":\\\"200000\\\",\\\"description\\\":\\\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\\\",\\\"category\\\":\\\"TECHNICAL\\\",\\\"venue\\\":\\\"ANNA AUDI\\\",\\\"attendance\\\":\\\"4000\\\",\\\"expectedParticipants\\\":\\\"4000\\\",\\\"facultyCoordinator\\\":{\\\"name\\\":\\\"NOORU MAA\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"studentCoordinator\\\":{\\\"name\\\":\\\"NOOR\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"PROrequest\\\":\\\"SAJDOOSIJANDFSAKFDSAFD\\\",\\\"campusEngineerRequest\\\":\\\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\\\",\\\"duration\\\":\\\"16 hours\\\",\\\"status\\\":\\\"\\\"}\"\n           },\n           {\n               \"ID\": 2,\n               \"CreatedAt\": \"2019-02-21T20:47:28.002589821Z\",\n               \"UpdatedAt\": \"2019-02-21T20:47:28.002589821Z\",\n               \"DeletedAt\": null,\n               \"Subject\": \"hades.event.CreateEvent\",\n               \"Timestamp\": \"2019-02-21 20:47:28.002438401 +0000 UTC m=+1069.548468003\",\n               \"Body\": \"{\\\"clubName\\\":\\\"GDG\\\",\\\"name\\\":\\\"DEVRELCONF\\\",\\\"toDate\\\":\\\"10TH OCTOBER\\\",\\\"fromDate\\\":\\\"8TH OCTOBER\\\",\\\"toTime\\\":\\\"10 PM\\\",\\\"fromTime\\\":\\\"11 AM\\\",\\\"budget\\\":\\\"200000\\\",\\\"description\\\":\\\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\\\",\\\"category\\\":\\\"TECHNICAL\\\",\\\"venue\\\":\\\"ANNA AUDI\\\",\\\"attendance\\\":\\\"4000\\\",\\\"expectedParticipants\\\":\\\"4000\\\",\\\"facultyCoordinator\\\":{\\\"name\\\":\\\"NOORU MAA\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"studentCoordinator\\\":{\\\"name\\\":\\\"NOOR\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"PROrequest\\\":\\\"SAJDOOSIJANDFSAKFDSAFD\\\",\\\"campusEngineerRequest\\\":\\\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\\\",\\\"duration\\\":\\\"16 hours\\\",\\\"status\\\":\\\"\\\"}\"\n           },\n           {\n               \"ID\": 3,\n               \"CreatedAt\": \"2019-03-09T10:29:40.344850552Z\",\n               \"UpdatedAt\": \"2019-03-09T10:29:40.344850552Z\",\n               \"DeletedAt\": null,\n               \"Subject\": \"hades.event.CreateEvent\",\n               \"Timestamp\": \"2019-03-09 10:29:40.344694322 +0000 UTC m=+327.442605414\",\n               \"Body\": \"{\\\"clubName\\\":\\\"\\\",\\\"name\\\":\\\"\\\",\\\"toDate\\\":\\\"\\\",\\\"fromDate\\\":\\\"\\\",\\\"toTime\\\":\\\"\\\",\\\"fromTime\\\":\\\"\\\",\\\"budget\\\":\\\"\\\",\\\"description\\\":\\\"\\\",\\\"category\\\":\\\"\\\",\\\"venue\\\":\\\"\\\",\\\"attendance\\\":\\\"\\\",\\\"expectedParticipants\\\":\\\"\\\",\\\"facultyCoordinator\\\":{\\\"name\\\":\\\"\\\",\\\"registrationNumber\\\":\\\"\\\",\\\"email\\\":\\\"\\\",\\\"phoneNumber\\\":\\\"\\\",\\\"gender\\\":\\\"\\\"},\\\"studentCoordinator\\\":{\\\"name\\\":\\\"\\\",\\\"registrationNumber\\\":\\\"\\\",\\\"email\\\":\\\"\\\",\\\"phoneNumber\\\":\\\"\\\",\\\"gender\\\":\\\"\\\"},\\\"PROrequest\\\":\\\"\\\",\\\"campusEngineerRequest\\\":\\\"\\\",\\\"duration\\\":\\\"\\\",\\\"status\\\":\\\"\\\"}\"\n           },\n           {\n               \"ID\": 4,\n               \"CreatedAt\": \"2019-03-09T10:39:59.807360183Z\",\n               \"UpdatedAt\": \"2019-03-09T10:39:59.807360183Z\",\n               \"DeletedAt\": null,\n               \"Subject\": \"hades.event.CreateEvent\",\n               \"Timestamp\": \"2019-03-09 10:39:59.807139694 +0000 UTC m=+4.791159116\",\n               \"Body\": \"{\\\"clubName\\\":\\\"\\\",\\\"name\\\":\\\"\\\",\\\"toDate\\\":\\\"\\\",\\\"fromDate\\\":\\\"\\\",\\\"toTime\\\":\\\"\\\",\\\"fromTime\\\":\\\"\\\",\\\"budget\\\":\\\"\\\",\\\"description\\\":\\\"\\\",\\\"category\\\":\\\"\\\",\\\"venue\\\":\\\"\\\",\\\"attendance\\\":\\\"\\\",\\\"expectedParticipants\\\":\\\"\\\",\\\"facultyCoordinator\\\":{\\\"name\\\":\\\"\\\",\\\"registrationNumber\\\":\\\"\\\",\\\"email\\\":\\\"\\\",\\\"phoneNumber\\\":\\\"\\\",\\\"gender\\\":\\\"\\\"},\\\"studentCoordinator\\\":{\\\"name\\\":\\\"\\\",\\\"registrationNumber\\\":\\\"\\\",\\\"email\\\":\\\"\\\",\\\"phoneNumber\\\":\\\"\\\",\\\"gender\\\":\\\"\\\"},\\\"PROrequest\\\":\\\"\\\",\\\"campusEngineerRequest\\\":\\\"\\\",\\\"duration\\\":\\\"\\\",\\\"status\\\":\\\"\\\"}\"\n           },\n           {\n               \"ID\": 5,\n               \"CreatedAt\": \"2019-03-09T10:42:27.834845067Z\",\n               \"UpdatedAt\": \"2019-03-09T10:42:27.834845067Z\",\n               \"DeletedAt\": null,\n               \"Subject\": \"hades.event.CreateEvent\",\n               \"Timestamp\": \"2019-03-09 10:42:27.834629395 +0000 UTC m=+18.339562524\",\n               \"Body\": \"{\\\"clubName\\\":\\\"GDG\\\",\\\"name\\\":\\\"DEVFEST\\\",\\\"toDate\\\":\\\"10TH OCTOBER\\\",\\\"fromDate\\\":\\\"8TH OCTOBER\\\",\\\"toTime\\\":\\\"10 PM\\\",\\\"fromTime\\\":\\\"11 AM\\\",\\\"budget\\\":\\\"200000\\\",\\\"description\\\":\\\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\\\",\\\"category\\\":\\\"TECHNICAL\\\",\\\"venue\\\":\\\"ANNA AUDI\\\",\\\"attendance\\\":\\\"4000\\\",\\\"expectedParticipants\\\":\\\"4000\\\",\\\"facultyCoordinator\\\":{\\\"name\\\":\\\"NOORU MAA\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"studentCoordinator\\\":{\\\"name\\\":\\\"NOORU BAAP\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"PROrequest\\\":\\\"SAJDOOSIJANDFSAKFDSAFD\\\",\\\"campusEngineerRequest\\\":\\\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\\\",\\\"duration\\\":\\\"16 hours\\\",\\\"status\\\":\\\"\\\"}\"\n           },\n           {\n               \"ID\": 6,\n               \"CreatedAt\": \"2019-03-09T10:43:26.780351991Z\",\n               \"UpdatedAt\": \"2019-03-09T10:43:26.780351991Z\",\n               \"DeletedAt\": null,\n               \"Subject\": \"hades.event.CreateEvent\",\n               \"Timestamp\": \"2019-03-09 10:43:26.77995934 +0000 UTC m=+77.284892489\",\n               \"Body\": \"{\\\"clubName\\\":\\\"CC\\\",\\\"name\\\":\\\"DEVSOC\\\",\\\"toDate\\\":\\\"10TH OCTOBER\\\",\\\"fromDate\\\":\\\"8TH OCTOBER\\\",\\\"toTime\\\":\\\"10 PM\\\",\\\"fromTime\\\":\\\"11 AM\\\",\\\"budget\\\":\\\"200000\\\",\\\"description\\\":\\\"TECHNICAL EVENT AT CC VIT. ITS GONNA BE AMAZING\\\",\\\"category\\\":\\\"TECHNICAL\\\",\\\"venue\\\":\\\"ANNA AUDI\\\",\\\"attendance\\\":\\\"4000\\\",\\\"expectedParticipants\\\":\\\"4000\\\",\\\"facultyCoordinator\\\":{\\\"name\\\":\\\"NOOR\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"studentCoordinator\\\":{\\\"name\\\":\\\"NOORU\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"PROrequest\\\":\\\"SAJDOOSIJANDFSAKFDSAFD\\\",\\\"campusEngineerRequest\\\":\\\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\\\",\\\"duration\\\":\\\"16 hours\\\",\\\"status\\\":\\\"\\\"}\"\n           },\n           {\n               \"ID\": 7,\n               \"CreatedAt\": \"2019-03-09T10:52:49.264936787Z\",\n               \"UpdatedAt\": \"2019-03-09T10:52:49.264936787Z\",\n               \"DeletedAt\": null,\n               \"Subject\": \"hades.event.CreateEvent\",\n               \"Timestamp\": \"2019-03-09 10:52:49.264770904 +0000 UTC m=+10.214397568\",\n               \"Body\": \"{\\\"clubName\\\":\\\"GDG\\\",\\\"name\\\":\\\"WOMENTECHIxES\\\",\\\"toDate\\\":\\\"10TH OCTOBER\\\",\\\"fromDate\\\":\\\"8TH OCTOBER\\\",\\\"toTime\\\":\\\"10 PM\\\",\\\"fromTime\\\":\\\"11 AM\\\",\\\"budget\\\":\\\"200000\\\",\\\"description\\\":\\\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\\\",\\\"category\\\":\\\"TECHNICAL\\\",\\\"venue\\\":\\\"ANNA AUDI\\\",\\\"attendance\\\":\\\"4000\\\",\\\"expectedParticipants\\\":\\\"4000\\\",\\\"facultyCoordinator\\\":{\\\"name\\\":\\\"NOOR\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"studentCoordinator\\\":{\\\"name\\\":\\\"NOORU BAAP\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"PROrequest\\\":\\\"SAJDOOSIJANDFSAKFDSAFD\\\",\\\"campusEngineerRequest\\\":\\\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\\\",\\\"duration\\\":\\\"16 hours\\\",\\\"status\\\":\\\"\\\"}\"\n           }\n       ]\n   ],\n   \"errors\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "analytics/endpoints/handlerfuncs.go",
    "groupTitle": "analytics"
  },
  {
    "type": "get",
    "url": "/api/v1/analytics",
    "title": "read specific logs",
    "group": "analytics",
    "name": "read_specific_logs",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "subject",
            "description": "<p>event subject</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "timestamp",
            "description": "<p>event timestamp</p>"
          },
          {
            "group": "Parameter",
            "type": "String-Object",
            "optional": false,
            "field": "body",
            "description": "<p>event body</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-url-example",
          "content": "http://localhost:8085/api/v1/analytics?subject=hades.event.CreateEvent",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"logs\": [\n       [\n           {\n               \"ID\": 1,\n               \"CreatedAt\": \"2019-02-21T20:45:21.579518607Z\",\n               \"UpdatedAt\": \"2019-02-21T20:45:21.579518607Z\",\n               \"DeletedAt\": null,\n               \"Subject\": \"hades.event.CreateEvent\",\n               \"Timestamp\": \"2019-02-21 20:45:21.579267685 +0000 UTC m=+943.125297422\",\n               \"Body\": \"{\\\"clubName\\\":\\\"GDG\\\",\\\"name\\\":\\\"DEVRELCONF\\\",\\\"toDate\\\":\\\"10TH OCTOBER\\\",\\\"fromDate\\\":\\\"8TH OCTOBER\\\",\\\"toTime\\\":\\\"10 PM\\\",\\\"fromTime\\\":\\\"11 AM\\\",\\\"budget\\\":\\\"200000\\\",\\\"description\\\":\\\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\\\",\\\"category\\\":\\\"TECHNICAL\\\",\\\"venue\\\":\\\"ANNA AUDI\\\",\\\"attendance\\\":\\\"4000\\\",\\\"expectedParticipants\\\":\\\"4000\\\",\\\"facultyCoordinator\\\":{\\\"name\\\":\\\"NOORU MAA\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"studentCoordinator\\\":{\\\"name\\\":\\\"NOOR\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"PROrequest\\\":\\\"SAJDOOSIJANDFSAKFDSAFD\\\",\\\"campusEngineerRequest\\\":\\\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\\\",\\\"duration\\\":\\\"16 hours\\\",\\\"status\\\":\\\"\\\"}\"\n           },\n           {\n               \"ID\": 2,\n               \"CreatedAt\": \"2019-02-21T20:47:28.002589821Z\",\n               \"UpdatedAt\": \"2019-02-21T20:47:28.002589821Z\",\n               \"DeletedAt\": null,\n               \"Subject\": \"hades.event.CreateEvent\",\n               \"Timestamp\": \"2019-02-21 20:47:28.002438401 +0000 UTC m=+1069.548468003\",\n               \"Body\": \"{\\\"clubName\\\":\\\"GDG\\\",\\\"name\\\":\\\"DEVRELCONF\\\",\\\"toDate\\\":\\\"10TH OCTOBER\\\",\\\"fromDate\\\":\\\"8TH OCTOBER\\\",\\\"toTime\\\":\\\"10 PM\\\",\\\"fromTime\\\":\\\"11 AM\\\",\\\"budget\\\":\\\"200000\\\",\\\"description\\\":\\\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\\\",\\\"category\\\":\\\"TECHNICAL\\\",\\\"venue\\\":\\\"ANNA AUDI\\\",\\\"attendance\\\":\\\"4000\\\",\\\"expectedParticipants\\\":\\\"4000\\\",\\\"facultyCoordinator\\\":{\\\"name\\\":\\\"NOORU MAA\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"studentCoordinator\\\":{\\\"name\\\":\\\"NOOR\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"PROrequest\\\":\\\"SAJDOOSIJANDFSAKFDSAFD\\\",\\\"campusEngineerRequest\\\":\\\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\\\",\\\"duration\\\":\\\"16 hours\\\",\\\"status\\\":\\\"\\\"}\"\n           },\n           {\n               \"ID\": 3,\n               \"CreatedAt\": \"2019-03-09T10:29:40.344850552Z\",\n               \"UpdatedAt\": \"2019-03-09T10:29:40.344850552Z\",\n               \"DeletedAt\": null,\n               \"Subject\": \"hades.event.CreateEvent\",\n               \"Timestamp\": \"2019-03-09 10:29:40.344694322 +0000 UTC m=+327.442605414\",\n               \"Body\": \"{\\\"clubName\\\":\\\"\\\",\\\"name\\\":\\\"\\\",\\\"toDate\\\":\\\"\\\",\\\"fromDate\\\":\\\"\\\",\\\"toTime\\\":\\\"\\\",\\\"fromTime\\\":\\\"\\\",\\\"budget\\\":\\\"\\\",\\\"description\\\":\\\"\\\",\\\"category\\\":\\\"\\\",\\\"venue\\\":\\\"\\\",\\\"attendance\\\":\\\"\\\",\\\"expectedParticipants\\\":\\\"\\\",\\\"facultyCoordinator\\\":{\\\"name\\\":\\\"\\\",\\\"registrationNumber\\\":\\\"\\\",\\\"email\\\":\\\"\\\",\\\"phoneNumber\\\":\\\"\\\",\\\"gender\\\":\\\"\\\"},\\\"studentCoordinator\\\":{\\\"name\\\":\\\"\\\",\\\"registrationNumber\\\":\\\"\\\",\\\"email\\\":\\\"\\\",\\\"phoneNumber\\\":\\\"\\\",\\\"gender\\\":\\\"\\\"},\\\"PROrequest\\\":\\\"\\\",\\\"campusEngineerRequest\\\":\\\"\\\",\\\"duration\\\":\\\"\\\",\\\"status\\\":\\\"\\\"}\"\n           },\n           {\n               \"ID\": 4,\n               \"CreatedAt\": \"2019-03-09T10:39:59.807360183Z\",\n               \"UpdatedAt\": \"2019-03-09T10:39:59.807360183Z\",\n               \"DeletedAt\": null,\n               \"Subject\": \"hades.event.CreateEvent\",\n               \"Timestamp\": \"2019-03-09 10:39:59.807139694 +0000 UTC m=+4.791159116\",\n               \"Body\": \"{\\\"clubName\\\":\\\"\\\",\\\"name\\\":\\\"\\\",\\\"toDate\\\":\\\"\\\",\\\"fromDate\\\":\\\"\\\",\\\"toTime\\\":\\\"\\\",\\\"fromTime\\\":\\\"\\\",\\\"budget\\\":\\\"\\\",\\\"description\\\":\\\"\\\",\\\"category\\\":\\\"\\\",\\\"venue\\\":\\\"\\\",\\\"attendance\\\":\\\"\\\",\\\"expectedParticipants\\\":\\\"\\\",\\\"facultyCoordinator\\\":{\\\"name\\\":\\\"\\\",\\\"registrationNumber\\\":\\\"\\\",\\\"email\\\":\\\"\\\",\\\"phoneNumber\\\":\\\"\\\",\\\"gender\\\":\\\"\\\"},\\\"studentCoordinator\\\":{\\\"name\\\":\\\"\\\",\\\"registrationNumber\\\":\\\"\\\",\\\"email\\\":\\\"\\\",\\\"phoneNumber\\\":\\\"\\\",\\\"gender\\\":\\\"\\\"},\\\"PROrequest\\\":\\\"\\\",\\\"campusEngineerRequest\\\":\\\"\\\",\\\"duration\\\":\\\"\\\",\\\"status\\\":\\\"\\\"}\"\n           },\n           {\n               \"ID\": 5,\n               \"CreatedAt\": \"2019-03-09T10:42:27.834845067Z\",\n               \"UpdatedAt\": \"2019-03-09T10:42:27.834845067Z\",\n               \"DeletedAt\": null,\n               \"Subject\": \"hades.event.CreateEvent\",\n               \"Timestamp\": \"2019-03-09 10:42:27.834629395 +0000 UTC m=+18.339562524\",\n               \"Body\": \"{\\\"clubName\\\":\\\"GDG\\\",\\\"name\\\":\\\"DEVFEST\\\",\\\"toDate\\\":\\\"10TH OCTOBER\\\",\\\"fromDate\\\":\\\"8TH OCTOBER\\\",\\\"toTime\\\":\\\"10 PM\\\",\\\"fromTime\\\":\\\"11 AM\\\",\\\"budget\\\":\\\"200000\\\",\\\"description\\\":\\\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\\\",\\\"category\\\":\\\"TECHNICAL\\\",\\\"venue\\\":\\\"ANNA AUDI\\\",\\\"attendance\\\":\\\"4000\\\",\\\"expectedParticipants\\\":\\\"4000\\\",\\\"facultyCoordinator\\\":{\\\"name\\\":\\\"NOORU MAA\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"studentCoordinator\\\":{\\\"name\\\":\\\"NOORU BAAP\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"PROrequest\\\":\\\"SAJDOOSIJANDFSAKFDSAFD\\\",\\\"campusEngineerRequest\\\":\\\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\\\",\\\"duration\\\":\\\"16 hours\\\",\\\"status\\\":\\\"\\\"}\"\n           },\n           {\n               \"ID\": 6,\n               \"CreatedAt\": \"2019-03-09T10:43:26.780351991Z\",\n               \"UpdatedAt\": \"2019-03-09T10:43:26.780351991Z\",\n               \"DeletedAt\": null,\n               \"Subject\": \"hades.event.CreateEvent\",\n               \"Timestamp\": \"2019-03-09 10:43:26.77995934 +0000 UTC m=+77.284892489\",\n               \"Body\": \"{\\\"clubName\\\":\\\"CC\\\",\\\"name\\\":\\\"DEVSOC\\\",\\\"toDate\\\":\\\"10TH OCTOBER\\\",\\\"fromDate\\\":\\\"8TH OCTOBER\\\",\\\"toTime\\\":\\\"10 PM\\\",\\\"fromTime\\\":\\\"11 AM\\\",\\\"budget\\\":\\\"200000\\\",\\\"description\\\":\\\"TECHNICAL EVENT AT CC VIT. ITS GONNA BE AMAZING\\\",\\\"category\\\":\\\"TECHNICAL\\\",\\\"venue\\\":\\\"ANNA AUDI\\\",\\\"attendance\\\":\\\"4000\\\",\\\"expectedParticipants\\\":\\\"4000\\\",\\\"facultyCoordinator\\\":{\\\"name\\\":\\\"NOOR\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"studentCoordinator\\\":{\\\"name\\\":\\\"NOORU\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"PROrequest\\\":\\\"SAJDOOSIJANDFSAKFDSAFD\\\",\\\"campusEngineerRequest\\\":\\\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\\\",\\\"duration\\\":\\\"16 hours\\\",\\\"status\\\":\\\"\\\"}\"\n           },\n           {\n               \"ID\": 7,\n               \"CreatedAt\": \"2019-03-09T10:52:49.264936787Z\",\n               \"UpdatedAt\": \"2019-03-09T10:52:49.264936787Z\",\n               \"DeletedAt\": null,\n               \"Subject\": \"hades.event.CreateEvent\",\n               \"Timestamp\": \"2019-03-09 10:52:49.264770904 +0000 UTC m=+10.214397568\",\n               \"Body\": \"{\\\"clubName\\\":\\\"GDG\\\",\\\"name\\\":\\\"WOMENTECHIxES\\\",\\\"toDate\\\":\\\"10TH OCTOBER\\\",\\\"fromDate\\\":\\\"8TH OCTOBER\\\",\\\"toTime\\\":\\\"10 PM\\\",\\\"fromTime\\\":\\\"11 AM\\\",\\\"budget\\\":\\\"200000\\\",\\\"description\\\":\\\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\\\",\\\"category\\\":\\\"TECHNICAL\\\",\\\"venue\\\":\\\"ANNA AUDI\\\",\\\"attendance\\\":\\\"4000\\\",\\\"expectedParticipants\\\":\\\"4000\\\",\\\"facultyCoordinator\\\":{\\\"name\\\":\\\"NOOR\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"studentCoordinator\\\":{\\\"name\\\":\\\"NOORU BAAP\\\",\\\"registrationNumber\\\":\\\"17BBE1010\\\",\\\"email\\\":\\\"SDADAS@A.COM\\\",\\\"phoneNumber\\\":\\\"919191991911\\\",\\\"gender\\\":\\\"M\\\"},\\\"PROrequest\\\":\\\"SAJDOOSIJANDFSAKFDSAFD\\\",\\\"campusEngineerRequest\\\":\\\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\\\",\\\"duration\\\":\\\"16 hours\\\",\\\"status\\\":\\\"\\\"}\"\n           }\n       ]\n   ],\n   \"errors\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "analytics/endpoints/handlerfuncs.go",
    "groupTitle": "analytics"
  },
  {
    "type": "post",
    "url": "/api/v1/coupons/create-schema",
    "title": "create coupon schema",
    "name": "create_coupon_schema",
    "group": "coupons",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "event",
            "description": "<p>name of the event</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "name",
            "description": "<p>name of the coupon</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "description",
            "description": "<p>description of the coupon</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "day",
            "description": "<p>day of the event</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"event\":\"DEVRELCONF\",\n\t\"coupons\": [{\n\t\t\"name\":\"lunch\",\n\t\t\"description\":\"lunch\",\n\t\t\"day\":1\n\t},{\n\t\t\"name\":\"bf\",\n\t\t\"description\":\"bf\",\n\t\t\"day\":1\n\t},{\n\t\t\"name\":\"lunch\",\n\t\t\"description\":\"lunch\",\n\t\t\"day\":2\n\t},{\n\t\t\"name\":\"dinner\",\n\t\t\"description\":\"dinner\",\n\t\t\"day\":2\n\t}]\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": \"Successfully creted coupon schema for event DEVRELCONF\",\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "coupons/pkg/service/service.go",
    "groupTitle": "coupons"
  },
  {
    "type": "post",
    "url": "/api/v1/coupons/delete-coupon",
    "title": "delete coupon",
    "name": "delete_coupon",
    "group": "coupons",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "eventName",
            "description": "<p>name of the event</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "name",
            "description": "<p>name of the coupon</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "day",
            "description": "<p>day of the event</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "description",
            "description": "<p>description of the coupon</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "delete-specific-request",
          "content": "{\n\t\"event\":\"DEVRELCONF\",\n\t\"query\":{\n\t\t\"name\":\"lunch\",\n\t\t\"day\":2,\n\t\t\"description\":\"lunch\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "delete-specific-response",
          "content": "{\n   \"rs\": \"Deleted\",\n   \"err\": null\n}",
          "type": "json"
        },
        {
          "title": "delete-all-request",
          "content": "{\n\t\"event\":\"DEVRELCONF\"\n}",
          "type": "json"
        },
        {
          "title": "delete-all-response",
          "content": "{\n   \"rs\": \"Deleted\",\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "coupons/pkg/service/service.go",
    "groupTitle": "coupons"
  },
  {
    "type": "post",
    "url": "/api/v1/coupons/mark-present",
    "title": "mark attendee present",
    "name": "mark_attendee_present",
    "group": "coupons",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "eventName",
            "description": "<p>name of the event</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>email of the attendee</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "day",
            "description": "<p>day of the event</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"attendance\":{\n\t\t\"eventName\":\"DEVRELCONF\",\n\t\t\"email\":\"a@a.COM\",\n\t\t\"day\":2\n\t}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": \"Successfully marked present for the day\",\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "coupons/pkg/service/service.go",
    "groupTitle": "coupons"
  },
  {
    "type": "post",
    "url": "/api/v1/coupons/redeem-coupon",
    "title": "redeem a coupon",
    "name": "redeem_a_coupon",
    "group": "coupons",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "eventName",
            "description": "<p>name of the event</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>email of the attendee</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "day",
            "description": "<p>day of the event</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "couponName",
            "description": "<p>name of the coupon</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"attendance\": {\n\t\t\"couponName\":\"dinner\",\n\t\t\"day\":2,\n\t\t\"email\":\"a@a.COM\",\n\t\t\"eventName\":\"DEVRELCONF\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": \"No match found for this coupon\",\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "coupons/pkg/service/service.go",
    "groupTitle": "coupons"
  },
  {
    "type": "post",
    "url": "/api/v1/coupons/view-schema",
    "title": "view coupon schema",
    "name": "view_coupon_schema",
    "group": "coupons",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "event",
            "description": "<p>name of the event</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"event\":\"DEVRELCONF\"\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": [\n       {\n           \"name\": \"gargentaul\",\n           \"description\": \"dasd\",\n           \"day\": 1\n       },\n       {\n           \"name\": \"bf\",\n           \"description\": \"bf\",\n           \"day\": 1\n       },\n       {\n           \"name\": \"dinner\",\n           \"description\": \"dinner\",\n           \"day\": 2\n       },\n       {\n           \"name\": \"lunch\",\n           \"description\": \"lunch\",\n           \"day\": 2\n       },\n       {\n           \"name\": \"lunch\",\n           \"description\": \"lunch\",\n           \"day\": 1\n       }\n   ],\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "coupons/pkg/service/service.go",
    "groupTitle": "coupons"
  },
  {
    "type": "post",
    "url": "/api/v1/event/create-event",
    "title": "create a new event",
    "name": "create_a_new_event",
    "group": "events",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "sampleRequest": [
      {
        "url": "http://localhost:8800/"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "clubName",
            "description": "<p>Name of your club</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Name of your event</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "toDate",
            "description": "<p>ending date</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "fromDate",
            "description": "<p>start date</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "toTime",
            "description": "<p>start time</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "fromTime",
            "description": "<p>ending time</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "budget",
            "description": "<p>budget</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "description",
            "description": "<p>event description</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "category",
            "description": "<p>category of the event</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "venue",
            "description": "<p>event venue</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "attendance",
            "description": "<p>Name of your club</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "expectedParticipants",
            "description": "<p>Expected Participants in the event</p>"
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "facultyCoordinator",
            "description": "<p>faculty coordinator details (Participant Object)</p>"
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "studentCoordinator",
            "description": "<p>student coordinator details (Participant Object)</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "PROrequest",
            "description": "<p>PRO department requests</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "campusEngineerRequest",
            "description": "<p>Campus engineer requests</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "duration",
            "description": "<p>duration of event</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\"event\":{\n \"clubName\":\"GDG\",\n \"name\":\"DEVRELCONF\",\n \"toDate\":\"10TH OCTOBER\",\n \"fromDate\":\"8TH OCTOBER\",\n \"toTime\":\"10 PM\",\n \"fromTime\":\"11 AM\",\n \"budget\":\"200000\",\n \"description\":\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\",\n \"category\":\"TECHNICAL\",\n \"venue\":\"ANNA AUDI\",\n \"attendance\":\"4000\",\n \"expectedParticipants\":\"4000\",\n \"facultyCoordinator\":{\n    \"name\":\"NOORU MAA\",\n    \"registrationNumber\":\"17BBE1010\",\n    \"email\":\"SDADAS@A.COM\",\n    \"phoneNumber\":\"919191991911\",\n    \"gender\":\"M\",\n    \"eventsAttended\":\"ALL\"\n },\n \"studentCoordinator\":{\n    \"name\":\"NOOR\",\n    \"registrationNumber\":\"17BBE1010\",\n    \"email\":\"SDADAS@A.COM\",\n    \"phoneNumber\":\"919191991911\",\n    \"gender\":\"M\",\n    \"eventsAttended\":\"ALL\"\n },\n \"PROrequest\":\"SAJDOOSIJANDFSAKFDSAFD\",\n \"campusEngineerRequest\":\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\",\n \"duration\":\"16 hours\"\n}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n\trs:\"created\",\n\terr:null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "events/pkg/service/service.go",
    "groupTitle": "events"
  },
  {
    "type": "post",
    "url": "/api/v1/event/create-event",
    "title": "create a new event",
    "name": "create_a_new_event",
    "group": "events",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "sampleRequest": [
      {
        "url": "http://localhost:8800/"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "clubName",
            "description": "<p>Name of your club</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Name of your event</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "toDate",
            "description": "<p>ending date</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "fromDate",
            "description": "<p>start date</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "toTime",
            "description": "<p>start time</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "fromTime",
            "description": "<p>ending time</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "budget",
            "description": "<p>budget</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "description",
            "description": "<p>event description</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "category",
            "description": "<p>category of the event</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "venue",
            "description": "<p>event venue</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "attendance",
            "description": "<p>Name of your club</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "expectedParticipants",
            "description": "<p>Expected Participants in the event</p>"
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "facultyCoordinator",
            "description": "<p>faculty coordinator details (Participant Object)</p>"
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "studentCoordinator",
            "description": "<p>student coordinator details (Participant Object)</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "PROrequest",
            "description": "<p>PRO department requests</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "campusEngineerRequest",
            "description": "<p>Campus engineer requests</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "duration",
            "description": "<p>duration of event</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\"event\":{\n \"clubName\":\"GDG\",\n \"name\":\"DEVRELCONF\",\n \"toDate\":\"10TH OCTOBER\",\n \"fromDate\":\"8TH OCTOBER\",\n \"toTime\":\"10 PM\",\n \"fromTime\":\"11 AM\",\n \"budget\":\"200000\",\n \"description\":\"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\",\n \"category\":\"TECHNICAL\",\n \"venue\":\"ANNA AUDI\",\n \"attendance\":\"4000\",\n \"expectedParticipants\":\"4000\",\n \"facultyCoordinator\":{\n    \"name\":\"NOORU MAA\",\n    \"registrationNumber\":\"17BBE1010\",\n    \"email\":\"SDADAS@A.COM\",\n    \"phoneNumber\":\"919191991911\",\n    \"gender\":\"M\",\n    \"eventsAttended\":\"ALL\"\n },\n \"studentCoordinator\":{\n    \"name\":\"NOOR\",\n    \"registrationNumber\":\"17BBE1010\",\n    \"email\":\"SDADAS@A.COM\",\n    \"phoneNumber\":\"919191991911\",\n    \"gender\":\"M\",\n    \"eventsAttended\":\"ALL\"\n },\n \"PROrequest\":\"SAJDOOSIJANDFSAKFDSAFD\",\n \"campusEngineerRequest\":\"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\",\n \"duration\":\"16 hours\"\n}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n\trs:\"created\",\n\terr:null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "events/pkg/service/service.go",
    "groupTitle": "events"
  },
  {
    "type": "delete",
    "url": "/api/v1/event/delete-event",
    "title": "delete an event",
    "name": "delete_an_event",
    "group": "events",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "key",
            "description": "<p>key to query the event by</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "value",
            "description": "<p>value of the key</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"query\":{\n\t\t\"key\":\"clubName\",\n\t\t\"value\":\"GDG\"\n\t\t\"organization\":\"CodeChef-VIT\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n\t rs:\"deleted\",\n\t err:null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "events/pkg/service/service.go",
    "groupTitle": "events"
  },
  {
    "type": "delete",
    "url": "/api/v1/event/delete-event",
    "title": "delete an event",
    "name": "delete_an_event",
    "group": "events",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "key",
            "description": "<p>key to query the event by</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "value",
            "description": "<p>value of the key</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"query\":{\n\t\t\"key\":\"clubName\",\n\t\t\"value\":\"GDG\"\n\t\t\"organization\":\"CodeChef-VIT\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n\t rs:\"deleted\",\n\t err:null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "events/pkg/service/service.go",
    "groupTitle": "events"
  },
  {
    "type": "get",
    "url": "/api/v1/event/read-event",
    "title": "read an event",
    "name": "read_an_event",
    "group": "events",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "key",
            "description": "<p>key to query the event by</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "value",
            "description": "<p>value of the key</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "specific",
            "description": "<p>search by name of the event</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "   {\"query\":{\n\t\t\"key\":\"clubName\",\n\t\t\"value\":\"GDG\",\n\t\t\"specific\":\"DEVFEST 2019\"\n\t\t\"organization\":\"CodeChef-VIT\"\n\t}}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": [{\n       \"clubName\": \"GDG\",\n       \"name\": \"DEVRELCONF\",\n       \"toDate\": \"10TH OCTOBER\",\n       \"fromDate\": \"8TH OCTOBER\",\n       \"toTime\": \"10 PM\",\n       \"fromTime\": \"11 AM\",\n       \"budget\": \"200000\",\n       \"description\": \"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\",\n       \"category\": \"TECHNICAL\",\n       \"venue\": \"ANNA AUDI\",\n       \"attendance\": \"4000\",\n       \"expectedParticipants\": \"4000\",\n       \"facultyCoordinator\": {\n           \"name\": \"NOORU MAA\",\n           \"registrationNumber\": \"17BBE1010\",\n           \"email\": \"SDADAS@A.COM\",\n           \"phoneNumber\": \"919191991911\",\n           \"gender\": \"M\"\n       },\n       \"studentCoordinator\": {\n           \"name\": \"NOORU BAAP\",\n           \"registrationNumber\": \"17BBE1010\",\n           \"email\": \"SDADAS@A.COM\",\n           \"phoneNumber\": \"919191991911\",\n           \"gender\": \"M\"\n       },\n       \"PROrequest\": \"SAJDOOSIJANDFSAKFDSAFD\",\n       \"campusEngineerRequest\": \"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\",\n       \"duration\": \"16 hours\",\n       \"mainSponsor\": {\n           \"name\": \"\",\n           \"registrationNumber\": \"\",\n           \"email\": \"\",\n           \"phoneNumber\": \"\",\n           \"gender\": \"\"\n       },\n       \"status\": \"false\"\n\n   }],\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "events/pkg/service/service.go",
    "groupTitle": "events"
  },
  {
    "type": "get",
    "url": "/api/v1/event/read-event",
    "title": "read an event",
    "name": "read_an_event",
    "group": "events",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "key",
            "description": "<p>key to query the event by</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "value",
            "description": "<p>value of the key</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "specific",
            "description": "<p>search by name of the event</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "   {\"query\":{\n\t\t\"key\":\"clubName\",\n\t\t\"value\":\"GDG\",\n\t\t\"specific\":\"DEVFEST 2019\"\n\t\t\"organization\":\"CodeChef-VIT\"\n\t}}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": [{\n       \"clubName\": \"GDG\",\n       \"name\": \"DEVRELCONF\",\n       \"toDate\": \"10TH OCTOBER\",\n       \"fromDate\": \"8TH OCTOBER\",\n       \"toTime\": \"10 PM\",\n       \"fromTime\": \"11 AM\",\n       \"budget\": \"200000\",\n       \"description\": \"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING\",\n       \"category\": \"TECHNICAL\",\n       \"venue\": \"ANNA AUDI\",\n       \"attendance\": \"4000\",\n       \"expectedParticipants\": \"4000\",\n       \"facultyCoordinator\": {\n           \"name\": \"NOORU MAA\",\n           \"registrationNumber\": \"17BBE1010\",\n           \"email\": \"SDADAS@A.COM\",\n           \"phoneNumber\": \"919191991911\",\n           \"gender\": \"M\"\n       },\n       \"studentCoordinator\": {\n           \"name\": \"NOORU BAAP\",\n           \"registrationNumber\": \"17BBE1010\",\n           \"email\": \"SDADAS@A.COM\",\n           \"phoneNumber\": \"919191991911\",\n           \"gender\": \"M\"\n       },\n       \"PROrequest\": \"SAJDOOSIJANDFSAKFDSAFD\",\n       \"campusEngineerRequest\": \"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD\",\n       \"duration\": \"16 hours\",\n       \"mainSponsor\": {\n           \"name\": \"\",\n           \"registrationNumber\": \"\",\n           \"email\": \"\",\n           \"phoneNumber\": \"\",\n           \"gender\": \"\"\n       },\n       \"status\": \"false\"\n\n   }],\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "events/pkg/service/service.go",
    "groupTitle": "events"
  },
  {
    "type": "put",
    "url": "/api/v1/event/update-event",
    "title": "update an event",
    "name": "update_an_event",
    "group": "events",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "key",
            "description": "<p>key to query the event by</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "value",
            "description": "<p>value of the key</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "changeKey",
            "description": "<p>key of the value which needs to be altered</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "changeValue",
            "description": "<p>the new value</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"query\":{\n\t\t\"key\":\"clubName\",\n\t\t\"value\":\"GDG\",\n\t\t\"changeKey\":\"clubName\",\n\t\t\"changeValue\":\"codechef\"\n\t\t\"organization\":\"CodeChef-VIT\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n\trs:\"updated\",\n\terr:null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "events/pkg/service/service.go",
    "groupTitle": "events"
  },
  {
    "type": "put",
    "url": "/api/v1/event/update-event",
    "title": "update an event",
    "name": "update_an_event",
    "group": "events",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "key",
            "description": "<p>key to query the event by</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "value",
            "description": "<p>value of the key</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "changeKey",
            "description": "<p>key of the value which needs to be altered</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "changeValue",
            "description": "<p>the new value</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"query\":{\n\t\t\"key\":\"clubName\",\n\t\t\"value\":\"GDG\",\n\t\t\"changeKey\":\"clubName\",\n\t\t\"changeValue\":\"codechef\"\n\t\t\"organization\":\"CodeChef-VIT\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n\trs:\"updated\",\n\terr:null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "events/pkg/service/service.go",
    "groupTitle": "events"
  },
  {
    "type": "get",
    "url": "/api/v1/exporter/excel",
    "title": "export participants as excel/csv",
    "name": "export_participants_as_excel_csv",
    "group": "exporter",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "key",
            "description": "<p>key to query the event by</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "value",
            "description": "<p>value of the key</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "event",
            "description": "<p>event name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "specific",
            "description": "<p>types of participants to export i.e. present/absent/all</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example1",
          "content": "{\n\t\"event\":\"DEVFEST 2019\",\n\t\"query\":{\n\t\t\"key\":\"gender\",\n\t\t\"value\":\"F\",\n\t\t\"specific\":\"project-all\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "request-example2",
          "content": "{\n\t\"event\":\"DEVFEST 2019\",\n\t\"query\":{\n\t\t\"key\":\"gender\",\n\t\t\"value\":\"F\",\n\t\t\"specific\":\"project-present\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "request-example3",
          "content": "{\n\t\"event\":\"DEVFEST 2019\",\n\t\"query\":{\n\t\t\"key\":\"gender\",\n\t\t\"value\":\"F\",\n\t\t\"specific\":\"project-absent\"\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "exporter/endpoints/controller/handlerfuncs.go",
    "groupTitle": "exporter"
  },
  {
    "type": "get",
    "url": "/api/v1/exporter/json",
    "title": "export participants as json",
    "name": "export_participants_as_json",
    "group": "exporter",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "key",
            "description": "<p>key to query the event by</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "value",
            "description": "<p>value of the key</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "event",
            "description": "<p>event name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "specific",
            "description": "<p>types of participants to export i.e. present/absent/all</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example1",
          "content": "{\n\t\"event\":\"DEVFEST 2019\",\n\t\"query\":{\n\t\t\"key\":\"gender\",\n\t\t\"value\":\"F\",\n\t\t\"specific\":\"project-all\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "request-example2",
          "content": "{\n\t\"event\":\"DEVFEST 2019\",\n\t\"query\":{\n\t\t\"key\":\"gender\",\n\t\t\"value\":\"F\",\n\t\t\"specific\":\"project-present\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "request-example3",
          "content": "{\n\t\"event\":\"DEVFEST 2019\",\n\t\"query\":{\n\t\t\"key\":\"gender\",\n\t\t\"value\":\"F\",\n\t\t\"specific\":\"project-absent\"\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "exporter/endpoints/controller/handlerfuncs.go",
    "groupTitle": "exporter"
  },
  {
    "type": "post",
    "url": "/api/v1/guests/create-guest",
    "title": "create a guest",
    "name": "create_a_guest",
    "group": "guest",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "event",
            "description": "<p>name of the event</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "name",
            "description": "<p>name of the guest</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>email of the guest</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "phoneNumber",
            "description": "<p>phone number of the guest</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "gender",
            "description": "<p>gender of the guest</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "stake",
            "description": "<p>stake of the guest</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "locationOfStay",
            "description": "<p>where does the guest stay? (origin)</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"event\":\"DEVRELCONF\",\n\t\"guest\": {\n\t\t\"name\":\"angad\",\n\t\t\"email\":\"sdaasd@a.com\",\n\t\t\"phoneNumber\":\"9999999999\",\n\t\t\"gender\":\"M\",\n\t\t\"stake\":\"speaker\",\n\t\t\"locationOfStay\":\"Bangalore\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": \"Guest created\",\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "guests/pkg/service/service.go",
    "groupTitle": "guest"
  },
  {
    "type": "post",
    "url": "/api/v1/guests/delete-guest",
    "title": "delete a guest",
    "name": "delete_a_guest",
    "group": "guest",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "key",
            "description": "<p>key of the event field</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "value",
            "description": "<p>value of the event field</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "changeKey",
            "description": "<p>Name of the guest to be deleted</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"query\": {\n\t\t\"key\":\"name\",\n\t\t\"value\":\"DEVRELCONF\",\n\t\t\"changeKey\":\"angad\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": \"Guest deleted\",\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "guests/pkg/service/service.go",
    "groupTitle": "guest"
  },
  {
    "type": "post",
    "url": "/api/v1/guests/read-guest",
    "title": "read a guest",
    "name": "read_a_guest",
    "group": "guest",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "key",
            "description": "<p>key of the event field</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "value",
            "description": "<p>value of the event field</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"query\": {\n\t\t\"key\":\"name\",\n\t\t\"value\":\"DEVRELCONF\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": [\n       {\n           \"name\": \"dafds\",\n           \"email\": \"sdadsaadasd@a.com\",\n           \"phoneNumber\": \"9699999999\",\n           \"gender\": \"M\",\n           \"stake\": \"speaker\",\n           \"locationOfStay\": \"Bangalore\"\n       },\n       {\n           \"name\": \"angad\",\n           \"email\": \"sdaasd@a.com\",\n           \"phoneNumber\": \"9999999999\",\n           \"gender\": \"M\",\n           \"stake\": \"speaker\",\n           \"locationOfStay\": \"Bangalore\"\n       }\n   ],\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "guests/pkg/service/service.go",
    "groupTitle": "guest"
  },
  {
    "type": "post",
    "url": "/api/v1/org/accept",
    "title": "Accept join request",
    "name": "Accept_join_request",
    "group": "organization",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "admin",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "org",
            "description": "<p>name of the organization as query param</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"email\":\"test1@test.com\",\n\t\"org\":\"GDG-VIT\"\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "\n{\"message\":\"Added as a member\"}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "organization/endpoints/join_handler.go",
    "groupTitle": "organization"
  },
  {
    "type": "get",
    "url": "/api/v1/org/org-events",
    "title": "Get events and organizations",
    "name": "Get_events_and_organizations",
    "group": "organization",
    "permission": [
      {
        "name": "user"
      }
    ],
    "parameter": {
      "examples": [
        {
          "title": "response-example",
          "content": "\n{\n   \"message\": \"Successful\",\n   \"data\": {\n       \"events\": null,\n       \"organizations\": [\n           {\n               \"name\": \"CodeChef-VIT\",\n               \"location\": \"India\",\n               \"description\": \"Developer Student Clubs\",\n               \"tag\": \"technical\",\n               \"createdAt\": \"2019-08-05 23:12:14.963858896 +0000 UTC m=+14.778325532\",\n               \"website\": \"\"\n           }\n       ]\n   }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "organization/endpoints/get_handler.go",
    "groupTitle": "organization"
  },
  {
    "type": "get",
    "url": "/api/v1/org/events",
    "title": "Get events of organization",
    "name": "Get_events_of_organization",
    "group": "organization",
    "permission": [
      {
        "name": "user"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "org",
            "description": "<p>organization name as query param</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "curl localhost/api/v1/org/events?org=DSC-VIT",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "\n{\n   \"message\": \"Successful\",\n   \"data\": {\n       \"events\": [\n           {\n               \"clubName\": \"DSC-VIT\",\n               \"name\": \"WomenTechies'19\",\n               \"toDate\": \"11th March\",\n               \"fromDate\": \"10th March\",\n               \"toTime\": \"7 PM\",\n               \"fromTime\": \"6 PM\",\n               \"budget\": \"115000\",\n               \"description\": \"Women centric hackathon\",\n               \"category\": \"TECHNICAL\",\n               \"venue\": \"Kamaraj Auditorium\",\n               \"attendance\": \"315\",\n               \"expectedParticipants\": \"315\",\n               \"facultyCoordinator\": {\n                   \"name\": \"\",\n                   \"registrationNumber\": \"\",\n                   \"email\": \"\",\n                   \"phoneNumber\": \"\",\n                   \"gender\": \"\"\n               },\n               \"studentCoordinator\": {\n                   \"name\": \"\",\n                   \"registrationNumber\": \"\",\n                   \"email\": \"\",\n                   \"phoneNumber\": \"\",\n                   \"gender\": \"\"\n               },\n               \"PROrequest\": \"\",\n               \"campusEngineerRequest\": \"mics, podium, projector\",\n               \"duration\": \"24 hours\",\n               \"status\": \"\"\n           }\n       ]\n   }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "organization/endpoints/get_handler.go",
    "groupTitle": "organization"
  },
  {
    "type": "get",
    "url": "/api/v1/org/join",
    "title": "Get organization list",
    "name": "Get_organization_list",
    "group": "organization",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "org",
            "description": "<p>name of the organization as query param</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\ncurl localhost/api/v1/org/search?q=GDG-VIT",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "\n{\n   \"message\": \"Successful\",\n   \"data\": [\n       {\n           \"name\": \"GDG-VIT\",\n           \"location\": \"India\",\n           \"description\": \"Developer Student Clubs\",\n           \"tag\": \"technical\",\n           \"createdAt\": \"2019-08-05 21:49:11.875594638 +0000 UTC m=+30.409044543\",\n           \"website\": \"https://dsv-vit-vellore.com\"\n       }\n   ]\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "organization/endpoints/get_handler.go",
    "groupTitle": "organization"
  },
  {
    "type": "get",
    "url": "/api/v1/org/",
    "title": "Get organizations of user",
    "name": "Get_organizations_of_user",
    "group": "organization",
    "permission": [
      {
        "name": "user"
      }
    ],
    "parameter": {
      "examples": [
        {
          "title": "response-example",
          "content": "\n{\n   \"message\": \"Successful\",\n   \"data\": {\n       \"organizations\": [\n           {\n               \"name\": \"CodeChef-VIT\",\n               \"location\": \"India\",\n               \"description\": \"Developer Student Clubs\",\n               \"tag\": \"technical\",\n               \"createdAt\": \"2019-08-05 23:12:14.963858896 +0000 UTC m=+14.778325532\",\n               \"website\": \"\"\n           }\n       ]\n   }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "organization/endpoints/get_handler.go",
    "groupTitle": "organization"
  },
  {
    "type": "get",
    "url": "/api/v1/org/join",
    "title": "Send a join request to an org",
    "name": "Send_a_join_request_to_an_org",
    "group": "organization",
    "permission": [
      {
        "name": "user"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "org",
            "description": "<p>name of the organization as query param</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\ncurl -H '{\"Authorization\":\"dbasjbdasbdbasdjkbasjkda\"}' localhost/api/v1/org/join?org=GDG-VIT",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\"message\":\"Some error occurred\",\"data\":\"A join request is already pending\"}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "organization/endpoints/join_handler.go",
    "groupTitle": "organization"
  },
  {
    "type": "get",
    "url": "/api/v1/org/view-req",
    "title": "View join requests",
    "name": "View_join_requests",
    "group": "organization",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "org",
            "description": "<p>name of the organization as query param</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\ncurl localhost/api/v1/org/view-req?q=GDG-VIT -H '{\"Authorization\":\"asdbasbdbasjdbasjhbdhasd\"}'",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "\n{\n   \"message\": \"Successful\",\n   \"data\": [\n       {\n           \"firstName\": \"test\",\n           \"lastName\": \"test\",\n           \"password\": \"\",\n           \"email\": \"test1@test.com\",\n           \"phoneNumber\": \"\",\n           \"linkedIn\": \"test\",\n           \"facebook\": \"test\",\n           \"description\": \"test\",\n           \"createdAt\": \"20-01-01\"\n       },\n       {\n           \"firstName\": \"test\",\n           \"lastName\": \"test\",\n           \"password\": \"\",\n           \"email\": \"test1@test.com\",\n           \"phoneNumber\": \"\",\n           \"linkedIn\": \"test\",\n           \"facebook\": \"test\",\n           \"description\": \"test\",\n           \"createdAt\": \"20-01-01\"\n       },\n       {\n           \"firstName\": \"test\",\n           \"lastName\": \"test\",\n           \"password\": \"\",\n           \"email\": \"test1@test.com\",\n           \"phoneNumber\": \"\",\n           \"linkedIn\": \"test\",\n           \"facebook\": \"test\",\n           \"description\": \"test\",\n           \"createdAt\": \"20-01-01\"\n       }\n   ]\n}",
          "type": "json"
        }
      ]
    },
    "permission": [
      {
        "name": "admin"
      }
    ],
    "version": "0.0.0",
    "filename": "organization/endpoints/get_handler.go",
    "groupTitle": "organization"
  },
  {
    "type": "post",
    "url": "/api/v1/org/create-org",
    "title": "create an organization",
    "name": "create_an_organization",
    "group": "organization",
    "permission": [
      {
        "name": "user"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "name",
            "description": "<p>name of the org</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "location",
            "description": "<p>location of the org</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "description",
            "description": "<p>description of the org</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "tag",
            "description": "<p>tag of the org</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "website",
            "description": "<p>website of the org</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"name\":\"DSC-VIT\",\n\t\"location\":\"India\",\n\t\"description\":\"Developer Student Clubs\",\n\t\"tag\":\"technical\",\n\t\"website\":\"https://dsv-vit-vellore.com\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "organization/endpoints/auth_handler.go",
    "groupTitle": "organization"
  },
  {
    "type": "post",
    "url": "/api/v1/org/add-members",
    "title": "invite a user org",
    "name": "invite_a_user_to_an_org",
    "group": "organization",
    "permission": [
      {
        "name": "organization admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>email of the user</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "org",
            "description": "<p>name of the organization</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"name\":\"DSC-VIT\",\n\t\"location\":\"India\",\n\t\"description\":\"Developer Student Clubs\",\n\t\"tag\":\"technical\",\n\t\"website\":\"https://dsv-vit-vellore.com\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "organization/endpoints/auth_handler.go",
    "groupTitle": "organization"
  },
  {
    "type": "post",
    "url": "/api/v1/org/login",
    "title": "login as a user",
    "name": "login_as_a_user",
    "group": "organization",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "password",
            "description": "<p>password of the user</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>email of the user</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"email\":\"test1@test.com\",\n\t\"password\":\"test\"\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": \"Done\",\n   \"token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQHRlc3QuY29tIiwicm9sZSI6IkRFRkFVTFQiLCJvcmdhbml6YXRpb24iOiIifQ.3Qj3eu8iwXL2v1Rb7qGEf5USQ-WVjRvYiLALWIbWZ5c\",\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "organization/endpoints/auth_handler.go",
    "groupTitle": "organization"
  },
  {
    "type": "post",
    "url": "/api/v1/org/login-org",
    "title": "login to the org workspace (for privellage escalation)",
    "name": "login_to_the_org_workspace__for_privellage_escalation_",
    "group": "organization",
    "permission": [
      {
        "name": "organization member"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "password",
            "description": "<p>password of the user</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>email of the user</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"name\":\"GDGVIT\",\n\t\"role\":\"admin\"\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQHRlc3QuY29tIiwicm9sZSI6IkRFRkFVTFQiLCJvcmdhbml6YXRpb24iOiIifQ.3Qj3eu8iwXL2v1Rb7qGEf5USQ-WVjRvYiLALWIbWZ5c\",\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "organization/endpoints/auth_handler.go",
    "groupTitle": "organization"
  },
  {
    "type": "post",
    "url": "/api/v1/org/signup",
    "title": "signup as a user",
    "name": "signup_as_a_user",
    "group": "organization",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "firstName",
            "description": "<p>first name of the user</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "lastName",
            "description": "<p>last name of the user</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "password",
            "description": "<p>password of the user</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>email of the user</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "phoneNumber",
            "description": "<p>phoneNumber of the user</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "linkedIn",
            "description": "<p>linkedIn URL of the user</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "facebook",
            "description": "<p>facebook URL of the user</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "description",
            "description": "<p>description of the user</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "createdAt",
            "description": "<p>when was the user created</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"firstName\": \"test\",\n\t\"lastName\": \"test\",\n\t\"password\": \"test\",\n\t\"email\": \"test1@test.com\",\n\t\"phoneNumber\": \"998171818\",\n\t\"linkedIn\": \"test\",\n\t\"facebook\": \"test\",\n\t\"description\": \"test\",\n\t\"createdAt\": \"20-01-01\"\n\t\t\"deviceToken\": \"hasndbaskdjbsajd\"\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n    \"rs\": \"Done\",\n    \"token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQHRlc3QuY29tIiwicm9sZSI6IkRFRkFVTFQiLCJvcmdhbml6YXRpb24iOiIifQ.3Qj3eu8iwXL2v1Rb7qGEf5USQ-WVjRvYiLALWIbWZ5c\",\n    \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "organization/endpoints/auth_handler.go",
    "groupTitle": "organization"
  },
  {
    "type": "post",
    "url": "/api/v1/participants/create-attendee",
    "title": "create an attendee",
    "name": "create_an_attendee",
    "group": "participants",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>name of the participant</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "registrationNumber",
            "description": "<p>registration number of the participant</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email",
            "description": "<p>email of the participant</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "phoneNumber",
            "description": "<p>phoneNumber of the participant</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "gender",
            "description": "<p>gender of the participant</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "eventName",
            "description": "<p>name of the event registering for</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"details\":{\n     \"name\":\"angad sharma\",\n     \"registrationNumber\":\"17BBE1010\",\n     \"email\":\"SDADAS@A.COM\",\n     \"phoneNumber\":\"919191991911\",\n     \"gender\":\"M\",\n     \"eventsAttended\":\"ALL\",\n     \"eventName\":\"DEVSOC\"\n  }\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": \"created\",\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "participants/pkg/service/service.go",
    "groupTitle": "participants"
  },
  {
    "type": "delete",
    "url": "/api/v1/participants/delete-attendee",
    "title": "delete an attendee",
    "name": "delete_an_attendee",
    "group": "participants",
    "permission": [
      {
        "name": "super-admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "key",
            "description": "<p>key to query the attendee by</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "value",
            "description": "<p>value of the key</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"query\":{\n\t\t\"key\":\"name\",\n\t\t\"Value\":\"dhruv sharma\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": \"deleted\",\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "participants/pkg/service/service.go",
    "groupTitle": "participants"
  },
  {
    "type": "get",
    "url": "/api/v1/participants/read-attendee",
    "title": "read an attendee",
    "name": "read_an_attendee",
    "permission": [
      {
        "name": "super-admin"
      }
    ],
    "group": "participants",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "key",
            "description": "<p>key to query the attendee by</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "value",
            "description": "<p>value of the key</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"query\":{\n\t\t\"key\":\"name\",\n\t\t\"Value\":\"angad sharma\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": [\n       {\n           \"name\": \"angad sharma\",\n           \"registrationNumber\": \"17BBE1010\",\n           \"email\": \"SDADAS@A.COM\",\n           \"phoneNumber\": \"919191991911\",\n           \"gender\": \"M\",\n\t\t\t \"attended\":\"absent\"\n       }\n   ],\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "participants/pkg/service/service.go",
    "groupTitle": "participants"
  },
  {
    "type": "delete",
    "url": "/api/v1/participants/rm-attendee",
    "title": "remove attendee from an event",
    "name": "remove_attendee_from_an_event",
    "group": "participants",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "key",
            "description": "<p>key to query the attendee by</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "value",
            "description": "<p>value of the key</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "changeKey",
            "description": "<p>Name of the club</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "changeValue",
            "description": "<p>Name of the event</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"query\":{\n\t\t\"key\":\"name\",\n\t\t\"Value\":\"dhruv sharma\",\n\t\t\"changeKey\":\"DSC VIT\",\n\t\t\"changeValue\":\"DEVFEST 2019\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "\n{\n   \"rs\": \"updated\",\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "participants/pkg/service/service.go",
    "groupTitle": "participants"
  },
  {
    "type": "get",
    "url": "/api/v1/simple-projection/project-absent",
    "title": "show absent participants",
    "name": "show_absent_participants",
    "group": "simple_projection",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "key",
            "description": "<p>key to query the event by</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "value",
            "description": "<p>value of the key</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "event",
            "description": "<p>event name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "day",
            "description": "<p>day of the event</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"event\":\"DEVFEST 2019\",\n\t\"day\":2,\n\t\"query\":{\n\t\t\"key\":\"gender\",\n\t\t\"value\":\"F\",\n\t\"specific\" : \"DSC-VIT\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": [\n       {\n           \"name\": \"das\",\n           \"registrationNumber\": \"17BCE2009\",\n           \"email\": \"x@axd.com\",\n           \"phoneNumber\": \"919191991911\",\n           \"gender\": \"F\"\n       }\n   ],\n   \"err\": null\n}\n\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "simple_projection/pkg/service/service.go",
    "groupTitle": "simple_projection"
  },
  {
    "type": "get",
    "url": "/api/v1/simple-projection/project-all",
    "title": "show participants of an event",
    "name": "show_participants_of_an_event",
    "group": "simple_projection",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "key",
            "description": "<p>key to query the event by</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "value",
            "description": "<p>value of the key</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "event",
            "description": "<p>event name</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"event\":\"DEVFEST 2019\",\n\t\"query\":{\n\t\t\"key\":\"gender\",\n\t\t\"value\":\"F\",\n\t\"specific\" : \"DSC-VIT\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": [\n       {\n           \"name\": \"das\",\n           \"registrationNumber\": \"17BCE2009\",\n           \"email\": \"x@axd.com\",\n           \"phoneNumber\": \"919191991911\",\n           \"gender\": \"F\"\n       }\n   ],\n   \"err\": null\n}",
          "type": "json"
        },
        {
          "title": "request-example-all",
          "content": "{\n\t\"event\":\"DEVFEST 2019\"\n}",
          "type": "json"
        },
        {
          "title": "response-example-all",
          "content": "{\n   \"rs\": [\n       {\n           \"name\": \"das\",\n           \"registrationNumber\": \"17BCE2009\",\n           \"email\": \"x@axd.com\",\n           \"phoneNumber\": \"919191991911\",\n           \"gender\": \"F\"\n       },\n       {\n           \"name\": \"das\",\n           \"registrationNumber\": \"17BCE2009\",\n           \"email\": \"x@ax.com\",\n           \"phoneNumber\": \"919191991911\",\n           \"gender\": \"M\"\n       },\n       {\n           \"name\": \"das\",\n           \"registrationNumber\": \"17BCE2009\",\n           \"email\": \"x@x.com\",\n           \"phoneNumber\": \"919191991911\",\n           \"gender\": \"M\"\n       }\n   ],\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "simple_projection/pkg/service/service.go",
    "groupTitle": "simple_projection"
  },
  {
    "type": "get",
    "url": "/api/v1/simple-projection/project-present",
    "title": "show present participants",
    "name": "show_present_participants",
    "group": "simple_projection",
    "permission": [
      {
        "name": "admin"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "key",
            "description": "<p>key to query the event by</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "value",
            "description": "<p>value of the key</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "event",
            "description": "<p>event name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "day",
            "description": "<p>day of the event</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "{\n\t\"event\":\"DEVFEST 2019\",\n\t\"day\":2,\n\t\"query\":{\n\t\t\"key\":\"gender\",\n\t\t\"value\":\"F\",\n\t\"specific\" : \"DSC-VIT\"\n\t}\n}",
          "type": "json"
        },
        {
          "title": "response-example",
          "content": "{\n   \"rs\": [\n       {\n           \"name\": \"das\",\n           \"registrationNumber\": \"17BCE2009\",\n           \"email\": \"x@axd.com\",\n           \"phoneNumber\": \"919191991911\",\n           \"gender\": \"F\"\n       }\n   ],\n   \"err\": null\n}",
          "type": "json"
        },
        {
          "title": "request-example-all",
          "content": "{\n\t\"event\":\"DEVFEST 2019\",\n\t\"day\":2\n}",
          "type": "json"
        },
        {
          "title": "response-example-all",
          "content": "{\n   \"rs\": [\n       {\n           \"name\": \"das\",\n           \"registrationNumber\": \"17BCE2009\",\n           \"email\": \"x@axd.com\",\n           \"phoneNumber\": \"919191991911\",\n           \"gender\": \"F\"\n       },\n       {\n           \"name\": \"das\",\n           \"registrationNumber\": \"17BCE2009\",\n           \"email\": \"x@ax.com\",\n           \"phoneNumber\": \"919191991911\",\n           \"gender\": \"M\"\n       },\n       {\n           \"name\": \"das\",\n           \"registrationNumber\": \"17BCE2009\",\n           \"email\": \"x@x.com\",\n           \"phoneNumber\": \"919191991911\",\n           \"gender\": \"M\"\n       }\n   ],\n   \"err\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "simple_projection/pkg/service/service.go",
    "groupTitle": "simple_projection"
  }
] });
