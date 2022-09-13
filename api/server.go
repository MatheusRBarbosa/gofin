package api

import (
	"github.com/gin-gonic/gin"
	utils "github.com/matheusrbarbosa/gofin/api/v1/utils"
)

func StartHttpServer() {
	router := gin.Default()

	router.GET("ping", utils.Ping)

	router.Run(":8080")
}