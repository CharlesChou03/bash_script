package main

import (
	"log"

	"github.com/CharlesChou03/_git/testmodule.git/client"
	"github.com/CharlesChou03/_git/testmodule.git/config"
	_ "github.com/CharlesChou03/_git/testmodule.git/docs"
	"github.com/CharlesChou03/_git/testmodule.git/internal/handlers"
	"github.com/CharlesChou03/_git/testmodule.git/logger"
	"github.com/CharlesChou03/_git/testmodule.git/services"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func setup() {
	logger.Setup()
	config.Setup()
	client.Setup(logger.Info)
	services.SetupAuthCache(config.Env.AuthTTL)
}

func errorHandlingMiddleWare(log *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		log.Printf("unexpected error: %s\n", err.Error())
	}
}

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger(), errorHandlingMiddleWare(logger.Error))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", handlers.HealthHandler)
	r.GET("/version", handlers.VersionHandler)

	return r
}

// @title Template Swagger
// @version 0.0.1
// @description this service is golang template
func main() {
	setup()
	r := setupRouter()
	r.Run(":9999")
}
