package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Load the json file
	f, err := ioutil.ReadFile("./assets/json_example.json")
	check(err)

	// Parse de json file using the model
	myModel := &MyModel{}
	json.Unmarshal([]byte(f), myModel)
	fmt.Printf("%+v",myModel)
}
