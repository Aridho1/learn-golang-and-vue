package main

import (
	"zunn/backend-api/config"
	"zunn/backend-api/database"
	"zunn/backend-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.InitDB()

	// router := gin.Default()

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message":	 "Hello world",
	// 	})
	// })
	
	// port := config.GetEnv("APP_PORT", "3000")
	// log.Printf("Server starting on port %s", port)
	// router.Run(":" + port)
	
	r := routes.SetupRouter()
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":	 "Hello world",
		})
	})

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
