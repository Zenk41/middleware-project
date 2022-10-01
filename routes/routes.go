package routes

import (
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	c "middleware-project/controllers"
	m "middleware-project/middlewares"
)

func New() *echo.Echo {
	e := echo.New()
 // Middleware 
	m.LogMiddleware(e)
	// Login & Register
	e.POST("/users/register", c.Register)
	e.POST("/users/login", c.Login)
	// Authentication
	privateRoutes := e.Group("")
	privateRoutes.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningKey: []byte("secretToken"),
	}))

	privateRoutes.Use(m.CheckToken)
	// user routing
	privateRoutes.GET("/users", c.GetUsersController)
	privateRoutes.GET("/users/:id", c.GetUserController)
	e.POST("/users", c.CreateUserController)
	privateRoutes.DELETE("/users/:id", c.DeleteUserController)
	privateRoutes.PUT("/users/:id", c.UpdateUserController)

	// book routing
	privateRoutes.GET("/books", c.GetBooksController)
	e.GET("/books/:id", c.GetBookController)
	privateRoutes.POST("/books", c.CreateBookController)
	privateRoutes.DELETE("/books/:id", c.DeleteBookController)
	privateRoutes.PUT("/books/:id", c.UpdateBookController)

	return e
}