package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProductType(t *testing.T) {
	//AAA
	//arrange and act
	p, err := NewProductType("ProductType 1", "Desc 1")

	//assert
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, p.Name, "ProductType 1")
}

func TestProductTypeWhenNameIsRequired(t *testing.T) {
	//AAA
	//arrange and act
	_, err := NewProductType("", "Desc 1")

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrNameIsRequired)
}

func TestProductTypeWhenDescriptionIsRequired(t *testing.T) {
	//AAA
	//arrange and act
	_, err := NewProductType("ProductType 1", "")

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrDescriptionIsRequired)
}
