package diff

import (
	"github.com/getkin/kin-openapi/openapi3"
)

type SchemaCollectionDiff struct {
	AddedSchemas    []string        `json:"addedSchemas,omitempty"`
	DeletedSchemas  []string        `json:"deletedSchemas,omitempty"`
	ModifiedSchemas ModifiedSchemas `json:"modifiedSchemas,omitempty"`
}

func (diff *SchemaCollectionDiff) empty() bool {
	return len(diff.AddedSchemas) == 0 &&
		len(diff.DeletedSchemas) == 0 &&
		len(diff.ModifiedSchemas) == 0
}

func newSchemaDiff() *SchemaCollectionDiff {
	return &SchemaCollectionDiff{
		AddedSchemas:    []string{},
		DeletedSchemas:  []string{},
		ModifiedSchemas: ModifiedSchemas{},
	}
}

type schemaRefPair struct {
	SchemaRef1 *openapi3.SchemaRef
	SchemaRef2 *openapi3.SchemaRef
}

type schemaRefPairs map[string]*schemaRefPair

func diffSchemaCollection(schemas1 openapi3.Schemas, schemas2 openapi3.Schemas) *SchemaCollectionDiff {

	result := newSchemaDiff()

	addedSchemas, deletedSchemas, otherSchemas := diffSchemas(schemas1, schemas2)

	for schema := range addedSchemas {
		result.addAddedSchema(schema)
	}

	for schema := range deletedSchemas {
		result.addDeletedSchema(schema)
	}

	for schema, schemaRefPair := range otherSchemas {
		result.addModifiedSchema(schema, schemaRefPair.SchemaRef1, schemaRefPair.SchemaRef2)
	}

	return result
}

func diffSchemas(schemas1 openapi3.Schemas, schemas2 openapi3.Schemas) (openapi3.Schemas, openapi3.Schemas, schemaRefPairs) {

	added := openapi3.Schemas{}
	deleted := openapi3.Schemas{}
	other := schemaRefPairs{}

	for schemaName1, schemaRef1 := range schemas1 {
		schemaRef2, ok := schemas2[schemaName1]
		if !ok {
			deleted[schemaName1] = schemaRef1
			continue
		}

		other[schemaName1] = &schemaRefPair{
			SchemaRef1: schemaRef1,
			SchemaRef2: schemaRef2,
		}
	}

	for schemaName2, schemaRef2 := range schemas2 {
		_, ok := schemas1[schemaName2]
		if !ok {
			added[schemaName2] = schemaRef2
		}
	}

	return added, deleted, other
}

func (diff *SchemaCollectionDiff) addAddedSchema(schema string) {
	diff.AddedSchemas = append(diff.AddedSchemas, schema)
}

func (diff *SchemaCollectionDiff) addDeletedSchema(schema string) {
	diff.DeletedSchemas = append(diff.DeletedSchemas, schema)
}

func (diff *SchemaCollectionDiff) addModifiedSchema(schema1 string, schemaRef1 *openapi3.SchemaRef, schemaRef2 *openapi3.SchemaRef) {
	diff.ModifiedSchemas.addSchemaDiff(schema1, schemaRef1, schemaRef2)
}

func (diff *SchemaCollectionDiff) getSummary() SchemaDiffSummary {
	return SchemaDiffSummary{
		AddedSchemas:    len(diff.AddedSchemas),
		DeletedSchemas:  len(diff.DeletedSchemas),
		ModifiedSchemas: len(diff.ModifiedSchemas),
	}
}