package services

import (
	"encoding/json"
	"os"

	"github.com/GDGVIT/Project-Hades/model"
)

func CreateJSON(data []model.Participant, c chan error) {
	f, err := os.Create("./participants.json")
	if err != nil {
		c <- err
		return
	}
	defer f.Close()

	d, err := json.Marshal(data)

	if err != nil {
		c <- err
		return
	}

	f.Write(d)
	c <- nil
	return
}
