package entity

import uuid "github.com/satori/go.uuid"

// Entity is the interface use to handle data on the application
type Entity interface {
	GetID() string
	String() string
	Init(params ...interface{}) error
}

// Std is a struct that helps to compose an entity setting default parameter to entity's procedures
type Std struct{}

// GetID retrieves entity's ID
func (s Std) GetID() string {
	return ""
}

// String cast entity to redable string
func (s Std) String() string {
	return ""
}

// Init initializes default entity's parameters
func (s Std) Init(params ...interface{}) error {
	return nil
}

// GenerateID generates a random UUID V4
func GenerateID() string {
	return uuid.NewV4().String()
}
