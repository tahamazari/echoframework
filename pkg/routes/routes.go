package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/tahamazari/echo-framework/pkg/controllers"
)

func SetupBookRoutes(e *echo.Echo) {
	e.GET("/book/:id", controllers.GetBook)
	e.POST("/book", controllers.CreateBook)
	e.PUT("/book/:id", controllers.UpdateBook)
	e.DELETE("/book/:id", controllers.DeleteBook)

	e.GET("/books", controllers.GetBooks)
}
