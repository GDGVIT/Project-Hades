# Project hades
The last event coordinator you will ever need

<br />


### Steps to run

<br />

```
docker-compose up
```

<br />

#### Event schema

<br />

```go
{
	ClubName              string      `json:"clubName"`
	Name                  string      `json:"name"`
	ToDate                string      `json:"toDate"`
	FromDate              string      `json:"fromDate"`
	ToTime                string      `json:"toTime"`
	FromTime              string      `json:"fromTime"`
	Budget                string      `json:"budget"`
	Description           string      `json:"description"`
	Category              string      `json:"category"`
	Venue                 string      `json:"venue"`
	Attendance            string      `json:"attendance"`
	ExpectedParticipants  string      `json:"expectedParticipants"`
	FacultyCoordinator    Participant `json:"sacultyCoordinator"`
	StudentCoordinator    Participant `json:"studentCoordinator"`
	GuestDetails          Guest       `json:"guest"`
	PROrequest            string      `json:"PROrequest"`
	CampusEngineerRequest string      `json:"campusEngineerRequest"`
	Duration              string      `json:"duration"`
}

```

<br />

#### Guest schema

<br />

```go
{
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phoneNumber"`
	Gender         string `json:"gender"`
	Stake          string `json:"stake"`
	LocationOfStay string `json:"locationOfStay"`
}
```


<br />

#### Participant schema

<br />

```go
{
	Name               string `json:"name"`
	RegistrationNumber string `json:"registrationNumber"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phoneNumber"`
	Gender             string `json:"gender"`
}
```

<br />

#### Secret file
model/creds.go

```go
package model

type secret struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_ENDPOINT string
}

var DB_SECRET = secret{

	DB_USERNAME: "your DB username",
	DB_PASSWORD: "your DB password",
	DB_ENDPOINT: "your DB endpoint URI eg: localhost:7687",
}

```

<br />
