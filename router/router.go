package router

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/korman/go-ComicGlass/handler"
	"time"
)

func Init() {
	router := gin.New()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, X-Token",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	getTitleList := router.Group("/")
	{
		getTitleList.GET("index", handler.Index)
	}

	router.Run(":8111")
}
