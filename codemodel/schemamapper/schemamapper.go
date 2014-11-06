package schemamapper

import (
	"github.com/daischio/daischeme/codemodel/util"
	parser "github.com/daischio/daischeme/codemodel/schemaparser"
	store "github.com/daischio/daischeme/codemodel/schemastore"
	"fmt"
)

type SchemaMapper struct {
	ModelName string
	schema *parser.Scheme
}

func New(modelName string, p *parser.Scheme) *SchemaMapper{
	sm := &SchemaMapper{modelName, p}
	return sm
}

func (sm *SchemaMapper) GetMappedStructs() []*Struct{

	structs := transformSchemasToStructs(sm)

	return structs
}

func transformSchemasToStructs(sm *SchemaMapper) []*Struct {

	structs := make([]*Struct, 0)

	var schemas []*store.NamedSchema = store.New(sm.ModelName, sm.schema).GetSchemas()

	for _, schema := range schemas {
		s := new(Struct)
		// Map schema properties to fields
		if schema.Properties == nil {
			continue
		}

		// Walk over properties
		for k, v := range schema.Properties {
			f := new(Field)
			f.Name = util.Capitalize(k)
			f.Tag = fmt.Sprintf(`json:"%s"`, k)
			f.Type = v.Type
			s.Fields = append(s.Fields, f)
		}

		//Add a name for the struct from the NamedSchema
		s.Name = schema.Name

		// Map the correct types
		resolvePropertyTypes(s)

		// Return the struct through the channel
		structs = append(structs,s)
	}
	return structs
}

func resolvePropertyTypes(s *Struct) {
	// Find the schema that maps each property
	for _, field := range s.Fields {
		switch field.Type {
		case "string":
			field.Type = "string"
		case "number":
			field.Type = "float64"
		case "boolean":
			field.Type = "bool"
		case "object":
			field.Type = field.Name
		default:
			field.Type = "string"
		}
	}
}

// ------------------------
// | Structs              |
// ------------------------

// Struct file holds a schema as a Struct representation
type Struct struct {
	Name string
	Fields []*Field
}

type Field struct {
	Tag string
	Name string
	Type string
}

// Write the struct header
func (s *Struct) head() string {
	return fmt.Sprintf("type %s struct {\n", s.Name)
}

// Write a field
func (s *Struct) field(i int) string {
	return fmt.Sprintf("\t%s %s \t\t\t`%s`\n", s.Fields[i].Name, s.Fields[i].Type, s.Fields[i].Tag)
}

// Write the end of file
func (s *Struct) tail() string {
	return fmt.Sprint("\n}\n")
}

// This function converts the struct to written code
func (s *Struct) Code() string {
	res := ""
	res += s.head()
	for i := range s.Fields {
		res += s.field(i)
	}
	res += s.tail()
	return res
}
