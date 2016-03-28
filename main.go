package main

import (
  "github.com/gin-gonic/gin"
  //"fmt"
)

type Mask struct {
  String  string `json:"string"`
  Lists   string `json:"list"`
  Key     string `json:"authorized_key"`
}

func main() {
  router := gin.Default()

  v1 := router.Group("/v1")
  {
    v1.POST("/mask", maskingString)
    v1.GET("/list/:id", getList)
    v1.POST("/list/add", addList)
    v1.POST("/list/edit", editList)
    //v1.GET("/list/del/:del_id", deleteList)
  }
  router.Run(":8080")
}

func maskingString(c *gin.Context) {
  var val Mask
  c.BindJSON(&val)
  c.JSON(200, gin.H{
    "status" : 200,
    "result" : val.String + " : " + val.Lists,
  })
}

func addList(c *gin.Context) {
}

func editList(c *gin.Context) {
}

//func deleteList(c *gin.Context) {
//}

func getList(c *gin.Context) {
}
