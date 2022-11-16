package main

import (
	"github.com/saidkrb/calculator_v2_web.git/internal/router"
	"log"
	"os"
)

func main() {

	err := router.StartRouter()
	if err != nil {
		log.Println("Error StartRouter: ", err)
		os.Exit(1)
	}
}
