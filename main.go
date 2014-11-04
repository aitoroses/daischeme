package main

import (
	"fmt"
	"github.com/daischio/daischeme/codemodel"
)

func main() {
	// Generate a model from a schema
	m := codemodel.New("MyPackage", "MyModel", "./assets/json_example_schema.json")
	fmt.Printf("%+v\n", m)

	// Generate the SchemaStore from a schema
	schemaStore := m.GetSchemaStore()
	fmt.Printf("main: %+v\n", *schemaStore.GetSchema())

	// Obtain the schemas
	schemas := schemaStore.GetSchemas()
	fmt.Printf("\nSchemas: %+v\n", schemas)

	// Obtain the mapped structures
	structs := m.GetSchemaMapper().GetMappedStructs()

	// Print the struct objects
	fmt.Printf("\nStructs: %+v\n", structs)
	for _, s := range structs {
		fmt.Printf("\t -%+v\n", s)
		fmt.Printf("\t\t Fields: %+v\n", s)
		for _, f := range s.Fields {
			fmt.Printf("\t\t\t -%+v\n", f)
		}

		// Print the struct as code
		fmt.Printf("\n%s", s.Code())
	}

}
