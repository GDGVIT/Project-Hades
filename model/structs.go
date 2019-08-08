package model

import (
	"reflect"

	jwt "github.com/dgrijalva/jwt-go"
)

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
	PROrequest            string      `json:"PROrequest"`
	CampusEngineerRequest string      `json:"campusEngineerRequest"`
	Duration              string      `json:"duration"`
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
	Key          string `json:"key"`
	Value        string `json:"value"`
	ChangeKey    string `json:"changeKey"`
	ChangeValue  string `json:"changeValue"`
	Organization string `json:"organization"`
	Specific     string `json:"specific"`
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

type CouponReturn struct {
	Coupons []Coupon `json:"coupons"`
	Err     error    `json:"err"`
}

type MessageReturn struct {
	Message string
	Err     error
}

type GuestReturn struct {
	Guests []Guest
	Err    error
}

// role as relation
type User struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	LinkedIn    string `json:"linkedIn"`
	Facebook    string `json:"facebook"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	DeviceToken string `json:"deviceToken"`
}

type UserReturn struct {
	User    User
	Err     error
	Message string
}

type Organization struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	CreatedAt   string `json:"createdAt"`
	Website     string `json:"website"`
}

type Token struct {
	Email        string `json:"email"`
	Role         string `json:"role"`
	Organization string `json:"organization"`
	jwt.StandardClaims
}

type TokenReturn struct {
	Token   string
	Err     error
	Message string
}

func (v Event) GetField(field string, value string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.FieldByName(value).String()
}
