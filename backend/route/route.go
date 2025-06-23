package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sharukh010/url-shortner/controller"
)

func RegisterRoutes(router *gin.Engine){
	router.GET("/test",controller.TestAPI())
	router.POST("/shorturls",controller.CreateLink())
	router.GET("/shorturls/:shortCode",controller.GetLinkStatistics())
	router.GET("/:shortCode",controller.RedirectToLink())
}

