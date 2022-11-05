package main

import (
	"github.com/arunsingh/echo/CRUD/cmd/handlers"
	"github.com/labstack/echo"
)

// CRUD APIs for products

func main() {
	e := echo.New()
	e.Static("/", "../index.html")

	//routes
	e.POST("/products/add", handlers.AddProducts).Name = "arun" // giving names to each routes #inline-naming
	e.GET("/products", handlers.GetAllProducts)
	e.GET("/products/:id", handlers.GetProduct)
	e.PUT("/products/:id", handlers.UpdateProduct)
	e.DELETE("/products/:id", handlers.DeleteProduct)
	e.POST("/upload", handlers.Upload)

	e.Logger.Fatal(e.Start("localhost:1900"))

}
