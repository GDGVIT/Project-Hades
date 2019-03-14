package model

import (
	"crypto/md5"
	"log"
)

func (u *User) Get(c chan UserReturn) {
	var user User
	log.Println(u)
	// check if user exists
	data, _, _, err := con.QueryNeoAll(
		`MATCH (u:USER) WHERE u.email=$email
		 RETURN u.firstName ,u.lastName, u.email, u.phoneNumber,
		 u.linkedIn, u.facebook, u.description, u.createdAt  
		`,
		map[string]interface{}{
			"email": u.Email,
		})

	if err != nil {
		c <- UserReturn{user, err, "Some error occurred"}
		return
	}

	if len(data) < 1 {

		// if not, hash password and save user
		pwhash := md5.Sum([]byte(u.Password))
		_, err := con.ExecNeo(`
			CREATE (u:USER {firstName:$1, lastname:$2, email:$3, phoneNumber:$4,
			linkedIn:$5, facebook:$6, description:$7, createdAt:$8, password:$9 })
		`, map[string]interface{}{
			"1": u.FirstName,
			"2": u.LastName,
			"3": u.Email,
			"4": u.PhoneNumber,
			"5": u.LinkedIn,
			"6": u.Facebook,
			"7": u.Description,
			"8": u.CreatedAt,
			"9": string(pwhash[:]),
		})

		if err != nil {
			c <- UserReturn{user, err, "Some error occurred"}
			return
		}
		c <- UserReturn{*u, nil, "Created new user"}
		return

	} else {
		c <- UserReturn{user, nil, "User already exists"}
		return
	}
}
