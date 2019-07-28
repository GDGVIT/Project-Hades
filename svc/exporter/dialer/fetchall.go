package dialer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GDGVIT/Project-Hades/model"
)

type FetchAllRequest struct {
	Event string      `json:"event"`
	Query model.Query `json:"query"`
}

type FetchAllResponse struct {
	Rs  []model.Participant `json:"rs"`
	Err error               `json:"err"`
}

func (req FetchAllRequest) FetchAll(c chan FetchAllResponse) {

	byteMsg, err := json.Marshal(req)

	if err != nil {
		c <- FetchAllResponse{nil, err}
	}

	resp, err := http.Post("http://simple_projection:8083/api/v1/simple-projection/"+req.Query.Specific,
		"application/json",
		bytes.NewBuffer(byteMsg),
	)

	if err != nil {
		c <- FetchAllResponse{nil, fmt.Errorf("Error making request")}
		return
	}

	var result FetchAllResponse

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		c <- result
		return
	}

	c <- result

	close(c)
}
