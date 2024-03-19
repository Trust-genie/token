package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liezner/token/controller"
)

func Newrouter(controller *controller.Newtagcontroller) *gin.Engine {

	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Connection Succesful")
	})
	baseRouter := router.Group("/api")
	tagRouter := baseRouter.Group("/tags")
	tagRouter.GET("/:tagId", Tagscontroller.Find)
	tagRouter.GET("", tagcontroller.Find_all)
	tagRouter.POST("", tagcontroller.Create)
	tagRouter.PATCH("/:tagId", tagcontroller.Update)
	tagRouter.DELETE("/:tagId", tagcontroller.Delete)

	return router
}
