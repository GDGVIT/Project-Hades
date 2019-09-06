package endpoints

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/GDGVIT/Project-Hades/model"
	"github.com/GDGVIT/Project-Hades/organization/endpoints/views"
)

/**
* @api {post} /api/v1/org/login login as a user
* @apiName login as a user
* @apiGroup organization
*
* @apiParam {string} password password of the user
* @apiParam {string} email email of the user
*
*
* @apiParamExample {json} request-example
*{
*	"email":"test1@test.com",
*	"password":"test"
*}
*
* @apiParamExample {json} response-example
*{
*    "rs": "Done",
*    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQHRlc3QuY29tIiwicm9sZSI6IkRFRkFVTFQiLCJvcmdhbml6YXRpb24iOiIifQ.3Qj3eu8iwXL2v1Rb7qGEf5USQ-WVjRvYiLALWIbWZ5c",
*    "err": null
*}
**/
func login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		req := views.Auth{}
		json.NewDecoder(r.Body).Decode(&req)
		token, err := model.Login(req.Email, req.Password, "DEFAULT", "")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(views.Token{
				Message: "Some error occurred",
				Err:     err.Error(),
			})
			return
		}

		json.NewEncoder(w).Encode(views.Token{
			Token: token,
		})
	}
}

/**
* @api {post} /api/v1/org/create-org create an organization
* @apiName create an organization
* @apiGroup organization
*
* @apiPermission user
* @apiParam {string} name name of the org
* @apiParam {string} location location of the org
* @apiParam {string} description description of the org
* @apiParam {string} tag tag of the org
* @apiParam {string} website website of the org
*
* @apiParamExample {json} request-example
*
*{
*	"name":"DSC-VIT",
*	"location":"India",
*	"description":"Developer Student Clubs",
*	"tag":"technical",
*	"website":"https://dsv-vit-vellore.com"
*}
*
**/
func createOrg() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := model.Organization{}
		json.NewDecoder(r.Body).Decode(&data)
		token := r.Header.Get("Authorization")
		tk, err := model.ValidateToken(token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(views.Msg{"", err.Error()})
			return
		}
		if model.Enforce(tk.Email, data.Name, "member") || model.Enforce(tk.Email, data.Name, "admin") {
			json.NewEncoder(w).Encode(views.Msg{"Policy for this user already exists", nil})
			return
		}

		if data.Name == "" {
			json.NewEncoder(w).Encode(views.Msg{"Organization name needed", nil})
			return
		}
		if err = model.CreateNewOrg(data, tk.Email); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err.Error()})
			return
		}
		if err := model.AddPolicy(tk.Email, data.Name, "admin"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(views.Msg{"Error creating policy", err.Error()})
			return
		}

		if er := model.AddPolicy(tk.Email, data.Name, "member"); er != nil {
			json.NewEncoder(w).Encode(views.Msg{"Error creating policy", err.Error()})
			return
		}
		json.NewEncoder(w).Encode(views.Msg{"Created Organization", nil})
		return
	}
}

/**
* @api {post} /api/v1/org/login-org login to the org workspace (for privellage escalation)
* @apiName login to the org workspace (for privellage escalation)
* @apiGroup organization
* @apiPermission organization member
*
* @apiParam {string} password password of the user
* @apiParam {string} email email of the user
*
*
* @apiParamExample {json} request-example
*{
*	"name":"GDGVIT",
*	"role":"admin"
*}
*
* @apiParamExample {json} response-example
*{
*    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQHRlc3QuY29tIiwicm9sZSI6IkRFRkFVTFQiLCJvcmdhbml6YXRpb24iOiIifQ.3Qj3eu8iwXL2v1Rb7qGEf5USQ-WVjRvYiLALWIbWZ5c",
*    "err": null
*}
**/
func loginOrg() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := views.Role{}
		json.NewDecoder(r.Body).Decode(&data)
		token, err := model.ValidateToken(r.Header.Get("Authorization"))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(views.Msg{"Invalid token", err.Error()})
			return
		}
		if data.Role != "member" && data.Role != "admin" {
			json.NewEncoder(w).Encode(views.Msg{
				Message: "Invalid role. Only <member|admin> allowed",
				Data:    nil,
			})
			return
		}

		// check if user is authorized for the role
		if !model.Enforce(token.Email, data.Name, data.Role) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(views.Msg{"failed to authenticate user", errors.New("failed to authenticate user")})
			return
		}
		// generate and return token
		cc := make(chan model.TokenReturn)
		go model.TokenGen(token.Email, data.Role, data.Name, cc)
		tk := <-cc
		json.NewEncoder(w).Encode(views.Token{
			Token: tk.Token,
			Err:   tk.Err.Error(),
		})
		return
	}
}

/**
* @api {post} /api/v1/org/signup signup as a user
* @apiName signup as a user
* @apiGroup organization
*
* @apiParam {string} firstName first name of the user
* @apiParam {string} lastName last name of the user
* @apiParam {string} password password of the user
* @apiParam {string} email email of the user
* @apiParam {string} phoneNumber phoneNumber of the user
* @apiParam {string} linkedIn linkedIn URL of the user
* @apiParam {string} facebook facebook URL of the user
* @apiParam {string} linkedIn linkedIn URL of the user
* @apiParam {string} description description of the user
* @apiParam {string} createdAt when was the user created
*
*
* @apiParamExample {json} request-example
*
* {
* 	"firstName": "test",
* 	"lastName": "test",
* 	"password": "test",
* 	"email": "test1@test.com",
* 	"phoneNumber": "998171818",
* 	"linkedIn": "test",
* 	"facebook": "test",
* 	"description": "test",
* 	"createdAt": "20-01-01"
*		"deviceToken": "hasndbaskdjbsajd"
* }
*
*
*
* @apiParamExample {json} response-example
* {
*     "rs": "Done",
*     "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQHRlc3QuY29tIiwicm9sZSI6IkRFRkFVTFQiLCJvcmdhbml6YXRpb24iOiIifQ.3Qj3eu8iwXL2v1Rb7qGEf5USQ-WVjRvYiLALWIbWZ5c",
*     "err": null
* }
**/
func signup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := model.User{}
		fmt.Println("Signup")
		json.NewDecoder(r.Body).Decode(&user)
		if user.Email == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(views.Token{
				Message: "User email not found",
				Err:     "User email not found",
			})
			return
		}
		c := make(chan model.UserReturn)
		go user.Get(c)
		msg := <-c
		close(c)
		if err := msg.Err; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(views.Token{
				Message: msg.Message,
				Err:     err.Error(),
			})
			return
		} else if msg.User.Email == "" {
			json.NewEncoder(w).Encode(views.Token{
				Message: msg.Message,
			})
			return
		}

		// generate JWT
		cc := make(chan model.TokenReturn)
		go model.TokenGen(user.Email, "DEFAULT", "", cc)
		msg2 := <-cc
		close(cc)
		if err := msg2.Err; err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(views.Token{
				Message: msg2.Message,
				Err:     err.Error(),
			})
			return
		}
		json.NewEncoder(w).Encode(views.Token{
			Message: msg2.Message,
			Token:   msg2.Token,
		})
		return
	}
}

/**
* @api {post} /api/v1/org/add-members invite a user org
* @apiName invite a user to an org
* @apiGroup organization
* @apiPermission organization admin
*
* @apiParam {string} email email of the user
* @apiParam {string} org name of the organization
*
* @apiParamExample {json} request-example
*
*{
*	"name":"DSC-VIT",
*	"location":"India",
*	"description":"Developer Student Clubs",
*	"tag":"technical",
*	"website":"https://dsv-vit-vellore.com"
*}
 */
func addMembers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := views.AddMembers{}
		json.NewDecoder(r.Body).Decode(&data)
		tk, err := model.ValidateToken(r.Header.Get("Authorization"))
		if err != nil || !model.Enforce(tk.Email, tk.Organization, "admin") {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(views.Msg{"error authorizing user", err.Error()})
			return
		}
		if err := model.InviteUserToOrg(data.Email, data.Org); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(views.Msg{"", err.Error()})
			return
		}
		if err := model.AddPolicy(data.Email, data.Org, "member"); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(views.Msg{"error adding policy", err.Error()})
			return
		}

		json.NewEncoder(w).Encode(views.Msg{"successful", nil})
		return
	}
}
