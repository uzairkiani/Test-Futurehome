package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"

	"net/http"
	"os"

)


type State struct {

	ID     string  `json:"id"`
	Name  string  `json:"Name"`
	Action string `json:"action"`
	Transition string `json:"Transition"`

}

var states []State

const jsonStates ="./states.db.json"



type Password struct {
	Pass     string  `json:"pass"`
}
var password []Password
func getPassword(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	for _, item := range password {
		fmt.Println("Endpoint Hit:", params["pass"])
		if (item.Pass== params["pass"]) {
			createJsonState()

			json.NewEncoder(w).Encode(item)
		}
		if (item.Pass!= params["pass"]){
			fmt.Fprintf(w, "Password didnot match!")
		}
	}
json.NewEncoder(w).Encode(&Password{})
}

// Delete book

func homePage(w http.ResponseWriter, r *http.Request ){
	fmt.Fprintf(w, "Welcome")
	fmt.Println("Endpoint Hit: homePage")
}


func createJsonState() {

	states = append(states, State{ID: "1", Name : "Lock", Action:"Unlock", Transition:"Unlocked" })
	states = append(states, State{ID: "2", Name : "Unlock", Action:"Lock", Transition:"Locked" })

	//	fmt.Println(users)
	var buf = new(bytes.Buffer)

	enc := json.NewEncoder(buf)

	enc.Encode(states)


	g, err := os.Create(jsonStates)
	if nil != err {
		log.Fatalln(err)
	}
	defer g.Close()
	io.Copy(g, buf)
}

func main() {
	// Init router
	r := mux.NewRouter()
	// Hardcoded data - @todo: add database
	password = append(password, Password{Pass:"1234"})
	// Route handles & endpoints
	r.HandleFunc("/", homePage	)
	r.HandleFunc("/password", getPassword	).Methods("GET")
	r.HandleFunc("/password/{pass}", getPassword	).Methods("GET")


	var buf = new(bytes.Buffer)

	db:=states
	enc1 := json.NewEncoder(buf)


	enc1.Encode(db)

	g, err := os.Create(jsonStates)
	if nil != err {
		log.Fatalln(err)
	}
	defer g.Close()
	io.Copy(g, buf)


	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}