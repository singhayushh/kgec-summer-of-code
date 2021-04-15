package main

import (
	"os"

	"github.com/singhayushh/ksoc/api"

	"github.com/joho/godotenv"
)

var server = api.Server{}
var g = api.GitHubAPI{}

func main() {
	godotenv.Load()
	Port := os.Getenv("PORT")
	g.Init()
	server.Run(Port)
}
