package model

import bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"

var con bolt.Conn

func DBInit(c bolt.Conn) {
	con = c
}
