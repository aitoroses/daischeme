package main

import (
	"fmt"
	"github.com/daischio/daischeme/codemodel"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Generate a model from a schema
	m := codemodel.New("main", "MyModel", "./assets/json_example_schema.json")

	// Obtain the mapped model code
	code := m.GetCode()
	fmt.Println(code)

	// Write model to disk
	err := ioutil.WriteFile("./model.go", []byte(m.GetCode()), 0644)
	check(err)
}


