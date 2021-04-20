package main

import (
	"os"

	"github.com/singhayushh/kgec-summer-of-code/api"

	"github.com/joho/godotenv"
)

var server = api.Server{}

func main() {
	godotenv.Load()
	Port := os.Getenv("PORT")
	server.Run(Port)
}
