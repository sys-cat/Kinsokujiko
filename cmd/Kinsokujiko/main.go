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
	e.Logger.Fatal(e.Start(":9090"))
}
