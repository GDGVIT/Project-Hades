package model

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func CreateCouponSchema(event string, coupons []Coupon, c chan error) {

	// check if coupon exists
	data, _, _, err := con.QueryNeoAll(
		`MATCH (n:EVENT)-[r:COUPON]->(:COUPON_SCHEMA)
		 WHERE n.name = $event
		 RETURN r
	`, map[string]interface{}{
			"event": event,
		})
	if err != nil {
		c <- err
		return
	}

	str := fmt.Sprintf("%v", data)
	if str == "[[]]" && str == "[]" {
		c <- fmt.Errorf("already exists")
		return
	}

	// create schema
	_, err = con.ExecNeo(`
		CREATE (:COUPON_SCHEMA { coupons: $coupons })<-[:COUPON]-(n:EVENT)
		WHERE n.event = $event
	`, map[string]interface{}{
		"event":   event,
		"coupons": coupons,
	})

	if err != nil {
		c <- err
		return
	}

	c <- nil
	return

}

func MarkPresent(attendance Attendance, c chan MessageReturn) {

	// check if user exists or not
	data, _, _, err := con.QueryNeoAll(`
		MATCH(n:EVENT)-[:ATTENDS]->(b)
		WHERE n.name=$name AND b.email=$rn
		RETURN b.email
	`, map[string]interface{}{
		"name": attendance.EventName,
		"rn":   attendance.Email,
	})
	if err != nil {
		c <- MessageReturn{"Error marking attendance", err}
		return
	}

	if len(data) < 1 {
		c <- MessageReturn{"No participant found", nil}
		return
	}

	// check if already given attendance
	data, _, _, err = con.QueryNeoAll(`
		MATCH(n:EVENT)-[r:PRESENT`+strconv.Itoa(attendance.Day)+`]->(b)
		WHERE n.name=$name AND b.email=$rn
		RETURN b.email
	`, map[string]interface{}{
		"name": attendance.EventName,
		"rn":   attendance.Email,
	})
	if err != nil {
		c <- MessageReturn{"Error marking attendance", err}
		return
	}
	if len(data) > 0 {
		c <- MessageReturn{"Already marked present", nil}
		return
	}

	// check if schema exists
	data, _, _, err = con.QueryNeoAll(
		`MATCH (n:EVENT)-[r:COUPON]->(:COUPON_SCHEMA)
		 WHERE n.name = $event
		 RETURN r
	`, map[string]interface{}{
			"event": attendance.EventName,
		})
	if err != nil {
		c <- MessageReturn{"Error occurred while checking if coupon schema exists", err}
		return
	}

	str := fmt.Sprintf("%v", data)
	if str != "[[]]" && str != "[]" {
		// mark present
		_, err = con.ExecNeo(`
			MATCH(n:EVENT)-[:ATTENDS]->(b)
			WHERE n.name=$name AND b.email=$rn
			CREATE (n)-[:PRESENT`+strconv.Itoa(attendance.Day)+`]->(b) 
		`, map[string]interface{}{
			"name": attendance.EventName,
			"rn":   attendance.Email,
		})
		if err != nil {
			c <- MessageReturn{"Error creating present relation", err}
			return
		}
		c <- MessageReturn{"Done", nil}

	}

	// if schema exists,generate hash, and save for each coupon
	var (
		couponArr []Coupon
		cps       []Coupon
	)
	cps = data[0][0].([]Coupon)

	for _, cp := range cps {
		if cp.Day == attendance.Day {
			couponArr = append(couponArr)
		}
	}
	go couponGen(couponArr, attendance, c)

}

// generate coupons and save
func couponGen(coupons []Coupon, attendance Attendance, c chan MessageReturn) {

	var (
		str string
		cps []string
	)
	SALT, _ := strconv.Atoi(os.Getenv("SALT"))

	for _, coupon := range coupons {

		str = attendance.EventName + strconv.Itoa(attendance.Day) + coupon.Name + attendance.Email

		bytes, err := bcrypt.GenerateFromPassword([]byte(str), SALT)
		if err != nil {
			c <- MessageReturn{"Error while hashing", err}
			return
		}
		cps = append(cps, string(bytes))
	}

	// create coupon relation
	_, err := con.ExecNeo(`
			MATCH(n:EVENT)-[:ATTENDS]->(b)
			WHERE n.name=$name AND b.email=$rn
			CREATE (n)-[:PRESENT`+strconv.Itoa(attendance.Day)+`{coupons:$cps}]->(b) 
		`, map[string]interface{}{
		"name": attendance.EventName,
		"rn":   attendance.Email,
		"cps":  cps,
	})
	if err != nil {
		c <- MessageReturn{"Error creating coupon relation", err}
		return
	}
	c <- MessageReturn{"Successfully marked present for the day", nil}
	return
}

// redeem a coupon

func RedeemCoupon(attendance Attendance, couponName string, c chan MessageReturn) {

	// build coupon
	SALT, _ := strconv.Atoi(os.Getenv("SALT"))
	str := attendance.EventName + strconv.Itoa(attendance.Day) + couponName + attendance.Email
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), SALT)
	if err != nil {
		c <- MessageReturn{"Error while hashing", err}
		return
	}
	coupon := string(bytes)

	// check if coupon exists
	data, _, _, err := con.QueryNeoAll(`
	MATCH (n:EVENT)-[r:PRESENT`+strconv.Itoa(attendance.Day)+`]->(a)
	WHERE a.email=$email
	RETURN [x IN r.coupons WHERE x = $coupon];
	`, map[string]interface{}{
		"email":  attendance.Email,
		"coupon": coupon,
	})

	if err != nil {
		c <- MessageReturn{"Error checking if coupon exists", err}
		return
	}

	str = fmt.Sprintf("%v", data)

	if str == "[[[]]]" || str == "[]" {
		c <- MessageReturn{"No match found for this coupon", nil}
		return
	}

	// check if empty coupon node
	// cp := ViewCoupon(attendance)
	// if len(cp) < 1 || cp[0] == "" {
	// 	ce := make(chan MessageReturn)
	// 	go DeleteCoupons(attendance, ce)

	// 	msg := <-ce
	// 	if err = msg.Err; err != nil {
	// 		c <- msg
	// 		return
	// 	}

	// 	c <- MessageReturn{"No more coupons exist for this user", nil}
	// 	return

	// }

	// remove from array
	_, err = con.ExecNeo(`
		MATCH (n:EVENT)-[c:PRESENT`+strconv.Itoa(attendance.Day)+`]->(a)
		WHERE a.email=$email AND n.name=$eventName
		SET c.coupons=[x IN c.coupons WHERE x <> $coupon];
		`, map[string]interface{}{
		"eventName": attendance.EventName,
		"email":     attendance.Email,
		"coupon":    coupon,
	})

	if err != nil {
		c <- MessageReturn{"Some error occurred", err}
		return
	}

	// check if empty node
	// cp = ViewCoupon(attendance)
	// if cp[0] == "" {
	// 	ce := make(chan MessageReturn)
	// 	go DeleteCoupons(attendance, ce)

	// 	msg := <-ce
	// 	if err := msg.Err; err != nil {
	// 		c <- msg
	// 		return
	// 	}

	// 	c <- MessageReturn{"No more coupons exist for this user", nil}
	// 	return

	// }

	c <- MessageReturn{"Successfully posted coupon", nil}
	return
}
