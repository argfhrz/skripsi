package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

//Object
type successResponseObject struct {
	Message string
	Data    []byte
}

type userObject struct {
	Email    string
	Password string
}

//Function Helper
func initDatabase(database *sql.DB) *sql.Tx {
	tx, err2 := database.Begin()
	if err2 != nil {
		log.Println(err2)
	}

	stmt, err3 := tx.Prepare("CREATE TABLE IF NOT EXISTS userList (email TEXT PRIMARY KEY, password TEXT)")
	if err3 != nil {
		log.Println(err3)
	}
	stmt.Exec()
	defer stmt.Close()

	return tx

}

func main() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)

	http.ListenAndServe(":8080", nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		log.Println("Request method is not POST")
		w.WriteHeader(404)
		http.Error(w, "Request method is wrong", http.StatusNotFound)
		return

	}

	if err0 := r.ParseMultipartForm(0); err0 != nil {
		log.Println("Post body is wrong")
		w.WriteHeader(404)
		http.Error(w, "Post body is wrong", http.StatusNotFound)
		return
	}

	mEmail := r.FormValue("email")
	mPassword := r.FormValue("password")

	if (mEmail == "") || (mPassword == "") {
		w.WriteHeader(404)
		http.Error(w, "Post body is wrong", http.StatusNotFound)
		return

	}

	database, err1 := sql.Open("sqlite3", "./rs_ningsih_tinampi.db")
	if err1 != nil {
		log.Println(err1)
		w.WriteHeader(500)
		http.Error(w, "Open database error", http.StatusInternalServerError)
		return
	}
	tx := initDatabase(database)
	defer database.Close()
	defer tx.Commit()

	mEmailDatabase := ""
	isUserMatch := false
	rows, err2 := tx.Query("SELECT email FROM userList WHERE email=?", mEmail)
	if err2 != nil {
		log.Println(err2)
		w.WriteHeader(500)
		http.Error(w, "Read Database error", http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		rows.Scan(&mEmailDatabase)
		isUserMatch = true
		break

	}

	if !isUserMatch {
		stmt, err2 := tx.Prepare("INSERT INTO userList (email, password) VALUES (?, ?)")
		if err2 != nil {
			log.Println(err2)
			w.WriteHeader(500)
			http.Error(w, "Register user error", http.StatusInternalServerError)
			return
		}
		stmt.Exec(mEmail, mPassword)
		defer stmt.Close()
	}

	if isUserMatch {
		log.Println("User already exists")
		w.WriteHeader(409)
		http.Error(w, "User already exists", http.StatusConflict)
		return

	}

	m2 := successResponseObject{"Register success", nil}
	a, err3 := json.Marshal(m2)
	if err3 != nil {
		log.Println(err3)
		w.WriteHeader(500)
		http.Error(w, "Format response error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(a)

}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		log.Println("Request method is not POST")
		w.WriteHeader(404)
		http.Error(w, "Request method is wrong", http.StatusNotFound)
		return

	}

	if err0 := r.ParseMultipartForm(0); err0 != nil {
		log.Println("Post body is wrong")
		w.WriteHeader(404)
		http.Error(w, "Post body is wrong", http.StatusNotFound)
		return
	}

	mEmail := r.FormValue("email")
	mPassword := r.FormValue("password")

	if (mEmail == "") || (mPassword == "") {
		w.WriteHeader(404)
		http.Error(w, "Post body is wrong", http.StatusNotFound)
		return

	}

	database, err1 := sql.Open("sqlite3", "./rs_ningsih_tinampi.db")
	if err1 != nil {
		log.Println(err1)
		w.WriteHeader(500)
		http.Error(w, "Open database error", http.StatusInternalServerError)
		return
	}
	tx := initDatabase(database)
	defer database.Close()
	defer tx.Commit()

	mEmailDatabase := ""
	mPasswordDatabase := ""
	isUserMatch := false
	log.Println(mEmail)
	rows, err2 := tx.Query("SELECT email, password FROM userList WHERE email=?", mEmail)
	if err2 != nil {
		log.Println(err2)
		w.WriteHeader(500)
		http.Error(w, "Read Database error", http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		rows.Scan(&mEmailDatabase, &mPasswordDatabase)
		if (mEmail == mEmailDatabase) && (mPassword == mPasswordDatabase) {
			isUserMatch = true
			break

		}

	}

	m2 := successResponseObject{"Login success", nil}
	if !isUserMatch {
		log.Println("User not found")
		w.WriteHeader(404)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	a, err3 := json.Marshal(m2)
	if err3 != nil {
		log.Println(err3)
		w.WriteHeader(500)
		http.Error(w, "Format response error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(a)

}
