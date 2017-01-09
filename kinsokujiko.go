package kinsokujiko

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatucOK, "hello world")
    })
    e.Logger.Fatal(e.Start(":9090"))
}