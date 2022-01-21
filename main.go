package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"go-http/models"
	"log"
	"net/http"
	"os"
)

type PersonsHandle struct{}
type WorldHandler struct{}

func initDb() *sql.DB {
	pswd, _ := os.LookupEnv("PSQL_PASS")
	//db, err := sql.Open("postgres", fmt.Sprintf("postgres://dev2:%s@localhost/cities?sslmode=disable", pswd))
	db, err := sql.Open("postgres",
		fmt.Sprintf("postgres://hostman:%s@94.228.113.90:5439/database?sslmode=disable", pswd))
	if err != nil {
		log.Panicln(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("You connected to your database.")
	return db
}

//var Db *sql.DB

//func init() {
//	var err error
//	pass, _ := os.LookupEnv("PSQL_PASS")
//	Db, err := sql.Open("postgres",
//		fmt.Sprintf("postgresql://hostman:%s@94.228.113.90:5439/database?sslmode=false", pass))
//	if err != nil {
//		log.Panicln(err)
//	}
//	Db.Ping()
//}

func retrieve() []models.Person {
	db := initDb()
	//person := models.Person{}
	rows, err := db.Query("select person_id, name, gender, status from persons")
	if err != nil {
		log.Panicln(err)
	}
	persons := make([]models.Person, 0)
	for rows.Next() {
		person := models.Person{}
		err := rows.Scan(&person.Id, &person.Name, &person.Gender, &person.Status)
		if err != nil {
			log.Panicln(err)
		}
		persons = append(persons, person)
	}
	defer db.Close()

	return persons
}

func (p PersonsHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//rURL := r.RequestURI
	//name := r.Form.Get("name")
	//qValues := r.URL.Query()
	//encQuers := qValues.Encode()
	//qName := r.URL.Query().Get("name")
	//usName := r.URL.User.Username()
	//msg := "Hello Handler! URI is " + rURL + " Name is " + name + " Query name is " + qName + " Encoded queries are " + encQuers
	persons := retrieve()
	peBytes, err := json.Marshal(persons)
	if err != nil {
		log.Panicln(err)
	}
	w.Write([]byte(peBytes))
	//w.Write([]byte("user name is " + usName))
	//w.Write([]byte(name))
}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
	w.Write([]byte("World handler"))
}

func main() {
	pers := retrieve()
	log.Println(pers)
	pHandler := PersonsHandle{}
	world := WorldHandler{}
	server := http.Server{Addr: "127.0.0.1:8087"}
	http.Handle("/persons", &pHandler)
	http.Handle("/world", &world)
	server.ListenAndServe()
}
