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

type TypeDefinition = []FieldDefinition

type CatalogueModel struct {
	DB *pgxpool.Pool
}

func (c CatalogueModel) FetchType(typeName string) (TypeDefinition, error) {
	var typeDefinition TypeDefinition
	rows, err := c.DB.Query(
		context.Background(),
		"SELECT field, type, data_type, optional FROM catalogue WHERE type = $1",
		typeName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var field FieldDefinition
		err = rows.Scan(&field.Field, &field.Type, &field.DataType, &field.Optional)
		if err != nil {
			return nil, err
		}
		typeDefinition = append(typeDefinition, field)
	}

	return typeDefinition, nil
}

func (c CatalogueModel) CreateField(field FieldDefinition) error {
	_, err := c.DB.Exec(
		context.Background(),
		"INSERT INTO catalogue (field, type, data_type, optional) VALUES (@field, @type, @data_type, @optional)",
		pgx.NamedArgs{
			"field":     field.Field,
			"type":      field.Type,
			"data_type": field.DataType,
			"optional":  field.Optional,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
