package models

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type HistoryElement struct {
	Digit1   float64 `json:"digit_1"`
	Digit2   float64 `json:"digit_2"`
	Operator string  `json:"operator"`
	Result   float64 `json:"result"`
}
