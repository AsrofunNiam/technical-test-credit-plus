package main

import (
	"log"
	"net/http"

	"github.com/AsrofunNiam/technical-test-credit-plus/app"
	c "github.com/AsrofunNiam/technical-test-credit-plus/configuration"
	"github.com/AsrofunNiam/technical-test-credit-plus/helper"
	"github.com/go-playground/validator/v10"
)

func main() {
	configuration, err := c.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	port := configuration.Port
	db := app.ConnectDatabase(configuration.User, configuration.Host, configuration.Password, configuration.PortDB, configuration.Db)

	validate := validator.New()
	router := app.NewRouter(db, validate)
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Printf("Server is running on port %s", port)

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
