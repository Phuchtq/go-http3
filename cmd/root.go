package cmd

import (
	"fmt"
	"http3-integrate/constants"
	"http3-integrate/infrastructures/routes"
	"log"
	"os"

	"github.com/quic-go/quic-go/http3"
	"github.com/spf13/cobra"
)

const (
	backUpPetApiPort string = "Your backup api port"
)

func setUp() {
	var port string = os.Getenv(constants.PetApiPort)
	var service string = "Pet"

	if port == "" {
		// Log message to inform that env variable has not been set
		log.Println(fmt.Sprintf(constants.EnvApiPortNotSetMsg, service))

		// Set env variable
		os.Setenv(constants.PetApiPort, backUpPetApiPort)
		log.Println(fmt.Sprintf(constants.EnvApiPortSetMsg, service) + backUpPetApiPort)

		port = backUpPetApiPort
	}

	var server = &http3.Server{
		Addr:      port,
		Handler:   routes.InitializePetApi(),
		TLSConfig: generateTlsConfig(),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(fmt.Sprintf(constants.ApiInitErrMsg, service, port) + err.Error())
	}

	log.Println("Pet service starts on port: ", port)
}

var rootCmd = &cobra.Command{
	Use:     "your-service",
	Short:   "your brief description of this service.",
	Aliases: []string{"command 1", "command 2", "command 3"}, // alternative commands
	Run: func(cmd *cobra.Command, args []string) { // Start command
		log.Println("Run service")

		// Load configuration
		config()

		// Set up service
		setUp()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(fmt.Sprintf(constants.CmdExecuteErrMsg, "Pet") + err.Error())
	}
}