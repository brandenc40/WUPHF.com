package main

import (
	"net/http"
	"time"

	"github.com/brandenc40/wuphf.com/common"
	"github.com/brandenc40/wuphf.com/config"
	"github.com/brandenc40/wuphf.com/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	if err := config.LoadConfig(); err != nil {
		logger.Error(
			"Unable to load config files",
			zap.Error(err),
		)
	}
	context := common.NewAppContext()
	handlers := handlers.New(context)

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

	// Load HTML and static from React App
	r.LoadHTMLGlob("wuphf-frontend/build/*.html")
	r.Static("/static", "./wuphf-frontend/build/static/")
	r.StaticFile("/manifest.json", "./wuphf-frontend/build/manifest.json")
	r.StaticFile("/favicon.ico", "./wuphf-frontend/build/favicon.ico")
	r.StaticFile("/logo192.png", "./wuphf-frontend/build/logo192.png")
	r.StaticFile("/logo512.png", "./wuphf-frontend/build/logo512.png")
	// send all non API or Auth traffic to the React App
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// API Routes
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Ping")
	})
	r.POST("/wuphf", handlers.WUPHF)

	s.ListenAndServe()
}
