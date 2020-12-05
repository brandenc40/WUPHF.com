package main

import (
	"net/http"
	"os"
	"time"

	"github.com/brandenc40/wuphf.com/common"
	"github.com/brandenc40/wuphf.com/config"
	"github.com/brandenc40/wuphf.com/handlers"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mGin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
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

	// 2000 per 30 days limit ($20 worth)
	rate := limiter.Rate{
		Period: 1 * (time.Hour * 24 * 30),
		Limit:  1000,
	}
	rateLimiter := mGin.NewMiddleware(limiter.New(memory.NewStore(), rate))

	// Build router
	r := gin.Default()

	// Load HTML and static from React App
	r.LoadHTMLGlob("wuphf-frontend/build/*.html")
	r.Static("/static", "./wuphf-frontend/build/static/")
	r.StaticFile("/manifest.json", "./wuphf-frontend/build/manifest.json")
	r.StaticFile("/favicon.ico", "./wuphf-frontend/build/favicon.ico")
	r.StaticFile("/logo192.png", "./wuphf-frontend/build/logo.png")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// Endpoint to check if the app is running properly
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Ping")
	})

	// API Routes
	api := r.Group("/api")
	{
		// Rate limiters
		api.Use(rateLimiter)
		api.POST("/wuphf", handlers.WUPHF)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	r.Run(":" + port)
}
