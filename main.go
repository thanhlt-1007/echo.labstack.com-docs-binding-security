package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CreateUserRequest struct {
	Email string `json:"email" form:"email"`
	Name  string `json:"name" form:"name"`
}

type User struct {
	Email   string
	Name    string
	IsAdmin bool
}

func postCreateUser(context echo.Context) error {
	var request CreateUserRequest
	err := context.Bind(&request)
	if err != nil {
		return context.JSON(
			http.StatusBadRequest,
			err,
		)
	}

	user := User{
		Email:   request.Email,
		Name:    request.Name,
		IsAdmin: false,
	}

	return context.JSON(
		http.StatusOK,
		user,
	)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/users", postCreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}
