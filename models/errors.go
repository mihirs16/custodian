package models

import (
	"fmt"
)

type FieldDataTypeError struct {
	fieldOpts FieldDefinitionOpts
}

func (e *FieldDataTypeError) Error() string {
	return fmt.Sprintf("unexpected data type `%s`; ", e.fieldOpts.DataType)
}

// for when there's an unexpected field
type EntityUnexpectedFieldError struct {
	fieldKey string
}

func (e *EntityUnexpectedFieldError) Error() string {
	return fmt.Sprintf("unexpected field `%s`; ", e.fieldKey)
}

// for when a non-optional field is missing
type EntityMissingFieldError struct {
	fieldName string
}

func (e *EntityMissingFieldError) Error() string {
	return fmt.Sprintf("Missing non-optional field `%s`", e.fieldName)
}

// for when a type is not found in the catalogue
type EntityMismatchTypeError struct {
	fieldKey      string
	fieldDataType string
	kindOfValue   string
}

func (e *EntityMismatchTypeError) Error() string {
	return fmt.Sprintf(
		"`%s` expected a `%s` value but got `%s`; ",
		e.fieldKey,
		e.fieldDataType,
		e.kindOfValue,
	)
}

// for when the data type of the value does not
// match the data type expected by the field
type EntityMissingTypeError struct {
	entityOpts EntityOptions
}

func (e *EntityMissingTypeError) Error() string {
	return fmt.Sprintf("type `%s` not found; ", e.entityOpts.Type)
}
