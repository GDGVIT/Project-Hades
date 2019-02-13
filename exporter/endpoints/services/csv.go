package services

import (
	"encoding/csv"
	"os"

	"github.com/GDGVIT/Project-Hades/model"
)

func CreateCSV(data []model.Participant, c chan error) {
	f, err := os.Create("./participants.csv")
	if err != nil {
		c <- err
		return
	}
	defer f.Close()

	w := csv.NewWriter(f)

	for _, obj := range data {
		var record []string
		record = append(record, obj.Name)
		record = append(record, obj.Email)
		record = append(record, obj.RegistrationNumber)
		record = append(record, obj.Gender)
		record = append(record, obj.PhoneNumber)
		w.Write(record)
	}
	w.Flush()
	c <- nil
	return
}
