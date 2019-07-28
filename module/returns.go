package model

import "reflect"

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

type UserReturn struct {
	User    User
	Err     error
	Message string
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
