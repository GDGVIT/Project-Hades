package model

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

func GetOrgs(org string) ([]Organization, error) {
	res, _, _, err := con.QueryNeoAll(`
					MATCH (n:ORG)
					WHERE n.name STARTS WITH $name
					RETURN n.name, n.location, n.description, n.tag, n.createdAt, n.website
				`, map[string]interface{}{
		"name": org,
	})
	if err != nil {
		return nil, err
	}
	orgs := []Organization{}
	for i, _ := range res {
		orgs = append(orgs, Organization{
			Name:        res[i][0].(string),
			Location:    res[i][1].(string),
			Description: res[i][2].(string),
			Tag:         res[i][3].(string),
			CreatedAt:   res[i][4].(string),
			Website:     res[i][5].(string),
		})
	}
	return orgs, nil
}

func CreateJoinRequest(user string, org string) error {

	data, _, _, err := con.QueryNeoAll(`
				MATCH(n:ORG)<-[r:JOIN]-(a:USER)
				WHERE n.name = $org AND a.email = $user
				RETURN r
				`, map[string]interface{}{
		"org":  org,
		"user": user,
	})

	if len(data) > 1 {
		return errors.New("A join request is already pending")
	}
	_, err = con.ExecNeo(`
					MATCH(n:ORG) WHERE n.name = $org
					MATCH(a:USER) WHERE a.email = $user
					CREATE (n)<-[:JOIN]-(a)
				`, map[string]interface{}{
		"org":  org,
		"user": user,
	})
	if err != nil {
		return err
	}
	return nil
}

func AcceptJoinRequest(user string, org string) error {

	data, _, _, err := con.QueryNeoAll(`
				MATCH(n:ORG)<-[r:JOIN]-(a:USER)
				WHERE n.name = $org AND a.email = $user
				RETURN r
				`, map[string]interface{}{
		"org":  org,
		"user": user,
	})

	if len(data) <= 1 {
		return errors.New("No join request found")
	}

	_, err = con.ExecNeo(`
					MATCH(n:ORG)<-[r:JOIN]-(a:USER) 
					WHERE n.name = $org AND a.email = $user
					DELETE r
					CREATE (n)<-[:MEMBER]-(a)
				`, map[string]interface{}{
		"org":  org,
		"user": user,
	})
	if err != nil {
		return err
	}
	return nil
}
