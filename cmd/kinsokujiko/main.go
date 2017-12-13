package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/sys-cat/kinsokujiko"
	kTargets "github.com/sys-cat/kinsokujiko/targets"
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

// Target is Mask Target
type Target struct {
	Surf string `json:"surf" form:"surf"`
	Pos  string `json:"pos" form:"pos"`
	Proc string `json:"proc" form:"proc"`
}

// Targets is Slice any Target
type Targets struct {
	Name    string   `json:"name" form:"name"`       // ターゲット名
	Tag     []string `json:"tag" form:"tag"`         // タグ名リスト
	Targets []Target `json:"targets" form:"targets"` // ターゲットリスト
}

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
		targets := new(Targets)
		if err := c.Bind(targets); err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		var ts kTargets.Targets
		ts.Name = targets.Name
		ts.Tag = strings.Join(targets.Tag, ",")
		for _, t := range *ktargets.Targets {
			ts.Targets = append(ts.Targets, kTargets.Target{Surf: t.Surf, Pos: t.Pos, Proc: t.Proc})
		}
		res, err := kTargets.Create(ts)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, res)
	})
	e.GET("/targets/read", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "/targets/read")
	})
	e.POST("/targets/update", func(c echo.Context) error {
		targets := new(Targets)
		if err := c.Bind(targets); err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		var ts kTargets.Targets
		ts.Name = targets.Name
		ts.Tag = strings.Join(targets.Tag, ",")
		for _, t := range *ktargets.Targets {
			ts.Targets = append(ts.Targets, kTargets.Target{Surf: t.Surf, Pos: t.Pos, Proc: t.Proc})
		}
		res, err := kTargets.Update(ts)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, res)
	})
	e.DELETE("/targets/delete", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "/targets/delete")
	})

	// Start Server
	e.Logger.Fatal(e.Start(":9090"))
}
