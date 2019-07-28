package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var conn *gorm.DB

func DBconnect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "logs.db")
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	log.Println("Connected to sql for logs")
	conn = db

	// migrate models into db
	conn.AutoMigrate(&Logs{})
	return conn
}
