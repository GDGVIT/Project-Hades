package CouponModule

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"log"
	"strconv"
	"sync"
)

func CreateCouponSchema(event string, coupons []Coupon, c chan error) {

	// check if coupon exists
	data, _, _, err := con.QueryNeoAll(`
		MATCH (n:EVENT)-[r:COUPON]->(:COUPON_SCHEMA)
		WHERE n.name = $event
		RETURN r.coupons
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

	// create schema; TODO error handling
	mutex := &sync.Mutex{}
	for _, cps := range coupons {

		go func(cp Coupon, mu *sync.Mutex) {

			mu.Lock()
			rs, err := con.ExecNeo(`
				MATCH(n:EVENT) WHERE n.name = $event
				CREATE (:COUPON_SCHEMA {name:$name, description:$desc, day:$day})<-[:COUPON]-(n)
			`, map[string]interface{}{
				"event": event,
				"name":  cp.Name,
				"desc":  cp.Desc,
				"day":   cp.Day,
			})
			mu.Unlock()

			if err != nil {
				log.Println(err)
				c <- err
				return
			}
			log.Println(rs)
		}(cps, mutex)

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
		`MATCH (n:EVENT)-[r:COUPON]->(a:COUPON_SCHEMA)
		 WHERE n.name = $event
		 RETURN a.name, a.description, a.day
	`, map[string]interface{}{
			"event": attendance.EventName,
		})
	if err != nil {
		c <- MessageReturn{"Error occurred while checking if coupon schema exists", err}
		return
	}

	str := fmt.Sprintf("%v", data)

	if str == "[[]]" && str == "[]" {
		// mark present if schema does not exist
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
		return
	}

	// if schema exists,generate hash, and save for each coupon
	var couponArr []Coupon

	for _, o := range data {
		if o[2].(int64) == int64(attendance.Day) {
			couponArr = append(couponArr, Coupon{
				Name: o[0].(string),
				Desc: o[1].(string),
				Day:  int(o[2].(int64)),
			})
		}
	}
	go couponGen(couponArr, attendance, c)

}

func returnStrs(arr []string) string {
	var buffer bytes.Buffer
	for _, r := range arr {
		buffer.WriteString(r + ",")
	}
	return buffer.String()[:len(buffer.String())]
}

// generate coupons and save
func couponGen(coupons []Coupon, attendance Attendance, c chan MessageReturn) {

	var (
		str string
		cps []string
	)

	for _, coupon := range coupons {
		str = attendance.EventName + strconv.Itoa(attendance.Day) + coupon.Name + attendance.Email
		bytes := md5.Sum([]byte(str))
		cps = append(cps, string(bytes[:]))
	}

	// create coupon relation
	_, err := con.ExecNeo(`
			MATCH(n:EVENT)-[:ATTENDS]->(b)
			WHERE n.name=$name AND b.email=$rn
			CREATE (n)-[:PRESENT`+strconv.Itoa(attendance.Day)+`]->(b) 
		`, map[string]interface{}{
		"name": attendance.EventName,
		"rn":   attendance.Email,
	})
	if err != nil {
		c <- MessageReturn{"Error creating coupon relation", err}
		return
	}
	for _, coup := range cps {
		_, err := con.ExecNeo(`
			MATCH(n:EVENT)-[:ATTENDS]->(b)
			WHERE n.name=$name AND b.email=$rn
			CREATE (b)-[:COUPON]->(c:COUPON {coupon:$cp})
		`, map[string]interface{}{
			"name": attendance.EventName,
			"rn":   attendance.Email,
			"cp":   coup,
		})
		if err != nil {
			c <- MessageReturn{"Error creating coupon relation", err}
			return
		}
	}
	c <- MessageReturn{"Successfully marked present for the day", nil}
	return
}

// redeem a coupon

func RedeemCoupon(attendance Attendance, couponName string, c chan MessageReturn) {

	// build coupon
	str := attendance.EventName + strconv.Itoa(attendance.Day) + couponName + attendance.Email
	bytes := md5.Sum([]byte(str)) //bcrypt.GenerateFromPassword([]byte(str), SALT)

	coupon := string(bytes[:])
	fmt.Println(coupon)
	// check if coupon exists
	data, _, _, err := con.QueryNeoAll(`
	MATCH (n:EVENT)-[r:PRESENT`+strconv.Itoa(attendance.Day)+`]->(a)-[:COUPON]->(c)
	WHERE a.email=$email AND c.coupon = "`+coupon+`"
	RETURN c.coupon
	`, map[string]interface{}{
		"email": attendance.Email,
	})

	if err != nil {
		c <- MessageReturn{"Error checking if coupon exists", err}
		return
	}

	strData := fmt.Sprintf("%v", data)
	log.Println(strData)
	if strData == "[[[]]]" || strData == "[]" {
		c <- MessageReturn{"No match found for this coupon", nil}
		return
	}

	// check if empty coupon node todo

	// remove from array
	_, err = con.ExecNeo(`
		MATCH (n:EVENT)-[b:PRESENT`+strconv.Itoa(attendance.Day)+`]->(a)-[:COUPON]->(c)
		WHERE a.email=$email AND n.name=$eventName AND c.coupon = "`+coupon+`"
		DETACH DELETE c
		`, map[string]interface{}{
		"eventName": attendance.EventName,
		"email":     attendance.Email,
	})

	if err != nil {
		c <- MessageReturn{"Some error occurred", err}
		return
	}

	c <- MessageReturn{"Successfully posted coupon", nil}
	return
}

func ViewCouponSchema(event string, mutex *sync.Mutex, c chan CouponReturn) {

	mutex.Lock()
	data, _, _, err := con.QueryNeoAll(`
		MATCH(n:EVENT)-[:COUPON]->(a:COUPON_SCHEMA)
		WHERE n.name=$event
		RETURN a.name, a.day, a.description
	`, map[string]interface{}{
		"event": event,
	})
	mutex.Unlock()
	if err != nil {
		c <- CouponReturn{nil, err}
		return
	}
	var cp []Coupon
	for i := range data {
		cp = append(cp, Coupon{
			Name: data[i][0].(string),
			Day:  int(data[i][1].(int64)),
			Desc: data[i][2].(string),
		})
	}
	c <- CouponReturn{cp, nil}
	return
}

func DeleteCouponSchema(event string, q Coupon, c chan MessageReturn) {

	var str string

	if q.Name != "" {

		str = `
			MATCH(n:EVENT)-[:COUPON]->(a:COUPON_SCHEMA)
			WHERE n.name=$event AND a.name=$cname AND a.day=$cday AND a.description = $cdesc
			DETACH DELETE a
		`
	} else {
		str = `
			MATCH(n:EVENT)-[:COUPON]->(a:COUPON_SCHEMA)
			WHERE n.name=$event
			DETACH DELETE a
		`
	}

	_, err := con.ExecNeo(str, map[string]interface{}{
		"event": event,
		"cname": q.Name,
		"cday":  q.Day,
		"cdesc": q.Desc,
	})
	if err != nil {
		c <- MessageReturn{"Error occurred while deleting", err}
		return
	}
	c <- MessageReturn{"Deleted", nil}
	return
}
