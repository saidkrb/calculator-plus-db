package router

import (
	"github.com/saidkrb/calculator_v2_web.git/config"
	"github.com/saidkrb/calculator_v2_web.git/functions/calculator"
	"github.com/saidkrb/calculator_v2_web.git/functions/history"
	"log"
	"net"
	"net/http"
)

func StartRouter() error {

	getConfig, err := config.GetConfig()
	if err != nil {
		log.Println("Не удалось получить настройки")
		return err
	}

	address := net.JoinHostPort(getConfig.Host, getConfig.Port)

	mux := http.NewServeMux()

	mux.HandleFunc("/calc", calculator.Calculate)
	mux.HandleFunc("/history", history.GetHistory)

	err = http.ListenAndServe(address, mux)
	if err != nil {
		log.Println("Error ListenAndServe: ")
		return err
	}
	return nil
}
