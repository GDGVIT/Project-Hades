package db

import (
	"fmt"
	"log"
)

func CreateLogs(subject, timestamp string, data []byte) {
	conn.Create(&Logs{Subject: subject, Timestamp: timestamp, Body: string(data)})
	log.Println("Created %s log", subject)
}

func ReadLogs(key, value string, ch chan LogsReturn) {
	var res []Logs
	conn.Where(fmt.Sprintf("%s = ?", key), value).Find(&[]Logs{}).Scan(&res)
	ch <- LogsReturn{res, nil}
	return
}

func ReadAllLogs(ch chan LogsReturn) {
	var res []Logs
	conn.Find(&[]Logs{}).Scan(&res)
	ch <- LogsReturn{res, nil}
	return
}
