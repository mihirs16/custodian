package models

import (
	"context"
	"reflect"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Properties = map[string]any

type Entity struct {
	Id         string     `json:"id"`
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	CreatedAt  string     `json:"created_at"`
	UpdatedAt  string     `json:"updated_at"`
}

type EntityOptions struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
}

type EntityModel struct {
	DB *pgxpool.Pool
}

func (e EntityModel) FetchEntity(id string) (Entity, error) {
	var entity Entity
	err := e.DB.QueryRow(
		context.Background(),
		"SELECT id, type, properties, created_at::text, updated_at::text FROM entities WHERE id = $1",
		id,
	).Scan(
		&entity.Id,
		&entity.Type,
		&entity.Properties,
		&entity.CreatedAt,
		&entity.UpdatedAt,
	)
	if err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (e EntityModel) CreateEntity(entityOpts EntityOptions) (string, error) {
	err := e.validateEntityProps(entityOpts)
	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()
	_, err = e.DB.Exec(
		context.Background(),
		"INSERT INTO entities (\"id\", \"type\", properties) VALUES (@id, @type, @properties);",
		pgx.NamedArgs{
			"id":         uuid,
			"type":       entityOpts.Type,
			"properties": entityOpts.Properties,
		},
	)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

func (e EntityModel) validateEntityProps(entityOpts EntityOptions) error {
	// fetch the type definition to validate against
	var typeDefinition = make(TypeDefinition)
	rows, err := e.DB.Query(
		context.Background(),
		"SELECT field, type, data_type, optional FROM catalogue WHERE type = $1",
		entityOpts.Type,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var field FieldDefinition
		err = rows.Scan(
			&field.Field,
			&field.Type,
			&field.DataType,
			&field.Optional,
		)
		if err != nil {
			return err
		}
		typeDefinition[field.Field] = field
	}

	// validation failed if type not found
	if len(typeDefinition) == 0 {
		return &EntityMissingTypeError{entityOpts}
	}

	// iterate through every property and check against the field definition
	for k, v := range entityOpts.Properties {
		field, fieldExists := typeDefinition[k]
		if !fieldExists {
			return &EntityUnexpectedFieldError{k}
		}

		kindOfValue := reflect.TypeOf(v).Kind()
		kindExpected := KindDataTypeMapping[field.DataType]

		if kindOfValue != kindExpected {
			return &EntityMismatchTypeError{k, field.DataType, kindOfValue}
		}
	}

	// iterate through all field definitions to catch missing non-optional props
	for _, v := range typeDefinition {
		_, fieldExists := entityOpts.Properties[v.Field]
		if !v.Optional && !fieldExists {
			return &EntityMissingFieldError{v.Field}
		}
	}

	return nil
}
