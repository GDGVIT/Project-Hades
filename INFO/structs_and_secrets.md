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
	FacultyCoordinator    Participant `json:"facultyCoordinator"`
	StudentCoordinator    Participant `json:"studentCoordinator"`
	GuestDetails          Guest       `json:"guest"`
	PROrequest            string      `json:"PROrequest"`
	CampusEngineerRequest string      `json:"campusEngineerRequest"`
	Duration              string      `json:"duration"`
	MainSponsor           Participant `json:"mainSponsor"`
	Status                string      `json:"status"`
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

#### Attendee schema

<br />

```go
{
	Name               string `json:"name"`
	RegistrationNumber string `json:"registrationNumber"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phoneNumber"`
	Gender             string `json:"gender"`
	EventName          string `json:"eventName"`
}
```

<br />

#### .env environment variable file

Place .nv file with the following details in each of the microservice folders

```
URI=
PROD_URI=
GRAPHENE_URI=
SALT=
PROJECTION_URI=http://<IP of machine>/api/v1/simple-projection/
MAIL_FROM=
MAIL_PASSWORD=
MAIL_TO=
JWT_PASSWORD=
```

<br />
