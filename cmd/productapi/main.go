package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	// _ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/mattn/go-sqlite3"
	"github.com/raulsilva-tech/ProductAPI/configs"
	"github.com/raulsilva-tech/ProductAPI/internal/infra/database"
	"github.com/raulsilva-tech/ProductAPI/internal/infra/webserver/handler"
)

func main() {

	//load configuration
	cfg, _ := configs.LoadConfig(".")

	//starting database connection
	DataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	fmt.Println(DataSourceName)

	db, err := sql.Open(cfg.DBDriver, DataSourceName)
	if err != nil {
		panic(err)
	}

	repo := database.NewProductRepository(db)

	pHandler := handler.NewProductHandler(repo)

	//webserver
	router := gin.Default()

	pGroup := router.Group("/products")
	pGroup.POST("/", pHandler.Create)
	pGroup.GET("/", pHandler.List)
	pGroup.POST("/:id", pHandler.Update)
	pGroup.DELETE("/:id", pHandler.Delete)
	pGroup.GET("/:id", pHandler.GetById)

	router.Run(":" + cfg.WebServerPort)
}
