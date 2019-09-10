package endpoints

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/GDGVIT/Project-Hades/model"
)

func bulkAddAttendees() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// TODO: validate if the event is there on the org or not

		data := struct {
			Attendees []model.Participant `json:"attendees"`
			EventName string              `json:"eventName"`
		}{}

		json.NewDecoder(r.Body).Decode(&data)
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
