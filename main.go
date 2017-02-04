package main

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "log"
  "net/http"
)

type Person struct {
  ID          string    `json:"id,omitempty"`
  Firstname   string    `json:"firstname,omitempty"`
  Lastname    string    `json:"lastname,omitempty"`
  Address     *Address  `json:"address,omitempty"`
}

type Address struct {
  City  string `json:"city,omitempty"`
  State string `json:"state,omitempty"`
}

var people []Person


func GetPersonEndPoint(w http.ResponseWriter, req *http.Request){

}

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request){
  json.NewEncoder(w).Encode(people)
}

func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request){

}


func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request){

}
func main(){
  router := mux.NewRouter()
  people = append(people, Person{ID: "1", Firstname: "Lucien", Lastname: "What",
    Address: &Address{City: "Seattle", State: "WA"}})
  people = append(people, Person{ID: "2", Firstname: "Justin", Lastname: "BB"})

  router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
  router.HandleFunc("/people/{id}", GetPeopleEndPoint).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000", router))
}



