package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	dialer "github.com/GDGVIT/Project-Hades/exporter/dialer"
	"github.com/GDGVIT/Project-Hades/exporter/endpoints/services"
	"github.com/GDGVIT/Project-Hades/model"
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

	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	specific := r.URL.Query().Get("specific")
	event := r.URL.Query().Get("event")

	// fetch results for result query
	result := dialer.FetchAllRequest{
		Event: event,
		Query: model.Query{
			Key:      key,
			Value:    value,
			Specific: specific,
		},
	}
	ret, err := result.FetchAllQ(key, value, specific)

	if err != nil {
		log.Println(err)
		log.Println("Error occurred while fetching all")
		json.NewEncoder(w).Encode(struct {
			Err error `json:"err"`
		}{err})
		return
	}

	// create CSV
	ce := make(chan error)
	arg := *ret
	go services.CreateCSV(arg.Rs, ce)

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

	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	specific := r.URL.Query().Get("specific")
	event := r.URL.Query().Get("event")

	// fetch results for result query
	result := dialer.FetchAllRequest{
		Event: event,
		Query: model.Query{
			Key:      key,
			Value:    value,
			Specific: specific,
		},
	}
	ret, err := result.FetchAllQ(key, value, specific)

	if err != nil {
		log.Println(err)
		log.Println("Error occurred while fetching all")
		json.NewEncoder(w).Encode(struct {
			Err error `json:"err"`
		}{err})
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
