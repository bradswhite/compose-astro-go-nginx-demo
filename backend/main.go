package main

import (
  "log"
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/rs/cors"
)

func main() {
  type GreetingType struct {
    Id    int     `json:"id"`
    Name  string  `json:"name"`
  }

  r := mux.NewRouter()

  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    data := GreetingType{ Id: 1, Name: "Docker/Go" }
    jsonData, err := json.Marshal(data)
    if err != nil {
      log.Fatal("Error marshalling json data!")
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData)
    //w.Write([]byte("Docker"))
  })

  c := cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:8080"},
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

  fmt.Println("Server is starting at port 3000")
  log.Fatal(http.ListenAndServe(":3000", handler))
}
