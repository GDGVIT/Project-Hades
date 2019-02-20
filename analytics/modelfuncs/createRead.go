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
	rows, err := conn.Where(fmt.Sprintf("%s = ?", "Subject"), "subject").Find(&Logs{}).Rows()
	if err != nil {
		log.Printf("Error reading from logs database: %v", err)
		ch <- LogsReturn{nil, err}
		return
	}
	var logs []Logs
	for rows.Next() {
		var l Logs
		conn.ScanRows(rows, &l)
		logs = append(logs, l)
	}
	ch <- LogsReturn{logs, nil}
	return
}
