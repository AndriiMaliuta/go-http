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
	"strconv"
)

type PersonsHandle struct{}
type PersonHandler struct{}

func initDb() *sql.DB {
	pswd, _ := os.LookupEnv("PSQL_PASS")
	hmIP, _ := os.LookupEnv("HM_IP")
	hmUs, _ := os.LookupEnv("HM_USER")
	db, err := sql.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:5439/database?sslmode=disable", pswd, hmIP))
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
//		fmt.Sprintf("postgresql://hostman:%s@IP:5439/database?sslmode=false", pass))
//	if err != nil {
//		log.Panicln(err)
//	}
//	Db.Ping()
//}

func getPersons() []models.Person {
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

func personById(id int) models.Person {
	db := initDb()
	//person := models.Person{}
	row := db.QueryRow("select person_id, name, gender, status from persons where person_id = $1", id)
	person := models.Person{}
	row.Scan(&person.Id, &person.Name, &person.Gender, &person.Status)
	defer db.Close()

	return person
}

func (p PersonsHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//rURL := r.RequestURI
	//usName := r.URL.User.Username()
	//msg := "Hello Handler! URI is " + rURL + " Name is " + name + " Query name is " + qName + " Encoded queries are " + encQuers
	persons := getPersons()
	peBytes, err := json.Marshal(persons)
	if err != nil {
		log.Panicln(err)
	}
	w.Write([]byte(peBytes))
	//w.Write([]byte("user name is " + usName))
	//w.Write([]byte(name))
}

func (p PersonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//httpMethod := r.Method
	pId := r.URL.Query().Get("id")
	//id, _ := strconv.Atoi(path.Base(r.URL.Path))
	atoi, _ := strconv.Atoi(pId)
	person, _ := json.Marshal(personById(atoi))
	log.Println("Getting person " + pId)
	w.Write([]byte(person))
}

func main() {
	//pers := getPersons()
	//log.Println(pers)
	psHandler := PersonsHandle{}
	pHandler := PersonHandler{}
	server := http.Server{Addr: "127.0.0.1:8087"}
	http.Handle("/persons", &psHandler)
	http.Handle("/person", &pHandler)
	server.ListenAndServe()
}
