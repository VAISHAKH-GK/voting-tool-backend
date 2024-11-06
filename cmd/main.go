package main

import (
	"os"

	"github.com/VAISHAKH-GK/voting-tool-backend/pkg/mongodb"
	"github.com/VAISHAKH-GK/voting-tool-backend/server"
	"github.com/joho/godotenv"
)

func main() {
  var port string
	godotenv.Load(".env")

	mongodb.Connect()

  if port = os.Getenv("PORT"); port != "" {
    port = "9000"
  }

	server.Run(port)
}
