package model

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func CreateNewOrg(org Organization, user string) error {
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
				MATCH(u:USER)
				WHERE u.email = $user
					CREATE(n:ORG {name:$name, location: $location, description: $description, tag: $tag, createdAt: $cat, website: $website})<-[:ADMIN]-(u)
					CREATE (n)<-[:MEMBER]-(u)
				`, map[string]interface{}{
		"name":        org.Name,
		"location":    org.Location,
		"description": org.Description,
		"tag":         org.Tag,
		"cat":         time.Now().String(),
		"website":     org.Website,
		"user":        user,
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
		CREATE (n)<-[:MEMBER]-(a)
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

	rss, _, _, err := con.QueryNeoAll(`
MATCH(n:ORG)<-[:JOIN]-(a:USER) WHERE n.name = $name
RETURN n.createdAt
				`, map[string]interface{}{
		"name": org,
	})

	if err != nil {
		return err
	}

	// if org exists then throw error
	if len(rss) > 0 {
		return errors.New("Request is already pending")
	}

	data, _, _, err := con.QueryNeoAll(`
				MATCH(n:ORG)<-[r:JOIN]-(a:USER)
				WHERE n.name = $org AND a.email = $user
				RETURN r
				`, map[string]interface{}{
		"org":  org,
		"user": user,
	})

	if len(data) > 0 {
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

func DenyJoinRequest(user, org string) error {

	data, _, _, err := con.QueryNeoAll(`
				MATCH(n:ORG)<-[r:JOIN]-(a:USER)
				WHERE n.name = $org AND a.email = $user
				RETURN r
				`, map[string]interface{}{
		"org":  org,
		"user": user,
	})

	if len(data) < 1 {
		return errors.New("No join request found")
	}

	_, err = con.ExecNeo(`
					MATCH(n:ORG)<-[r:JOIN]-(a:USER) 
					WHERE n.name = $org AND a.email = $user
					DELETE r
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

	if len(data) < 1 {
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

func GetJoinRequests(org string) ([]User, error) {
	data, _, _, err := con.QueryNeoAll(`
					MATCH (n:ORG)<-[:JOIN]-(a:USER)
					WHERE n.name = $name
					RETURN a.email, a.firstName, a.lastname, a.description,
					a.createdAt, a.facebook, a.linkedIn
				`, map[string]interface{}{
		"name": org,
	})
	if err != nil {
		return nil, err
	}

	users := []User{}
	for i, _ := range data {
		users = append(users, User{
			Email:       data[i][0].(string),
			FirstName:   data[i][1].(string),
			LastName:    data[i][2].(string),
			Description: data[i][3].(string),
			CreatedAt:   data[i][4].(string),
			Facebook:    data[i][5].(string),
			LinkedIn:    data[i][6].(string),
		})
	}
	return users, nil
}

func GetUserDetailsOnlyOrg(user string) (orgs []Organization, err error) {

	data, _, _, err := con.QueryNeoAll(`
MATCH (u:USER)-[:MEMBER]->(n:ORG)
WHERE u.email = $user
RETURN n.name, n.tag, n.location, n.description, n.createdAt
`, map[string]interface{}{
		"user": user,
	})
	fmt.Println(data, user)

	for i, _ := range data {
		orgs = append(orgs, Organization{
			Name:        data[i][0].(string),
			Tag:         data[i][1].(string),
			Location:    data[i][2].(string),
			Description: data[i][3].(string),
			CreatedAt:   data[i][4].(string),
		})

	}
	return orgs, nil
}

func GetUserDetails(user string) (events []Event, orgs []Organization, err error) {
	data, _, _, err := con.QueryNeoAll(`
MATCH (u:USER)-[:MEMBER]->(n:ORG)
MATCH(n)<-[:EVENT]-(e:EVENT)
WHERE u.email = $user
RETURN n.name, n.tag, n.location, n.description, n.createdAt,
e.clubName, e.name, e.toDate, e.fromDate, e.toTime, e.fromTime, e.budget, e.description, e.category
`, map[string]interface{}{
		"user": user,
	})
	fmt.Println(data, user)

	for i, _ := range data {
		orgs = append(orgs, Organization{
			Name:        data[i][0].(string),
			Tag:         data[i][1].(string),
			Location:    data[i][2].(string),
			Description: data[i][3].(string),
			CreatedAt:   data[i][4].(string),
		})

		if data[i][5] == nil {
			continue
		}
		events = append(events, Event{
			ClubName:    data[i][5].(string),
			Name:        data[i][6].(string),
			ToDate:      data[i][7].(string),
			FromDate:    data[i][8].(string),
			ToTime:      data[i][9].(string),
			FromTime:    data[i][10].(string),
			Budget:      data[i][11].(string),
			Description: data[i][12].(string),
			Category:    data[i][13].(string),
		})
	}
	return events, orgs, nil
}

func IsEventOfOrg(eventname, orgname string) (bool, error) {
	data, _, _, err := con.QueryNeoAll(`
MATCH(n:ORG)<-[r]-(a:EVENT) WHERE n.name = $name 
AND a.name = $eventname
RETURN n.createdAt
				`, map[string]interface{}{
		"name":      orgname,
		"eventname": eventname,
	})

	if err != nil {
		return false, err
	}

	// if org exists then throw error
	if len(data) > 1 {
		return false, nil
	} else {
		return true, nil
	}
}

func EnforceRoleMember(email, org string) (bool, error) {

	data, _, _, err := con.QueryNeoAll(`
MATCH(n:ORG)<-[r:MEMBER]-(a:USER) WHERE n.name = $name 
AND a.email = $email
RETURN n.createdAt
				`, map[string]interface{}{
		"name":  org,
		"email": email,
	})

	if err != nil {
		return false, err
	}

	// if role exists
	if fmt.Sprintf("%t", data) != "[]" {
		return true, nil
	} else {
		return false, nil
	}
}

func EnforceRoleAdmin(email, org string) (bool, error) {

	data, _, _, err := con.QueryNeoAll(`
MATCH(n:ORG)<-[r:ADMIN]-(a:USER) WHERE n.name = $name 
AND a.email = $email
RETURN n.createdAt
				`, map[string]interface{}{
		"name":  org,
		"email": email,
	})

	if err != nil {
		return false, err
	}

	// if role exists
	if fmt.Sprintf("%t", data) != "[]" {
		return true, nil
	} else {
		return false, nil
	}
}

func EnforceRoleEither(email, org string) (bool, error) {

	data, _, _, err := con.QueryNeoAll(`
MATCH(n:ORG)<-[r:MEMBER]-(a:USER) WHERE n.name = $name 
AND a.email = $email
RETURN n.createdAt
				`, map[string]interface{}{
		"name":  org,
		"email": email,
	})

	if err != nil {
		return false, err
	}

	// if role exists
	if fmt.Sprintf("%t", data) != "[]" {
		return true, nil
	} else {

		data, _, _, err = con.QueryNeoAll(`
MATCH(n:ORG)<-[r:ADMIN]-(a:USER) WHERE n.name = $name 
AND a.email = $email
RETURN n.createdAt
				`, map[string]interface{}{
			"name":  org,
			"email": email,
		})

		if err != nil {
			return false, err
		}

		// if role exists
		if fmt.Sprintf("%t", data) != "[]" {
			return true, nil
		} else {
			return false, nil
		}

	}
}
