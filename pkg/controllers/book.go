package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tahamazari/echo-framework/pkg/config"
	"github.com/tahamazari/echo-framework/pkg/models"
)

func GetBook(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()
	fmt.Println(id)

	fmt.Println((id))

	var book models.Book // Declare a single instance of model.Book

	if res := db.First(&book, id); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}
		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": book,
	}

	return c.JSON(http.StatusOK, response)
}

func CreateBook(c echo.Context) error {
	b := new(models.Book)
	db := config.DB()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	book := &models.Book{
		Name:   b.Name,
		Author: b.Author,
	}

	if err := db.Create(&book).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": book,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")
	b := new(models.Book)
	db := config.DB()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existing_book := new(models.Book)

	if err := db.First(&existing_book, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existing_book.Name = b.Name
	existing_book.Author = b.Author
	if err := db.Save(&existing_book).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existing_book,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	book := new(models.Book)

	if err := db.First(&book, id).Error; err != nil {
		data := map[string]interface{}{
			"message": "Book not found",
		}
		return c.JSON(http.StatusNotFound, data)
	}

	if err := db.Delete(&book).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Book has been soft deleted",
	}
	return c.JSON(http.StatusOK, response)
}

func GetBooks(c echo.Context) error {
	db := config.DB()

	var books []models.Book

	if err := db.Find(&books).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	return c.JSON(http.StatusOK, books)
}
