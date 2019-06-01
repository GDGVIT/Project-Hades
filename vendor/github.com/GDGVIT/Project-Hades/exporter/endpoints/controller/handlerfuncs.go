package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	dialer "github.com/GDGVIT/Project-Hades/exporter/dialer"
	"github.com/GDGVIT/Project-Hades/exporter/endpoints/services"
)

/**
*@api {get} /api/v1/exporter/excel export participants as excel/csv
*@apiName export participants as excel/csv
*@apiGroup exporter
*@apiPermission admin
*@apiParam {String} [key] key to query the event by
*@apiParam {String} [value] value of the key
*@apiParam {String} event event name
*@apiParam {String} specific types of participants to export i.e. present/absent/all
*
*@apiParamExample {json} request-example1
*{
*	"event":"DEVFEST 2019",
*	"query":{
*		"key":"gender",
*		"value":"F",
*		"specific":"project-all"
*	}
*}
*
*@apiParamExample {json} request-example2
*{
*	"event":"DEVFEST 2019",
*	"query":{
*		"key":"gender",
*		"value":"F",
*		"specific":"project-present"
*	}
*}
*
*
*@apiParamExample {json} request-example3
*{
*	"event":"DEVFEST 2019",
*	"query":{
*		"key":"gender",
*		"value":"F",
*		"specific":"project-absent"
*	}
*}
**/
func HandleExcel(w http.ResponseWriter, r *http.Request) {

	var result dialer.FetchAllRequest

	json.NewDecoder(r.Body).Decode(&result)

	c := make(chan dialer.FetchAllResponse)

	// fetch results for result query
	go result.FetchAll(c)

	ret := <-c
	if err := ret.Err; err != nil {
		log.Println(err)
		log.Println("Error occurred while fetching all")
		json.NewEncoder(w).Encode(ret)
		return
	}

	// create CSV
	ce := make(chan error)
	go services.CreateCSV(ret.Rs, ce)

	if err := <-ce; err != nil {
		log.Println(err)
	}

	// set file to download
	w.Header().Set("Content-Disposition", "attachment; filename=participants.csv")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	http.ServeFile(w, r, "participants.csv")

	// cleanup
	go os.Remove("participants.csv")
}

/**
*@api {get} /api/v1/exporter/json export participants as json
*@apiName export participants as json
*@apiGroup exporter
*@apiPermission admin
*@apiParam {String} [key] key to query the event by
*@apiParam {String} [value] value of the key
*@apiParam {String} event event name
*@apiParam {String} specific types of participants to export i.e. present/absent/all
*
*@apiParamExample {json} request-example1
*{
*	"event":"DEVFEST 2019",
*	"query":{
*		"key":"gender",
*		"value":"F",
*		"specific":"project-all"
*	}
*}
*
*@apiParamExample {json} request-example2
*{
*	"event":"DEVFEST 2019",
*	"query":{
*		"key":"gender",
*		"value":"F",
*		"specific":"project-present"
*	}
*}
*
*
*@apiParamExample {json} request-example3
*{
*	"event":"DEVFEST 2019",
*	"query":{
*		"key":"gender",
*		"value":"F",
*		"specific":"project-absent"
*	}
*}
**/
func HandleJson(w http.ResponseWriter, r *http.Request) {

	var result dialer.FetchAllRequest

	json.NewDecoder(r.Body).Decode(&result)

	c := make(chan dialer.FetchAllResponse)

	// fetch results for result query
	go result.FetchAll(c)

	ret := <-c
	if err := ret.Err; err != nil {
		log.Println(err)
		log.Println("Error occurred while fetching all")
		json.NewEncoder(w).Encode(ret)
		return
	}

	// create JSON
	ce := make(chan error)
	go services.CreateJSON(ret.Rs, ce)

	if err := <-ce; err != nil {
		log.Println(err)
	}

	// set file to download
	w.Header().Set("Content-Disposition", "attachment; filename=participants.json")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	http.ServeFile(w, r, "participants.json")

	// cleanup
	go os.Remove("participants.json")
}
