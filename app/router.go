package app

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("/lottery", SaveTicket)
		v1.GET("/lottery", QueryTicket)
	}
}
