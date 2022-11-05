package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Mock database
var products = []map[int]string{{1: "fridge"}, {2: "AC"}, {3: "Washing-machine"}}

//handlers
func addProducts(c echo.Context) error {
	type payload struct {
		Name string `json:"product_name"`
	}

	var pl payload
	if err := c.Bind(&pl); err != nil {
		log.Fatal(err)
	}
	product := map[int]string{
		len(products) + 1: pl.Name,
	}
	products = append(products, product)
	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {
	// product := map[int]string{}
	for k, _ := range products {
		id, _ := strconv.Atoi(c.Param("id"))
		if k == id {
			return c.JSON(http.StatusOK, products[k])
		}
	}

	return c.JSON(http.StatusNotFound, "No product with given ID")
}

func getAllProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	//routes
	e.GET("/product/:id", getProduct)
	e.POST("/products", addProducts)
	e.GET("/products", getAllProducts)

	e.Logger.Fatal(e.Start("localhost:8000"))
}
