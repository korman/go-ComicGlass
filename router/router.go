package router

import (
	"github.com/gin-gonic/gin"
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
		getTitleList
	}

	router.Run(":8111")
}
