package model

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func MarkPresent(eventName string, email string, coupons int, day int, c chan error) {

	// check if user exists or not
	// check if already given attendance
	data, _, _, err := con.QueryNeoAll(`
		MATCH(n:EVENT)-[r:PRESENT`+strconv.Itoa(day)+`]->(b)
		WHERE n.name=$name AND b.email=$rn
		RETURN b.email
	`, map[string]interface{}{
		"name": eventName,
		"rn":   email,
	})
	if err != nil {
		c <- err
		return
	}
	if len(data) > 0 {
		c <- fmt.Errorf("Already marked present")
		return
	}

	// if number of coupons = 0 just create a PRESENT relation
	if coupons == 0 {
		_, err := con.ExecNeo(`
			MATCH(n:EVENT)-[:ATTENDS]->(b)
			WHERE n.name=$name AND b.email=$rn
			CREATE (n)-[:PRESENT`+strconv.Itoa(day)+`]->(b) 
		`, map[string]interface{}{
			"name": eventName,
			"rn":   email,
		})
		if err != nil {
			c <- err
			return
		}
		c <- nil
		return
	}

	// goroutine for generating coupon hashes
	cc := make(chan []string)
	go couponGen(eventName, email, coupons, cc)

	// generate a hashed array of coupons, mark user present and add coupons in the relation
	couponArr := <-cc
	_, err = con.ExecNeo(`
			MATCH(n:EVENT)-[:ATTENDS]->(b)
			WHERE n.name=$name AND b.email=$rn
			CREATE (n)-[:PRESENT`+strconv.Itoa(day)+`]->(b) 
			CREATE (b)-[:COUPON_`+strings.Replace(eventName, " ", "", -1)+strconv.Itoa(day)+`]->(c:COUPONS {coupons:$cps})
		`, map[string]interface{}{
		"name": eventName,
		"rn":   email,
		"cps":  couponArr,
	})
	if err != nil {
		c <- err
		return
	}
	c <- nil
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
		MATCH(n:EVENT)-[:PRESENT`+strconv.Itoa(query.Day)+`]->(b)
		WHERE n.name=$en AND b.email=$rn
		MATCH (b)-[:COUPON_`+strings.Replace(query.EventName, " ", "", -1)+strconv.Itoa(query.Day)+`]->(c)
		RETURN c.coupons`, map[string]interface{}{

		"en": query.EventName,
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
