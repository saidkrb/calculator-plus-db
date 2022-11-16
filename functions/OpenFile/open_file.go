package OpenFile

import (
	"io"
	"log"
	"os"
)

func OpenJSON() ([]byte, error) {
	// открываем файл с историей - он сейчас в формате json
	file, err := os.OpenFile("./history/history.json", os.O_RDWR, 0777)
	if err != nil {
		log.Println("could not open history file")
		return nil, err
	}

	// читаем фйл (он все еще в json) и записываем его в переменную contentJSON
	ContentJSON, err := io.ReadAll(file)
	if err != nil {
		log.Println("could not read history file")
		return nil, err
	}
	return ContentJSON, nil
}
