package models

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FieldDefinition struct {
	Type      string `json:"type"`
	Field     string `json:"field"`
	DataType  string `json:"data_type"`
	Optional  bool   `json:"optional"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type FieldDefinitionOpts struct {
	Type     string `json:"type"`
	Field    string `json:"field"`
	DataType string `json:"data_type"`
	Optional bool   `json:"optional"`
}

type TypeDefinition = map[string]FieldDefinition

type CatalogueModel struct {
	DB *pgxpool.Pool
}

var KindDataTypeMapping = map[string]string{
	"text":    "string",
	"number":  "float64",
	"boolean": "bool",
	"array":   "slice",
}

func (c CatalogueModel) FetchType(typeName string) (TypeDefinition, error) {
	var typeDefinition = make(TypeDefinition)
	rows, err := c.DB.Query(
		context.Background(),
		"SELECT field, type, data_type, optional, created_at::text, updated_at::text FROM catalogue WHERE type = $1",
		typeName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var field FieldDefinition
		err = rows.Scan(
			&field.Field,
			&field.Type,
			&field.DataType,
			&field.Optional,
			&field.CreatedAt,
			&field.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		typeDefinition[field.Field] = field
	}

	return typeDefinition, nil
}

func (c CatalogueModel) CreateField(fieldOpts FieldDefinitionOpts) error {
	_, dataTypeMappingExists := KindDataTypeMapping[fieldOpts.DataType]
	if !dataTypeMappingExists {
		return &FieldDataTypeError{fieldOpts}
	}

	_, err := c.DB.Exec(
		context.Background(),
		"INSERT INTO catalogue (field, type, data_type, optional) VALUES (@field, @type, @data_type, @optional)",
		pgx.NamedArgs{
			"field":     fieldOpts.Field,
			"type":      fieldOpts.Type,
			"data_type": fieldOpts.DataType,
			"optional":  fieldOpts.Optional,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
