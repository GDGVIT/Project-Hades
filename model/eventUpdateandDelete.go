package model

import (
	"log"
)

// delete event with given query
func DeleteEvent(q Query, c chan error) {
	result, err := con.ExecNeo(`
		MATCH(n:EVENT)-[r]->(a)
		WHERE n.`+q.Key+`=$val
		DETACH DELETE n, a
	`, map[string]interface{}{
		"val": q.Value,
	})
	if err != nil {
		c <- err
	}

	log.Println(result)
	log.Println("Event deleted")
	c <- nil
	return
}

// update event with given query and new value
func UpdateEvent(q Query, c chan error) {
	result, err := con.ExecNeo(`
		MATCH(n:EVENT)
		WHERE n.`+q.Key+`=$val
		SET n.`+q.ChangeKey+`=$val1
		RETURN n.`+q.ChangeKey+`
	`, map[string]interface{}{
		"val":  q.Value,
		"val1": q.ChangeValue,
	})

	if err != nil {
		c <- err
		return
	}
	c <- nil

	log.Println(result)
	log.Printf("Updated")

}
