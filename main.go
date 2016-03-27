package main

import (
  "github.com/gin-gonic/gin"
  //"fmt"
)

type GetListName struct {
  Value string `json:"key"`
  //Text string `json:text`
}

func main() {
  router := gin.Default()

  router.POST("/list", func(c *gin.Context) {
    var val GetListName
    c.Bind(&val)
    //c.String(200, val.Value)
    c.JSON(200, val)
  })
  router.Run(":3000")
}

func getBlackList(c *gin.Context) {
  var val GetListName
  c.Bind(&val)
  c.JSON(200, val.Value)
/*  c.JSON(200, gin.H{
    "return":val.Value,
  })*/
}
