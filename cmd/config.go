package cmd

import (
	"http3-integrate/constants"
	"log"

	"github.com/joho/godotenv"
)

func config() {
	// Load env file
	if err := godotenv.Load(); err != nil {
		log.Fatal(constants.EnvLoadErrMsg + err.Error())
	}

	// More configuration
}
