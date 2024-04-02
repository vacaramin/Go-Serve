package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func Init() {
	setenv()
}
func setenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
