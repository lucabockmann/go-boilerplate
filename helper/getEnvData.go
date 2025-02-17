package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvDataByName(key string) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
		return "Cant Load Env Variable by name: " + key, err
	}

	return os.Getenv(key), nil
}
