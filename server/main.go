package main

import (
	"net/http"

	"github.com/code-injection/quizaction/server/internal/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//	@title			Quizaction Server
//	@version		0.0
//	@description	Quizaction Server.
//	@schemes		http
//	@host			localhost:8080

// TODO: Separate server, routes, etc. initializations
// TODO: Add config

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	ur := user.NewInMemRepo()
	us := user.NewService(ur)
	uh := user.NewHandler(us)

	docs := e.Group("/docs")
	docs.File("/spec", "docs/swagger.yaml")
	docs.File("", "docs/docs.html")

	e.GET("/hello", Hello)

	e.GET("/users/:username", uh.GetUser)
	e.POST("/users", uh.CreateUser)

	e.Logger.Fatal(e.Start(":80"))
}

// Example Hello Response DTO. Will be gone tomorrow... 🌨️
type HelloResponse struct {
	Message string `json:"message"`
}

// Hello always returns "Quiz, quiz, go"
//
//	@Summary	Hello
//	@Tags		Others
//	@Id			hello
//	@Produce	json
//	@Success	200	{object}	HelloResponse
//	@Router		/hello [get]
func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, HelloResponse{Message: "Quiz, quiz, go"})
}
