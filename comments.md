функция calculate должна возвращать ошибку

ответ арифметической операции

сериализуем в json
data, err := json.MarshallIndent(&HistoryElement)

потом отправляем ответ пользователю
w.Write(date)

...

сначала читаем файл history.json
объявляем переменную - массив структур заданного типа
потом в эту переменную записываем десериализуем данные из history.json

hisotry = append(HUstory, newElement)

newElement - это то что мы получили после арифметической операции
history - это то что мы первоначально прочитали из history.json

далее надо history сериализовать
применить marshall indent - навести красоту

далее его надо записать в файл history.json

т.е. перезаписываем файл history

ДАЛЕЕ
следующий роут - get history

открывается файл history json - можно функцию назвать GetDataFromHistory
читается файл
получаем массив байт
потом десериализуется
и отправляется

можно сделать короче - прочитать файл и отправить


...
новый файл server.go в директории internal

type Server struct {
Mux *http.Service
Service string
}

! слоистая архитектура
отдельно сервер (1 слой - хэндлеры) и отдельно сервисы (2 слой - бизнес  логика), 3 слой - работа с базой данных

еще папка в internal - services
там services.go

***
type Service struct {
}
пустая структруа

func NewService() *Service
***

в роутере прописываем
func NewServer(mux *http.ServerMux) *Server{
return &Server{
Mux: mux
Service: service
}


задание - добавить еще один роут
реализовать роут для удаления истории
= записать пустую структуру

функция init сама активизируется при импорте файла

написать роут для удаления истории
по этой новой слоистой архитектуре



