package main

import (
	"log"
	"varcelio/controller"
	"varcelio/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	router := gin.Default()

	router.Use(gzip.Gzip(gzip.BestSpeed))
	router.Use(cors.Default())

	basePath := "/api/"
	docs.SwaggerInfo.BasePath = basePath
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.StaticFile("/rapidoc.html", "docs/rapidoc.html")

	apiV1 := router.Group(basePath + "v1")
	apiV1.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Imusegipo"}) })

	controller.NewUserController(apiV1)

	router.Run(":8080")
}
