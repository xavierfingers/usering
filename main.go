package main
import (
 "github.com/gin-gonic/gin"
)
var ID = 1
func main() {
  type User struct {
   ID int `json:"id"` 
   Username string `json:"username"`
 } 
  router := gin.Default()
   router.GET("/users/new", func(c *gin.Context) {
  ID += 10               
  c.JSON(200, ID)
 })
 router.GET("/users/ids", func(c *gin.Context) {
  nextID  := ID
  var newUser User
  newUser.ID = nextID
  nextID++ 
  c.JSON(201, nextID)
 })
router.Run(":8080")
}
