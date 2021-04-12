package main

import (
	"ksoc/api"
	"os"

	"github.com/joho/godotenv"
)

var server = api.Server{}

func main() {
	godotenv.Load()
	Port := os.Getenv("PORT")
	server.Run(Port)
}
