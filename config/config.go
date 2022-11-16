package config

import (
	"encoding/json"
	"github.com/saidkrb/calculator_v2_web.git/internal/models"
	"io"
	"log"
	"os"
)

func GetConfig() (*models.Config, error) {

	// чтение и десериализация конфигов
	file, err := os.Open("./config/config.json")
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// создаем переменную config по типу структуры Config
	var config models.Config

	// записываем конфиги в переменную config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Println("Error UnMarshall config: ", err)
		return nil, err
	}
	return &config, nil
}
