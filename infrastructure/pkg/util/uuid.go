package util

import (
	"github.com/google/uuid"
)

// UUID Define alias
type UUID = uuid.UUID

// NewUUID Create uuid
func NewUUID() (UUID, error) {
	return uuid.NewRandom()
}

// MustUUID Create uuid(Throw panic if something goes wrong)
func MustUUID() UUID {
	v, err := NewUUID()
	if err != nil {
		panic(any(err))
	}
	return v
}

// UuidString Create uuid
func UuidString() string {
	return MustUUID().String()
}
