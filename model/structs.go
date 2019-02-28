package model

import "reflect"

type Participant struct {
	Name               string `json:"name"`
	RegistrationNumber string `json:"registrationNumber"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phoneNumber"`
	Gender             string `json:"gender"`
}

type Attendee struct {
	Name               string `json:"name"`
	RegistrationNumber string `json:"registrationNumber"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phoneNumber"`
	Gender             string `json:"gender"`
	EventName          string `json:"eventName"`
}

type Event struct {
	ClubName             string      `json:"clubName"`
	Name                 string      `json:"name"`
	ToDate               string      `json:"toDate"`
	FromDate             string      `json:"fromDate"`
	ToTime               string      `json:"toTime"`
	FromTime             string      `json:"fromTime"`
	Budget               string      `json:"budget"`
	Description          string      `json:"description"`
	Category             string      `json:"category"`
	Venue                string      `json:"venue"`
	Attendance           string      `json:"attendance"`
	ExpectedParticipants string      `json:"expectedParticipants"`
	FacultyCoordinator   Participant `json:"facultyCoordinator"`
	StudentCoordinator   Participant `json:"studentCoordinator"`
	// GuestDetails          Guest       `json:"guest"`
	PROrequest            string `json:"PROrequest"`
	CampusEngineerRequest string `json:"campusEngineerRequest"`
	Duration              string `json:"duration"`
	// MainSponsor           Participant `json:"mainSponsor"`
	Status string `json:"status"`
}

type Guest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phoneNumber"`
	Gender         string `json:"gender"`
	Stake          string `json:"stake"`
	LocationOfStay string `json:"locationOfStay"`
}

type Query struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	ChangeKey   string `json:"changeKey"`
	ChangeValue string `json:"changeValue"`
	Specific    string `json:"specific"`
}

type EventReturn struct {
	Event []Event `json:"event"`
	Err   error   `json:"err"`
}

type ParticipantReturn struct {
	Attendees []Attendee `json:"attendee"`
	Err       error      `json:"err"`
}
type SafeParticipantReturn struct {
	Participants []Participant `json:"participants"`
	Err          error         `json:"err"`
}

type Attendance struct {
	EventName  string `json:"eventName"`
	Email      string `json:"email"`
	Day        int    `json:"day"`
	CouponName string `json:"couponName"`
}

type Coupon struct {
	Name string `json:"name"`
	Desc string `json:"description"`
	Day  int    `json:"day"`
}

type MessageReturn struct {
	Message string
	Err     error
}

type GuestReturn struct {
	Guests []Guest
	Err    error
}

func (v Event) GetField(field string, value string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.FieldByName(value).String()
}
