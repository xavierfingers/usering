package main
import (
 "github.com/gin-gonic/gin"
)

func main() {
  type User struct {
   ID int `json:"id"` 
   Username string `json:"username"`
 } 
  router := gin.Default()
   router.GET("/index.html", func(c *gin.Context) {
    c.Redirect(307, "/index")
   })
   router.GET("/teapot", func(c *gin.Context) {
      c.String(418, "I'm a teapot not a coffee pot.")
  })
   router.GET("/index", func(c *gin.Context) {
     c.String(200, "Hello, World! go to /users/new to create a new user and go to /users/ids for user ids.")
   })
   router.GET("/users/new", func(c *gin.Context) {
  ID := c.Query("id")           
  c.JSON(200, ID)
 })
 router.GET("/users/ids", func(c *gin.Context) {
  nextID  := 1
  nextID += 10
  c.JSON(201, nextID)
 })
router.Run(":8080")
}
