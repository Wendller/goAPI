package entities

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(plainText string) (ID, error) {
	id, err := uuid.Parse(plainText)

	return ID(id), err
}
