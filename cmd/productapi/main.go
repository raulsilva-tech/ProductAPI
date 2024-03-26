package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/raulsilva-tech/ProductAPI/internal/infra/database"
	"github.com/raulsilva-tech/ProductAPI/internal/usecase"
)

func main() {

	ctx := context.Background()

	//creating connection with database
	db, err := sql.Open("sqlite3", "./productapi.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := database.NewProductRepository(db)
	uc := usecase.NewCreateProductUseCase(repo)

	output, err := uc.Execute(ctx, usecase.CreateProductInputDTO{
		Name:        "Product4",
		Description: "product4desc",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
