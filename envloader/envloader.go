package envloader

import (
	"github.com/joho/godotenv"
	"log"
)

func LoaderEnv() {
	envFilePath := "/home/u-andrey/GolandProjects/RestApiProject/local.env"
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
