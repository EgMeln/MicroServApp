package main

import (
	"github.com/EgMeln/MicroServApp/pkg/game"
	"github.com/labstack/echo/v4"

	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func main() {
	e := echo.New()
	game.Run()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "HELLO, WORLD")
	})
	e.GET("/users/:id", getUser)
	e.GET("/show", show)

	e.POST("/users", func(context echo.Context) error {
		u := new(User)
		if err := context.Bind(u); err != nil {
			return err
		}
		return context.JSON(http.StatusCreated, u)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
