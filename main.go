package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/liezner/token/config"
	"github.com/liezner/token/controller"
	"github.com/liezner/token/models"
	"github.com/liezner/token/repository"
	"github.com/liezner/token/router"
	"github.com/liezner/token/service"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Initialize connection")

	DB := config.DatabaseConnection()
	validate := validator.New()

	DB.Table(tags).AutoMigrate(&models.Tags{})

	tagrepo := repository.NewTagsRepo(DB)

	tagservice := service.NewTagServiceimpl(tagrepo, validate)

	tagcontroller := controller.Newtagcontroller(tagservice)

	routes := router.Newrouter(&tagcontroller)
	Server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := Server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
