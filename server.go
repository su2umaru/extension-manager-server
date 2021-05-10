package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Extension struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// type Extensions []Extension

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	var extensions []Extension
	result := Db.Find(&extensions)
	output, _ := json.Marshal(extensions)
	w.Write(output)
	return result.Error
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	var extension Extension
	json.NewDecoder(r.Body).Decode(&extension)
	result := Db.Create(&extension)
	return result.Error
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	var extension Extension
	json.NewDecoder(r.Body).Decode(&extension)
	fmt.Print(extension)
	result := Db.Delete(&extension)
	return result.Error
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")

	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
