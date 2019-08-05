package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GDGVIT/Project-Hades/model"
	"github.com/GDGVIT/Project-Hades/organization/endpoints/views"
)

func login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := views.Auth{}
		json.NewDecoder(r.Body).Decode(&req)
		token, err := model.Login(req.Email, req.Password, "DEFAULT", "")
		if err != nil {
			json.NewEncoder(w).Encode(views.Token{
				Message: "Some error occurred",
				Err:     err,
			})
		}
		json.NewEncoder(w).Encode(views.Token{
			Token: token,
		})
	}
}

func createOrg() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := model.Organization{}
		json.NewDecoder(r.Body).Decode(&data)
		token := r.Header.Get("Authorization")
		tk, err := model.ValidateToken(token)
		if err != nil {
			json.NewEncoder(w).Encode(views.Msg{"", err.Error()})
			return
		}

		if data.Name == "" {
			json.NewEncoder(w).Encode(views.Msg{"Organization name needed", nil})
			return
		}
		if err = model.CreateNewOrg(data); err != nil {
			json.NewEncoder(w).Encode(views.Msg{"", err})
			return
		}
		if err := model.AddPolicy(tk.Email, data.Name, "admin"); err != nil {
			json.NewEncoder(w).Encode(views.Msg{"Error creating policy", err})
			return
		}
		if er := model.AddPolicy(tk.Email, data.Name, "member"); er != nil {
			json.NewEncoder(w).Encode(views.Msg{"Error creating policy", err})
			return
		}
		json.NewEncoder(w).Encode(views.Msg{"Created Organization", nil})
		return
	}
}

func loginOrg() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := views.Role{}
		json.NewDecoder(r.Body).Decode(&data)
		token, err := model.ValidateToken(r.Header.Get("Authorization"))
		if err != nil {
			json.NewEncoder(w).Encode(views.Msg{"Invalid token", err})
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
			json.NewEncoder(w).Encode(views.Msg{"failed to authenticate user", nil})
			return
		}
		// generate and return token
		cc := make(chan model.TokenReturn)
		go model.TokenGen(token.Email, data.Role, data.Name, cc)
		tk := <-cc
		json.NewEncoder(w).Encode(views.Token{
			Token: tk.Token,
			Err:   tk.Err,
		})
		return
	}
}

func signup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := model.User{}
		fmt.Println("Signup")
		json.NewDecoder(r.Body).Decode(&user)
		if user.Email == "" {
			json.NewEncoder(w).Encode(views.Token{
				Message: "User email not found",
			})
			return
		}
		c := make(chan model.UserReturn)
		go user.Get(c)
		msg := <-c
		close(c)
		if err := msg.Err; err != nil {
			json.NewEncoder(w).Encode(views.Token{
				Message: msg.Message,
				Err:     err,
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
			json.NewEncoder(w).Encode(views.Token{
				Message: msg2.Message,
				Err:     err,
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

func addMembers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := views.AddMembers{}
		json.NewDecoder(r.Body).Decode(&data)
		tk, err := model.ValidateToken(r.Header.Get("Authorization"))
		if err != nil || !model.Enforce(tk.Email, tk.Organization, "admin") {
			json.NewEncoder(w).Encode(views.Msg{"error authorizing user", err})
			return
		}
		if err := model.InviteUserToOrg(data.Email, data.Org); err != nil {
			json.NewEncoder(w).Encode(views.Msg{"", err})
			return
		}
		if err := model.AddPolicy(data.Email, data.Org, "member"); err != nil {
			json.NewEncoder(w).Encode(views.Msg{"error adding policy", err})
			return
		}

		json.NewEncoder(w).Encode(views.Msg{"successful", nil})
		return
	}
}
