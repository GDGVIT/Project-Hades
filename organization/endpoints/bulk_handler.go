package endpoints

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"

	"github.com/GDGVIT/Project-Hades/model"
	"github.com/GDGVIT/Project-Hades/organization/endpoints/views"
)

func bulkAddAttendees() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		tk, err := model.ValidateToken(token)
		if err != nil {
			json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err.Error()})
			return
		}
		data := struct {
			Attendees []model.Participant `json:"attendees"`
			EventName string              `json:"eventName"`
		}{}

		json.NewDecoder(r.Body).Decode(&data)

		access, err := model.EnforceRoleAdmin(tk.Email, tk.Organization)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err.Error()})
			return
		}

		if !access {
			json.NewEncoder(w).Encode(views.Msg{"failed to authenticate user", errors.New("failed to authenticate user")})
			return
		}

		isEventOfOrg, err := model.IsEventOfOrg(data.EventName, tk.Organization)
		if err != nil {
			json.NewEncoder(w).Encode(views.Msg{"Error: Some error occurred while checking if the event belongs to this org", nil})
			return
		}
		if !isEventOfOrg {
			json.NewEncoder(w).Encode(views.Msg{"Error: The event does not belong to this organization", nil})
			return
		}

		mutex := &sync.Mutex{}
		errs := []error{}

		for _, i := range data.Attendees {
			if err := model.CreateAttendeeSync(data.EventName, i, mutex); err != nil {
				errs = append(errs, err)
			}
		}

		if len(errs) < 1 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(struct {
				Err []error `json:"error"`
			}{errs})
			return
		} else {
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}
