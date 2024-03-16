package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Initialize connection")
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Connection Succesful")
	})
	Server := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	err := Server.ListenAndServe()
	if err != nil {
		panic(err)  
	}
}
