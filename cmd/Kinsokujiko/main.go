package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sys-cat/Kinsokujiko"
)

// Kinsoku is masking sentence
type Kinsoku struct {
	Sentence string `json:"sentence" form:"sentence"`
}

// Item is dictionary Item
type Item struct {
	Surf  string `json:"surf" form:"surf"`
	Slice string `json:"slice" form:"slice"`
	Kana  string `json:"kana" form:"kana"`
	Pos   string `json:"pos" form:"pos"`
}

// Dictionary is Slice any Item
type Dictionary []Item

func main() {
	e := echo.New()
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	e.POST("/", func(c echo.Context) error {
		u := new(Kinsoku)
		if err := c.Bind(u); err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
		kin := Kinsokujiko.Tokenize(Kinsokujiko.Master{Sentence: u.Sentence})
		return c.JSON(http.StatusOK, kin)
	})
	// Dictionary
	e.GET("/dictionary/read", func(c echo.Context) error {
		res, err := Kinsokujiko.Show()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, res)
	})
	e.POST("/dictionary/update", func(c echo.Context) error {
		dics := new(Dictionary)
		if err := c.Bind(dics); err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		var dic Kinsokujiko.Dictionary
		for _, d := range *dics {
			dic = append(dic, Kinsokujiko.Item{Surf: d.Surf, Slice: d.Slice, Kana: d.Kana, Pos: d.Pos})
		}
		res, upErr := Kinsokujiko.Update(dic)
		if upErr != nil {
			return c.JSON(http.StatusInternalServerError, up_err)
		}
		return c.JSON(http.StatusOK, res)
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

	// Start Server
	e.Logger.Fatal(e.Start(":9090"))
}
