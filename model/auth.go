package model

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"log"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/transport/http"
	"github.com/mitchellh/mapstructure"
)

func TokenGen(email string, role string, organization string, c chan TokenReturn) {
	tk := &Token{
		Email:        email,
		Role:         role,
		Organization: organization,
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_PASSWORD")))

	if err != nil {
		c <- TokenReturn{tokenString, err, "Some error occurred"}
		return
	}
	c <- TokenReturn{tokenString, nil, "Done"}
	return
}

func VerifyToken(ctx context.Context) (tk Token, err error) {

	token := ctx.Value(http.ContextKeyRequestAuthorization).(string)
	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.GetSigningMethod("HS256") {
			return nil, errors.New("Error, invalid signing method")
		}
		return []byte(os.Getenv("JWT_PASSWORD")), nil
	})

	if err != nil {
		return tk, err
	}

	if tok.Valid {
		mapstructure.Decode(tok.Claims, &tk)
		return tk, nil
	}
	return tk, errors.New("Invalid token")
}

func AuthorizeUser(ctx context.Context) (bool, error) {
	tk, err := VerifyToken(ctx)
	if err != nil {
		return false, err
	}
	if tk.Role != "admin" || tk.Role != "manager" || tk.Role != "member" {
		return false, nil
	}
	if !Enforce(tk.Email, tk.Organization, tk.Role) {
		return false, nil
	}
	return true, nil
}

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

func Login(email string, password string, organization string, role string) (token string, err error) {
	pwhash := md5.Sum([]byte(password))
	data, _, _, err := con.QueryNeoAll(
		`MATCH (u:USER) WHERE u.email=$email AND u.password=$password
		 RETURN u.firstName ,u.lastName, u.email, u.phoneNumber,
		 u.linkedIn, u.facebook, u.description, u.createdAt  
		`,
		map[string]interface{}{
			"email":    email,
			"password": string(pwhash[:]),
		})

	if err != nil {
		return "", err
	}

	str := fmt.Sprintf("%v", data)
	if str == "[]" || str == "[[]]" {
		return "", fmt.Errorf("Wrong email or password")
	}

	// generate and return token
	cc := make(chan TokenReturn)
	go TokenGen(email, role, organization, cc)
	tk := <-cc
	return tk.Token, tk.Err
}
