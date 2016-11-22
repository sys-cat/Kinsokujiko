package main

import (
  "fmt"
  "log"
  "time"
  "os"
  "github.com/gin-gonic/gin"
  "./mecab"
  "./db"
  "./config"
)

type Mask struct {
  String  string `json:"string"`
  List    []string `json:"list"`
  Key     string `json:"auth"`
}

func init() {
  config.Set()
}

func main() {
  router := gin.Default()

  v1 := router.Group("/v1")
  {
    v1.POST("/mask/", maskingString)
    // list manage statement
    v1.POST("/list/add/", addList)
    v1.POST("/list/edit/", editList)
    v1.GET("/list/:id/", getList)
    v1.GET("/list/:id/del/", deleteList)
    // dictionary manage statement
    v1.POST("/dic/add/", func(c *gin.Context){})
    v1.GET("/dic/get/", func(c *gin.Context){})
    v1.POST("/dic/edit/", func(c *gin.Context){})
    v1.GET("/dic/del/:id/", func(c *gin.Context){})
    v1.GET("/dic/rehash/", func(c *gin.Context){})
    // auth statement
    v1.GET("/get/authorize/key/", getAuthorize)
  }
  router.Run(":8080")
}

func setAccessHeader(c *gin.Context) {
  c.Header("Access-Control-Allow-Origin", "*")
  c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
  c.Header("Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS")
}

func maskingString(c *gin.Context) {
  setAccessHeader(c)
  var val Mask
  c.BindJSON(&val)
  masked, err := mecab.Masking(val.String, val.List)
  if err == nil {
    c.JSON(200, gin.H{"status" : 200,"result" : masked,})
  } else {
    c.JSON(500, gin.H{"status" : 500,"error" : err,})
  }
}

func addList(c *gin.Context) {
  setAccessHeader(c)
  c.JSON(200, gin.H{
    "status" : 200,
    "result" : "nothing"
  })
}

func editList(c *gin.Context) {}

func deleteList(c *gin.Context) {}

func getList(c *gin.Context) {}

func getAuthorize(c *gin.Context) {}
