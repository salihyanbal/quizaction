package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//	@title			Quizaction Server
//	@version		0.0
//	@description	Quizaction Server.
//	@host			localhost:8080

func main() {
	e := echo.New()

	docs := e.Group("/docs")
	docs.File("/spec", "docs/swagger.yaml")
	docs.File("", "docs/docs.html")

	e.GET("/hello", Hello)

	e.Logger.Fatal(e.Start(":80"))
}

//	@Summary	Hello
//	@Tags		Others
//	@Produce	json
//	@Success	200
//	@Router		/hello [get]
func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Quiz, quiz, Go!"})
}
