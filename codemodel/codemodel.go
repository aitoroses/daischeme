package codemodel

import (
	parser "github.com/daischio/daischeme/codemodel/schemaparser"
	store "github.com/daischio/daischeme/codemodel/schemastore"
	mapper "github.com/daischio/daischeme/codemodel/schemamapper"
	"fmt"
)

type CodeModel struct {
	PackageName string
	ModelName   string
	Scheme      *parser.Scheme
}

/* Return the SchemaStore */
func (c *CodeModel) GetSchemaStore() *store.SchemaStore {
	s := store.New(c.ModelName, c.Scheme)
	return s
}

/* Return the SchemaMapper */
func (c *CodeModel) GetSchemaMapper() *mapper.SchemaMapper {
	mapper := mapper.New(c.ModelName ,c.Scheme)
	return mapper
}

/* Return the SchemaMapper */
func (c *CodeModel) GetCode() string {
	// Obtain the mapped structures
	structs := c.GetSchemaMapper().GetMappedStructs()

	result := fmt.Sprintf("package %s\n\n", c.PackageName)
	for _, s := range structs {
		// Print the struct as code
		result += fmt.Sprintf("\n%s", s.Code())
	}
	return result
}

/* Generate a new CodeModel instance */
func New(p string, md string, schema string) *CodeModel {
	model := CodeModel{p, md, FromSchemaFile(schema)}
	return &model
}

/* Generate a CodeModel for the input schema file */
func FromSchemaFile(p string) *parser.Scheme {
	scheme := parser.ParseSchema(p)
    return scheme
}
