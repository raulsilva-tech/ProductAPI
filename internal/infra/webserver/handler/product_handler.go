package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raulsilva-tech/ProductAPI/internal/infra/database"
	"github.com/raulsilva-tech/ProductAPI/internal/usecase"
)

type ProductHandler struct {
	ProductRepository database.ProductRepositoryInterface
}

func NewProductHandler(repo database.ProductRepositoryInterface) *ProductHandler {
	return &ProductHandler{
		ProductRepository: repo,
	}
}

func (h *ProductHandler) Create(c *gin.Context) {

	var input usecase.CreateProductInputDTO
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	uc := usecase.NewCreateProductUseCase(h.ProductRepository)
	output, err := uc.Execute(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (h *ProductHandler) List(c *gin.Context) {

	uc := usecase.NewListProductUseCase(h.ProductRepository)
	outputList, err := uc.Execute(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, outputList)
}

func (h *ProductHandler) GetById(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id is required to get a record"})
		return
	}
	uc := usecase.NewGetProductUseCase(h.ProductRepository)
	found, err := uc.Execute(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, found)
}

func (h *ProductHandler) Update(c *gin.Context) {

	id := c.Param("id")

	var input usecase.UpdateProductInputDTO
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	input.Id = id

	uc := usecase.NewUpdateProductUseCase(h.ProductRepository)
	err = uc.Execute(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.Status(http.StatusOK)

}

func (h *ProductHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id is required to delete a record"})
		return
	}

	uc := usecase.NewDeleteProductUseCase(h.ProductRepository)
	err := uc.Execute(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)

}
