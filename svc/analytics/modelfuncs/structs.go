package db

import "github.com/jinzhu/gorm"

type Logs struct {
	gorm.Model
	Subject   string
	Timestamp string
	Body      string
}

type LogsReturn struct {
	Logs []Logs `json:"logs"`
	Err  error  `json:"err"`
}
