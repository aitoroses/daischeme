package schemastore

import (
	parser "github.com/daischio/daischeme/codemodel/schemaparser"
	"fmt"
	"github.com/daischio/daischeme/codemodel/util"
)

type SchemaStore struct {
	ModelName string
	schema *parser.Scheme
}

// General named schema
type NamedSchema struct {
	Name string          // schema name
	*parser.Scheme
}

func New(modelName string, s *parser.Scheme) *SchemaStore {
	st := &SchemaStore{modelName, s};
	return st
}

func (s *SchemaStore) GetSchema() *NamedSchema {
	return &NamedSchema{s.ModelName, s.schema}
}

func (s *SchemaStore) GetSchemas() []*NamedSchema {
	// Create
	schemas := make([]*NamedSchema, 0)

	// Add the first schema
	schemas = append(schemas, &NamedSchema{s.ModelName, s.schema})

	// Define a walking function
	var walk func(s *parser.Scheme)
	walk = func(s *parser.Scheme) {
		// Iterate over properties and get nested schemas
		for k, v := range s.Properties {
			fmt.Printf("%v: %+v\n",k,v)
			if v.Type == "object" && v.Properties != nil {
				// append the schema
				schemas = append(schemas, &NamedSchema{util.Capitalize(k),&v})
				// walk that schema also
				walk(&v)
			}
		}
	}
	// Iterate
	walk(s.schema)
	return schemas
}
