package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVar() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}else{
		log.Println("Environment variable loading successfull")
	}
}