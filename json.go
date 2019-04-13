package main

// ref: https://golang.org/pkg/encoding/json/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/striversity/go-on-the-run/types"
)

type Log struct {

	ID     string  `json:"id"`
	Match  string  `json:"match"`
	Unmatch  string `json:"unmatch"`

}
const jsonFile = "./user.db.json"

func main() {
	createJsonFile()

	f, err := os.Open(jsonFile)
	g, err :=os.Open(jsonStates)
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()
	defer g.Close()

	dec := json.NewDecoder(f)
	db := types.UserDb{}

	db1:=states
	dec1:= json.NewDecoder(g)


	dec1.Decode(&db1)

	dec.Decode(&db)
	fmt.Println(db)
}

func createJsonFile() {

    states = append(states, State{ID: "1", Name : "Lock", Action:"Unlock", Transition:"Unlocked" })
	states = append(states, State{ID: "2", Name : "Unlock", Action:"Lock", Transition:"Locked" })






	db := types.UserDb{Users: users, Type: "Simple"}
	db1  :=states

	//	fmt.Println(users)
	var buf = new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	enc1 := json.NewEncoder(buf)
	enc.Encode(db1)

	enc1.Encode(db)

	f, err := os.Create(jsonFile)
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()
	io.Copy(f, buf)






	g, err := os.Create(jsonStates)
	if nil != err {
		log.Fatalln(err)
	}
	defer g.Close()
	io.Copy(g, buf)



}