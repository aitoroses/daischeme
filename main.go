package main

import (
	"fmt"
	//"github.com/daischio/daischeme/codemodel"
	"io/ioutil"
	"encoding/json"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Generate a model from a schema
	/*m := codemodel.New("main", "MyModel", "./assets/json_example_schema.json")

	// Obtain the mapped model code
	code := m.GetCode()
	fmt.Println(code)

	// Write model to disk
	err := ioutil.WriteFile("./model.go", []byte(m.GetCode()), 0644)
	check(err)*/

	// Load the json file
	f, err := ioutil.ReadFile("./assets/json_example.json")
	check(err)

	// Parse de json file using the model
	myModel := &MyModel{}
	json.Unmarshal([]byte(f), myModel)
	fmt.Printf("%+v",myModel)
}
