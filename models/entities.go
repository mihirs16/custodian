package models

import (
	"context"

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
		"SELECT id, type, properties FROM entities WHERE id = $1",
		id,
	).Scan(
		&entity.Id,
		&entity.Type,
		&entity.Properties,
	)
	if err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (e EntityModel) CreateEntity(entityOpts EntityOptions) (string, error) {
	uuid := uuid.New().String()
	_, err := e.DB.Exec(
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
