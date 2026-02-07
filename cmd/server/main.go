package main

import (
	"log"
	"testeFreteRapido/config"
	"testeFreteRapido/internal/adapter/controllers"
	"testeFreteRapido/internal/adapter/repositories"
	"testeFreteRapido/internal/usecase/services"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	config, err := config.New(r)
	if err != nil {
		log.Fatal("Error to init config: ", err.Error())
	}

	log.Println("Register Repository")
	repository := repositories.NewRepository(config.DB)

	log.Println("Register Service")
	service, err := services.NewService(repository)
	if err != nil {
		log.Fatal("Error to init service: ", err.Error())
	}

	log.Println("Register controllers")
	err = controllers.RegisterControllers(r, service)
	if err != nil {
		log.Fatal("Error registering controllers: ", err.Error())
	}

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Server are not running: ", err.Error())
	}

	log.Println("Server running at http://localhost:8080")
}
