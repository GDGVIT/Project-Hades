package model

import (
	"fmt"
	"log"
	"os"

	"github.com/casbin/casbin"
	mysqladapter "github.com/casbin/mysql-adapter"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

var (
	con      bolt.Conn
	enforcer *casbin.Enforcer
)

func DBInit(c bolt.Conn) {
	con = c
}

// connect to neo
func ConnectToDB() bolt.Conn {

	conn, err := bolt.NewDriver().OpenNeo(os.Getenv("PROD_URI"))
	if err != nil {
		log.Fatalln("Error connecting to DB")
	}
	return conn
}

// connect policy enforcer
// and wrappers around it
func ConnectEnforcer() {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/", "root", os.Getenv("MYSQL_ROOT_PASSWORD"), "db", "3306")
	adapter := mysqladapter.NewAdapter("mysql", uri)
	enforcer = casbin.NewEnforcer("../policy.conf", adapter)
	enforcer.EnableAutoSave(true)
}

func Enforce(who, resource, access string) bool {
	return enforcer.Enforce(who, resource, access)
}

func AddPolicy(who, resource, access string) error {
	enforcer.AddPolicy(who, resource, access)
	return enforcer.SavePolicy()
}

func RemovePolicy(who, resource, access string) error {
	enforcer.RemovePolicy(who, resource, access)
	return enforcer.SavePolicy()
}
