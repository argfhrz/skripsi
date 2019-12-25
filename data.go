package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

//Object
type responseObject struct {
	Response string
}

type updateDataObject struct {
	Temperature string
}

type readDataObject struct {
	Temperature string
}

//Function Helper (pembuatan tabel dan kolom database)
func initDatabase(database *sql.DB) *sql.Tx {
	tx, err2 := database.Begin()
	if err2 != nil {
		log.Println(err2)
	}

	stmt, err3 := tx.Prepare("CREATE TABLE IF NOT EXISTS data (log TEXT PRIMARY KEY, temperature TEXT)")
	if err3 != nil {
		log.Println(err3)
	}
	stmt.Exec()
	defer stmt.Close()

	return tx

}

func updateResponseParser(request *http.Request) *updateDataObject {
	body, err0 := ioutil.ReadAll(request.Body)
	if err0 != nil {
		log.Println(err0)
	}
	var m updateDataObject
	err1 := json.Unmarshal(body, &m)
	if err1 != nil {
		log.Println(err1)
	}

	return &m
}

var cTemperature string = ""

func updateDataAndroid(aTemperature string) {
	if cTemperature != aTemperature {
		client := &http.Client{}
		postData := []byte("{\"to\": \"/topics/update\", \"data\": {\"message\": \"Server data is updated\"}}")
		req, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewReader(postData))
		if err != nil {
			os.Exit(1)
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "key=AIzaSyBW8x9GQEJKygHUoKhc3kowNWGrvh6p4LI")
		resp, err := client.Do(req)
		defer resp.Body.Close()

	}
}

func main() {
	log.Println("Mulai")
	mux := http.NewServeMux()

	mux.HandleFunc("/createData", createDataHandler)
	mux.HandleFunc("/readData", readDataHandler)

	http.ListenAndServe(":8080", mux)
}

func readDataHandler(w http.ResponseWriter, r *http.Request) {

	m2 := readDataObject{cTemperature}
	b, err2 := json.Marshal(m2)
	if err2 != nil {
		log.Println(err2)
	}
	w.Write(b)

}

func createDataHandler(w http.ResponseWriter, r *http.Request) {
	m := updateResponseParser(r)

	database, err0 := sql.Open("sqlite3", "./skripsi.db")
	if err0 != nil {
		log.Println(err0)
	}
	tx := initDatabase(database)
	defer database.Close()
	defer tx.Commit()

	stmt, err1 := tx.Prepare("INSERT INTO data (log, temperature) VALUES (?,?)")
	if err1 != nil {
		log.Println(err1)
	}

	sourceTime := time.Now().UTC().Add(7 * time.Hour)
	nowTime := sourceTime.Format("2006-01-02 15:04:05")

	stmt.Exec(nowTime, m.Temperature)
	defer stmt.Close()

	cTemperature = m.Temperature

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	m2 := responseObject{"Create data success"}
	b, err2 := json.Marshal(m2)
	if err2 != nil {
		log.Println(err2)
	}
	w.Write(b)

	updateDataAndroid(m.Temperature)

}
