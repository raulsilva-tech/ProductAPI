package entity

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrDescriptionIsRequired = errors.New("description is required")
	ErrNameIsRequired        = errors.New("name is required")
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   sql.NullTime
}

func NewProduct(name, description string) (*Product, error) {
	p := &Product{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		CreatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	}
	return p, p.Valid()
}

func (p *Product) Valid() error {

	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Description == "" {
		return ErrDescriptionIsRequired
	}

	return nil
}
