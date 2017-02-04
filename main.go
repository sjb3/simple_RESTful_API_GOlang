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
  params := mux.Vars(req)
  for _, item := range people {
    if item.ID == params["id"]{
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Person{}) //no id empty object returns
}

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request){
  json.NewEncoder(w).Encode(people)
}

func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  var person Person
  _ = json.NewDecoder(req.Body).Decode(&person)
  person.ID = params["id"]
  people = append(people, person)
  json.NewEncoder(w).Encode(people)
}


func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  for index, item := range people{
    if item.ID == params["id"]{
      people = append(people[:index], people[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(people)
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



