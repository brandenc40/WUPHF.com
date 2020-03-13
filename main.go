package main

import (
	"net/http"
	"time"

	"github.com/brandenc40/wuphf.com/config"
	"github.com/brandenc40/wuphf.com/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	handlers := handlers.New()

	// controller := controllers.New()
	// params := controllers.WuphfParams{
	// 	Message:    "Decided to sell company. Thanks, bro. Hell of a ride.",
	// 	FromName:   "Ryan Howard",
	// 	SMSNumber:  "+1 563-343-5557",
	// 	CallNumber: "+1 563-343-5557",
	// 	ToEmail:    "brandencolen@gmail.com",
	// }
	// _ = controller.SendWuphf(&params)

	r := gin.Default()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Ping")
	})
	r.POST("/wuphf", handlers.WUPHF)

	s.ListenAndServe()
}
