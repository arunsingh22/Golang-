package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/arunsingh/echo/CRUD/internal/model"
	"github.com/labstack/echo"
)

func Upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
}

//handlers
func AddProducts(c echo.Context) error {
	product := model.Product{}
	// POST -> BODY
	err := json.NewDecoder(c.Request().Body).Decode(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to parse input")
	}
	model.DB = append(model.DB, product)
	return c.JSON(http.StatusAccepted, "Product added to inventory")
}

func GetAllProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, model.DB)
}
func GetProduct(c echo.Context) error {
	id := c.Param("id")
	for _, prod := range model.DB {
		if id, err := strconv.Atoi(id); err == nil {
			if prod.ID == id {
				return c.JSON(http.StatusOK, prod)
			}

		}
	}
	return c.JSON(http.StatusNotFound, "No such product in Inventory")
}

func UpdateProduct(c echo.Context) error {
	return nil
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	for k, prod := range model.DB {
		if id, err := strconv.Atoi(id); err == nil {
			if prod.ID == id {
				model.DB = append(model.DB[:k], model.DB[k+1:]...)
				return c.JSON(http.StatusOK, model.DB)
			}
		}
	}
	return c.JSON(http.StatusNoContent, "No prodcut in Inventory with given ID")
}
