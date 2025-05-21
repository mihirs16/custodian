package models

import (
	"time"

	"github.com/google/uuid"
)

type Properties = map[string]any

type Block struct {
	Id         string     `json:"id"`
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	CreatedAt  string     `json:"created_at"`
	UpdatedAt  string     `json:"updated_at"`
}

type BlockOptions struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
}

func CreateBlock(options BlockOptions) Block {
	blockUuid := uuid.New().String()
	block := Block{
		Id:         blockUuid,
		Type:       options.Type,
		Properties: options.Properties,
		CreatedAt:  time.Now().UTC().Format(time.RFC3339),
		UpdatedAt:  time.Now().UTC().Format(time.RFC3339),
	}
	return block
}
