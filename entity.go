package entity

import (
	er "github.com/pierre97t/error"
	uuid "github.com/satori/go.uuid"
)

// Entity is the interface use to handle data on the application
type Entity interface {
	GetID() string
	String() string
	Init(params ...interface{}) *er.Error
	Validate(params ...interface{}) *er.Error
}

// Std is a struct that helps to compose an entity setting default parameter to entity's procedures
type Std struct {
	ID *string `db:"id,omitempty" api:"id,omitempty"`
}

// GetID returns entity's ID
func (s Std) GetID() string {
	if s.ID != nil {
		return *s.ID
	}
	return ""
}

// String cast entity to redable string
func (s Std) String() string {
	return ""
}

// Init initializes default entity's parameters
func (s Std) Init(params ...interface{}) *er.Error {
	return nil
}

// Init initializes default entity's parameters
func (s Std) Validate(params ...interface{}) *er.Error {
	return nil
}

// GenerateID generates a random UUID V4
func GenerateID() string {
	return uuid.NewV4().String()
}
