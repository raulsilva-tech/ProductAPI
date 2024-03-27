package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	//AAA
	//arrange
	pt, err := NewProductType("type 1", "type description")
	assert.Nil(t, err)

	//act
	p, err := NewProduct("Product 1", "Desc 1", *pt)

	//assert
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, p.Name, "Product 1")
}

func TestWhenNameIsRequired(t *testing.T) {
	//AAA
	//arrange
	pt, err := NewProductType("type 1", "type description")
	assert.Nil(t, err)

	//act
	_, err = NewProduct("", "Desc 1", *pt)

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrNameIsRequired)
}

func TestWhenDescriptionIsRequired(t *testing.T) {
	//AAA
	//arrange
	pt, err := NewProductType("type 1", "type description")
	assert.Nil(t, err)

	//act
	_, err = NewProduct("Product 1", "", *pt)

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrDescriptionIsRequired)
}
