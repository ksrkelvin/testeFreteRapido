package config

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

func New(r *gin.Engine) (config *Config, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(
				"Erro to try NewFreteRapidoApi: %v\n\nstack trace:\n%s",
				r,
				debug.Stack(),
			)
		}
	}()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Env are not loaded: %v", err.Error())
	}

	c := &Config{}
	log.Println("Connect DB")
	if err = c.ConnectDB(); err != nil {
		log.Fatalf("DB cannot to be initialized: %v", err.Error())
		return
	}

	return c, err
}
