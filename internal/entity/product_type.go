package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type ProductType struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   sql.NullTime
}

func NewProductType(name, description string) (*ProductType, error) {
	p := &ProductType{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		CreatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	}
	return p, p.Valid()
}

func (p *ProductType) Valid() error {

	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Description == "" {
		return ErrDescriptionIsRequired
	}

	return nil
}
