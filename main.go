package main

import (
	"log"

	"github.com/joho/godotenv"

	"ws_amaris/interfaces/middleware/server"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {

	server.InitServer().RunServer()

}
