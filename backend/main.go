package main

import (
	"database/sql"
  "strings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
  "os"

  "github.com/joho/godotenv"
  "github.com/rs/cors"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type People struct {
  ID    int     `json:"id"`
  Name  string  `json:"name"`
}

func connect() (*sql.DB, error) {
  passwordFile := os.Getenv("POSTGRES_PASSWORD_FILE")
  dbPort := os.Getenv("POSTGRES_PORT")
  dbName := os.Getenv("POSTGRES_DB")
  dbUser := os.Getenv("POSTGRES_USER")

	bin, err := ioutil.ReadFile(passwordFile)
	if err != nil {
		return nil, err
	}
  password := strings.TrimRight(string(bin), "\n")
	return sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable", dbUser, password, dbPort, dbName))
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connect()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM people")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var people = []*People{}
	for rows.Next() {
    person := new(People)
		err = rows.Scan(
      &person.ID,
      &person.Name,
    )
    if err != nil {
      w.WriteHeader(500)
      return
    }
		people = append(people, person)
	}
  jsonData, err := json.Marshal(people)
  if err != nil {
    log.Fatal("Error marshalling json data!")
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(jsonData)
}

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	log.Print("Prepare db...")
	if err := prepare(); err != nil {
		log.Fatal(err)
	}

  apiPort := fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT"))
  webPort := fmt.Sprintf(":%s", os.Getenv("FRONTEND_PORT"))

  log.Print("Listening ", apiPort)
	r := mux.NewRouter()
	r.HandleFunc("/", peopleHandler)

  c := cors.New(cors.Options{
    AllowedOrigins: []string{webPort},
    AllowedMethods: []string{
      http.MethodGet,
      http.MethodPost,
      http.MethodPut,
      http.MethodPatch,
      http.MethodDelete,
      http.MethodOptions,
      http.MethodHead,
    },
    AllowedHeaders: []string{"*"},
    AllowCredentials: true,
  })

  handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(apiPort, handler))
}

func prepare() error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	for i := 0; i < 60; i++ {
		if err := db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	if _, err := db.Exec("DROP TABLE IF EXISTS people"); err != nil {
		return err
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS people (id SERIAL, name VARCHAR)"); err != nil {
		return err
	}

	for i := 0; i < 5; i++ {
		if _, err := db.Exec("INSERT INTO people (name) VALUES ($1);", fmt.Sprintf("Person #%d", i)); err != nil {
			return err
		}
	}
	return nil
}
