package calculator

import (
	"encoding/json"
	"fmt"
	"github.com/saidkrb/calculator_v2_web.git/internal/models"
	"log"
	"net/http"
	"strconv"
)

func Calculate(w http.ResponseWriter, r *http.Request) error {

	// получаем параметры query из http-запроса
	queries := r.URL.Query()

	queryOne := queries.Get("digit1")
	queryTwo := queries.Get("digit2")
	queryOperator := queries.Get("operator")

	//проверка получения параметров
	fmt.Println(queryOne, queryOperator, queryTwo)

	//переводим параметры в тип данных int
	digitOne, err := strconv.ParseFloat(queryOne, 64)
	if err != nil {
		log.Println(err)
	}

	digitTwo, err := strconv.ParseFloat(queryTwo, 64)
	if err != nil {
		log.Println(err)
	}

	// проверка № 2 получения всех параметров
	fmt.Println(digitOne, queryOperator, digitTwo)

	// проводим арифметическую операцию
	var result float64

	switch queryOperator {
	case "-":
		result = digitOne - digitTwo
	case "+", "%2B":
		result = digitOne + digitTwo
	case "*":
		result = digitOne * digitTwo
	case "/":
		result = digitOne / digitTwo
	}

	// проверка получения результатов арифметической операции
	fmt.Println(result)

	// запись результата в структуру
	var historyElement = models.HistoryElement{
		Digit1:   digitOne,
		Digit2:   digitTwo,
		Operator: queryOperator,
		Result:   result,
	}

	// сериализуем HistoryElement в JSON, чтобы дальше записать ее в файл history.json
	newJSON, err := json.Marshal(historyElement)
	if err != nil {
		log.Println("Ошибка при сериализации полученного результата")
		return err
	}

	// теперь записываем newJSON в файл history.json
	// а поскольку это у нас массив структур
	// то в массив мы записываем через функцию append

	//  но прежде, наведем красоту в JSON
	newJSON, err = json.MarshalIndent(newJSON, "", "")
	if err != nil {
		log.Println("Ошибка при наведении красоты в файле newJSON")
		return err
	}

	// далее записываем newJSON в файл history.json
	//ContentJSON, err := OpenFile.OpenJSON()

	return err
}

/*
	//ВОТ ЗДЕСЬ Я СЕЙЧАС ЗАСТРЯЛ
	ContentJSON = append(ContentJSON, newJSON)

	//  сериализуем данные для передачи в формате json в рамках http-ответа
	jsonResponse, err := json.MarshalIndent(&historyElement, "", "")
	if err != nil {
		log.Println("Error when Marshalling: ", err)
		return
	}

	// направляем ответ на http-запрос
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println("Error when buffering marshall: ", err)
		return
	}

	// записываем из буфера в структуру historyElement

	// записываем historyElement в файл history.json

}
*/
