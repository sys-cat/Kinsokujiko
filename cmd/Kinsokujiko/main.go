package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sys-cat/Kinsokujiko"
)

type Kinsoku struct {
	Sentence string `json:"sentence" form:"sentence"`
}

func main() {
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		u := new(Kinsoku)
		if err := c.Bind(u); err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
		kin := Kinsokujiko.Tokenize(Kinsokujiko.Master{u.Sentence})
		return c.JSON(http.StatusOK, kin)
	})
	// Dictionary
	e.GET("/dictionary/read", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "/dictionary/read")
	})
	e.POST("/dictionary/update", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "/dictionary/update")
	})
	// Targets
	e.PUT("/targets/create", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "/targets/create")
	})
	e.GET("/targets/read", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "/targets/read")
	})
	e.POST("/targets/update", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "/targets/update")
	})
	e.DELETE("/targets/delete", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "/targets/delete")
	})
	e.Logger.Fatal(e.Start(":9090"))
}
