package codemodel

import (
	parser "github.com/daischio/daischeme/codemodel/schemaparser"
	store "github.com/daischio/daischeme/codemodel/schemastore"
	mapper "github.com/daischio/daischeme/codemodel/schemamapper"
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
