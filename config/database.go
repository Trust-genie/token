package config

import (
	"fmt"

	"github.com/liezner/token/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	server   = "localhost"
	database = "postgres"
	port     = 5432
	username = "postgres"
	password = "cabineT123"
)

func DatabaseConnection() *gorm.DB {
	sqlcredentials := fmt.Sprintf("server=%s database=%s port=%d username=%s password=%s sslmode=disable", server, database, port, username, password)
	db, err := gorm.Open(postgres.Open(sqlcredentials), &gorm.Config{})
	helper.ErrorPanic(err)

	return db

}
