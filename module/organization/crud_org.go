package organizationModule

import (
	"errors"
	"log"
	"time"
)

func CreateNewOrg(org Organization) error {
	data, _, _, err := con.QueryNeoAll(`
MATCH(n:ORG) WHERE n.name = $name
RETURN n.createdAt
				`, map[string]interface{}{
		"name": org.Name,
	})

	if err != nil {
		return err
	}

	// if org exists then throw error
	if len(data) > 1 {
		return errors.New("org already exists")
	}

	res, err := con.ExecNeo(`
					CREATE(n:ORG {name:$name, location: $location, description: $description, tag: $tag, createdAt: $cat, website: $website})
				`, map[string]interface{}{
		"name":        org.Name,
		"location":    org.Location,
		"description": org.Description,
		"tag":         org.Tag,
		"cat":         time.Now().String(),
		"website":     org.Website,
	})

	if err != nil {
		return err
	}
	log.Println(res)
	return nil
}

func InviteUserToOrg(email, org string) error {
	data, _, _, err := con.QueryNeoAll(`
					MATCH (n:ORG)-[:MEMBER]->(a:USER)
					WHERE n.name=$org AND a.email=$user
					RETURN n.name
				`, map[string]interface{}{
		"org":  org,
		"user": email,
	})

	if err != nil {
		return err
	}

	if len(data) > 1 {
		return errors.New("user is already a member")
	}

	data, _, _, err = con.QueryNeoAll(`
MATCH(n:ORG) WHERE n.name = $name
RETURN n.createdAt
				`, map[string]interface{}{
		"name": org,
	})

	if err != nil {
		return err
	}

	// if org exists then throw error
	if len(data) < 1 {
		return errors.New("org DNE")
	}

	data, _, _, err = con.QueryNeoAll(`
MATCH(n:USER) WHERE n.email = $name
RETURN n.createdAt
				`, map[string]interface{}{
		"name": email,
	})

	if err != nil {
		return err
	}

	// if org exists then throw error
	if len(data) < 1 {
		return errors.New("org DNE")
	}

	res, err := con.ExecNeo(`
		MATCH(n:ORG) WHERE n.name=$org
		MATCH (a:USER) WHERE a.email=$user
		CREATE (n)-[:MEMBER]->(a)
	`, map[string]interface{}{
		"org":  org,
		"user": email,
	})

	log.Println(res)
	return err

}
