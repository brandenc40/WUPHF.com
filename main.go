package main

import (
	"github.com/brandenc40/wuphf.com/config"
	"github.com/brandenc40/wuphf.com/controllers"
)

func main() {
	config.LoadConfig()

	controller := controllers.New()
	params := controllers.WuphfParams{
		Message:   "test",
		SMSNumber: "+1 563-343-5557",
	}
	_ = controller.SendWuphf(&params)

	// r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}