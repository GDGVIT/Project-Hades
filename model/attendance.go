package model

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func MarkPresent(query Attendance, c chan MessageReturn) {

	// check if user exists or not
	// check if already given attendance
	data, _, _, err := con.QueryNeoAll(`
		MATCH(n:EVENT)-[r:PRESENT`+strconv.Itoa(query.Day)+`]->(b)
		WHERE n.name=$name AND b.email=$rn
		RETURN b.email
	`, map[string]interface{}{
		"name": query.EventName,
		"rn":   query.Email,
	})
	if err != nil {
		c <- MessageReturn{"Error marking attendance", err}
		return
	}
	if len(data) > 0 {
		c <- MessageReturn{"Already marked present", nil}
		return
	}

	// if number of coupons = 0 just create a PRESENT relation
	if query.Coupons == 0 {
		_, err := con.ExecNeo(`
			MATCH(n:EVENT)-[:ATTENDS]->(b)
			WHERE n.name=$name AND b.email=$rn
			CREATE (n)-[:PRESENT`+strconv.Itoa(query.Day)+`]->(b) 
		`, map[string]interface{}{
			"name": query.EventName,
			"rn":   query.Email,
		})
		if err != nil {
			c <- MessageReturn{"Error creating present relation", err}
			return
		}
		c <- MessageReturn{"Successfully marked present for the day", nil}
		return
	}

	// goroutine for generating coupon hashes
	cc := make(chan []string)
	go couponGen(query.EventName, query.Email, query.Coupons, cc)

	// generate a hashed array of coupons, mark user present and add coupons in the relation
	couponArr := <-cc
	_, err = con.ExecNeo(`
			MATCH(n:EVENT)-[:ATTENDS]->(b)
			WHERE n.name=$name AND b.email=$rn
			CREATE (n)-[:PRESENT`+strconv.Itoa(query.Day)+`]->(b) 
			CREATE (b)-[:COUPON_`+strings.Replace(query.EventName, " ", "", -1)+strconv.Itoa(query.Day)+`]->(c:COUPONS {coupons:$cps})
		`, map[string]interface{}{
		"name": query.EventName,
		"rn":   query.Email,
		"cps":  couponArr,
	})
	if err != nil {
		c <- MessageReturn{"Error creating coupon node", err}
		return
	}
	c <- MessageReturn{"Successfully marked present for the day", nil}
	return

}

func couponGen(eventName string, email string, coupons int, cc chan []string) {
	var couponArr []string
	SALT, _ := strconv.Atoi(os.Getenv("SALT"))
	for i := 0; i < coupons; i++ {

		bytes, err := bcrypt.GenerateFromPassword([]byte(eventName+strconv.Itoa(coupons)+email), SALT)
		if err != nil {
			log.Println("Error while hashing")
			cc <- couponArr
			return
		}
		couponArr = append(couponArr, string(bytes))
	}
	cc <- couponArr
	return
}

func ViewCoupon(query Attendance) []string {

	data, _, _, err := con.QueryNeoAll(`
		MATCH(n:EVENT)-[:PRESENT`+strconv.Itoa(query.Day)+`]->(b)-[:COUPON_`+strings.Replace(query.EventName, " ", "", -1)+strconv.Itoa(query.Day)+`]->(c)
		WHERE b.email=$rn
		RETURN c.coupons`, map[string]interface{}{
		"rn": query.Email,
	})
	if err != nil {
		log.Println(err)
		var arr []string
		return arr
	}
	if len(data) < 1 {
		var arr []string
		return arr
	}

	str := fmt.Sprintf("%v", data[0][0])
	str = strings.Replace(str, "[", "", -1)
	str = strings.Replace(str, "]", "", -1)
	return strings.Split(str, " ")

}

func PostCoupon(coupon string, query Attendance, c chan MessageReturn) {

	// check if coupon exists
	data, _, _, err := con.QueryNeoAll(`
	MATCH (n:EVENT)-[:PRESENT`+strconv.Itoa(query.Day)+`]->(a)-[:COUPON_`+strings.Replace(query.EventName, " ", "", -1)+strconv.Itoa(query.Day)+`]->(c)
	WHERE a.email=$email
	RETURN [x IN c.coupons WHERE x = $coupon];
	`, map[string]interface{}{
		"email":  query.Email,
		"coupon": coupon,
	})

	if err != nil {
		c <- MessageReturn{"Error checking if coupon exists", err}
		return
	}

	str := fmt.Sprintf("%v", data)

	if str == "[[[]]]" || str == "[]" {
		c <- MessageReturn{"No match found for this coupon", nil}
		return
	}

	// check if empty coupon node
	cp := ViewCoupon(query)
	if len(cp) < 1 || cp[0] == "" {
		ce := make(chan MessageReturn)
		go DeleteCoupons(query, ce)

		msg := <-ce
		if err = msg.Err; err != nil {
			c <- msg
			return
		}

		c <- MessageReturn{"No more coupons exist for this user", nil}
		return

	}

	// remove from array
	_, err = con.ExecNeo(`
		MATCH (n:EVENT)-[:PRESENT`+strconv.Itoa(query.Day)+`]->(a)-[:COUPON_`+strings.Replace(query.EventName, " ", "", -1)+strconv.Itoa(query.Day)+`]->(c)
		WHERE a.email=$email
		SET c.coupons=[x IN c.coupons WHERE x <> $coupon];
		`, map[string]interface{}{
		"eventName": query.EventName,
		"email":     query.Email,
		"coupon":    coupon,
	})

	if err != nil {
		c <- MessageReturn{"Some error occurred", err}
		return
	}

	// check if empty node
	cp = ViewCoupon(query)
	if cp[0] == "" {
		ce := make(chan MessageReturn)
		go DeleteCoupons(query, ce)

		msg := <-ce
		if err := msg.Err; err != nil {
			c <- msg
			return
		}

		c <- MessageReturn{"No more coupons exist for this user", nil}
		return

	}

	c <- MessageReturn{"Successfully posted coupon", nil}
	return

}

func DeleteCoupons(query Attendance, c chan MessageReturn) {
	_, err := con.ExecNeo(`
	MATCH (n:EVENT)-[:PRESENT`+strconv.Itoa(query.Day)+`]->(a)-[:COUPON_`+strings.Replace(query.EventName, " ", "", -1)+strconv.Itoa(query.Day)+`]->(c)
	WHERE a.email=$email
	DETACH DELETE c
	`, map[string]interface{}{
		"email": query.Email,
	})

	if err != nil {
		c <- MessageReturn{"Some error occurred while deleting node", err}
		return
	}

	c <- MessageReturn{"Successfully deleted coupon node", nil}
	return
}
