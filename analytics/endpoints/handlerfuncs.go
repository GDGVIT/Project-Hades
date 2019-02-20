package endpoints

import (
	"encoding/json"
	"net/http"

	db "github.com/GDGVIT/Project-Hades/analytics/modelfuncs"
)

func readFromDB() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// make a map of all GET queries
		queries := r.URL.Query()
		var (
			arr    [][]db.Logs
			errs   []error
			buflen = len(queries)
			count  int
		)

		// make a channel with buffer length same as the number of queries
		ch := make(chan db.LogsReturn, buflen)
		for key, value := range queries {
			go db.ReadLogs(key, value[0], ch)
		}

		// query channel buffers, then close channel
		for i := range ch {
			count++
			if i.Err != nil {
				errs = append(errs, i.Err)
			}
			arr = append(arr, i.Logs)
			if count == buflen {
				break
			}
		}
		close(ch)

		// send JSON response
		json.NewEncoder(w).Encode(struct {
			Logs   [][]db.Logs `json:"logs"`
			Errors []error     `json:"errors"`
		}{
			arr,
			errs,
		})

	}
}
