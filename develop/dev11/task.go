package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

/*
	HTTP-сервер

	Реализовать HTTP-сервер для работы с календарем. В рамках задания необходимо работать
	строго со стандартной HTTP-библиотекой.

	В рамках задания необходимо:
	Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные
	функции и объекты доменной области.
	Реализовать middleware для логирования запросов


	Методы API:
	POST /create_event
	POST /update_event
	POST /delete_event
	GET /events_for_day
	GET /events_for_week
	GET /events_for_month

	Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
	В GET методах параметры передаются через queryString, в POST через тело запроса.
	В результате каждого запроса должен возвращаться JSON-документ содержащий либо {"result": "..."}
	в случае успешного выполнения метода, либо {"error": "..."} в случае ошибки бизнес-логики.

	В рамках задачи необходимо:
	Реализовать все методы.
	Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	В случае ошибки бизнес-логики сервер должен возвращать HTTP 503.
	В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400.
	В случае остальных ошибок сервер должен возвращать HTTP 500.
	Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
*/

//Storage структура хранилища ивентов
type Storage struct {
	m      sync.RWMutex
	events map[string][]Event
}

//NewStorage конструктор хранилища
func NewStorage() *Storage {
	return &Storage{
		events: make(map[string][]Event),
	}
}

// Event структура события
type Event struct {
	UserID      string    `json:"User_id"`
	Date        time.Time `json:"date"`
	ID          string    `json:"event_id"`
	Description string    `json:"description"`
}

func main() {
	storage := NewStorage()
	server := &http.Server{
		Addr:         "localhost:8080",
		Handler:      storage.NewRouter(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server start")
	log.Fatalln(server.ListenAndServe())
}

// NewRouter конструктор хендлеров
func (s *Storage) NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", myMiddleware(s.create))
	mux.HandleFunc("/update_event", myMiddleware(s.update))
	mux.HandleFunc("/delete_event", myMiddleware(s.delete))
	mux.HandleFunc("/events_for_day", myMiddleware(s.dayEvent))
	mux.HandleFunc("/events_for_week", myMiddleware(s.weekEvent))
	mux.HandleFunc("/events_for_month", myMiddleware(s.monthEvent))
	return mux
}

// myMiddleware логирует запросы
func myMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request method " + r.Method + ", url " + r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

// create обрабатывает create_event
func (s *Storage) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("Invalid method: " + r.Method)
		http.Error(w, "Invalid method", 400)
		return
	}
	event, err := parsePost(r)
	if err != nil {
		fmt.Println(err)
		errorResp(w, err.Error(), 400)
		return
	}
	err = Validate(event)
	if err != nil {
		fmt.Println(err)
		errorResp(w, err.Error(), 400)
		return
	}
	err = s.createEvent(event)
	if err != nil {
		fmt.Println(err)
		errorResp(w, err.Error(), 400)
		return
	}
	resultResp(w, s.events[event.UserID])
}

// createEvent создает ивент
func (s *Storage) createEvent(event *Event) error {
	s.m.Lock()
	defer s.m.Unlock()
	events := s.events[event.UserID]
	for _, e := range events {
		if e.ID == event.ID {
			return fmt.Errorf("Event already exists")
		}
	}
	s.events[event.UserID] = append(s.events[event.UserID], *event)
	return nil
}

//update обрабатывает update_event
func (s *Storage) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("Invalid method: " + r.Method)
		http.Error(w, "Invalid method", 400)
		return
	}
	event, err := parsePost(r)
	if err != nil {
		fmt.Println(err)
		errorResp(w, err.Error(), 400)
		return
	}
	err = Validate(event)
	if err != nil {
		fmt.Println(err)
		errorResp(w, err.Error(), 400)
		return
	}
	err = s.updateEvent(event)
	if err != nil {
		fmt.Println(err)
		errorResp(w, err.Error(), 400)
		return
	}
	resultResp(w, s.events[event.UserID])
}

// updateEvent обновляет ивент
func (s *Storage) updateEvent(event *Event) error {
	s.m.Lock()
	defer s.m.Unlock()
	update := false
	events := s.events[event.UserID]
	for i, v := range events {
		if v.ID == event.ID {
			events[i].Date = event.Date
			events[i].Description = event.Description
			update = true
			break
		}
	}
	if update == false {
		return fmt.Errorf("Event does not exist")
	}
	return nil
}

//delete обрабатывает delete_event
func (s *Storage) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("Invalid method: " + r.Method)
		http.Error(w, "Invalid method", 400)
		return
	}
	event, err := parsePost(r)
	if err != nil {
		fmt.Println(err)
		errorResp(w, err.Error(), 400)
		return
	}
	err = Validate(event)
	if err != nil {
		fmt.Println(err)
		errorResp(w, err.Error(), 400)
		return
	}
	err = s.deleteEvent(event)
	if err != nil {
		fmt.Println(err)
		errorResp(w, err.Error(), 400)
		return
	}
	resultResp(w, s.events[event.UserID])
}

//deleteEvent удаляет ивент
func (s *Storage) deleteEvent(event *Event) error {
	s.m.Lock()
	defer s.m.Unlock()
	events := s.events[event.UserID]
	update := false
	for i, v := range events {
		if v.ID == event.ID {
			events = append(events[:i], events[i+1:]...)
			update = true
			break
		}
	}
	if update == false {
		return fmt.Errorf("Event does nor exist")
	}
	s.events[event.UserID] = events
	return nil
}

//dayEvent обрабатывает events_for_day
func (s *Storage) dayEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("Invalid method: " + r.Method)
		http.Error(w, "Invalid method", 400)
		return
	}
	event, err := parseGet(r)
	if err != nil {
		fmt.Println(err)
		errorResp(w, err.Error(), 400)
		return
	}
	events, err := s.GetForDay(event)
	if err != nil {
		fmt.Println(err)
		errorResp(w, err.Error(), 400)
		return
	}
	//	fmt.Println(events)
	resultResp(w, events)
}

//GetForDay получаем ивенты на день
func (s *Storage) GetForDay(event *Event) ([]Event, error) {
	s.m.Lock()
	defer s.m.Unlock()
	result := make([]Event, 0)
	events, ok := s.events[event.UserID]
	if !ok {
		return nil, fmt.Errorf("User does not exists")
	}
	for _, v := range events {
		if v.Date.Year() == event.Date.Year() && v.Date.Month() == event.Date.Month() && v.Date.Day() == event.Date.Day() {
			result = append(result, v)
		}
	}

	return result, nil
}

//weekEvent обрабатывает events_for_week
func (s *Storage) weekEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("Invalid method: " + r.Method)
		http.Error(w, "Invalid method", 400)
		return
	}
	event, err := parseGet(r)
	if err != nil {
		fmt.Println(err)
		errorResp(w, err.Error(), 400)
		return
	}
	events, err := s.GetForWeek(event)
	if err != nil {
		fmt.Println(err)
		errorResp(w, err.Error(), 400)
		return
	}
	resultResp(w, events)
}

//GetForWeek получаем ивенты на неделю
func (s *Storage) GetForWeek(event *Event) ([]Event, error) {
	s.m.Lock()
	defer s.m.Unlock()
	result := make([]Event, 0)
	events, ok := s.events[event.UserID]
	yearEvent, weekEvent := event.Date.ISOWeek()
	if !ok {
		return nil, fmt.Errorf("User does not exists")
	}
	for _, v := range events {
		year, week := v.Date.ISOWeek()
		fmt.Println(year, week, yearEvent, weekEvent)
		if week == weekEvent && year == yearEvent {

			result = append(result, v)
		}
	}
	return result, nil
}

//monthEvent обрабатывает events_for_month
func (s *Storage) monthEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("Invalid method: " + r.Method)
		http.Error(w, "Invalid method", 400)
		return
	}
	event, err := parseGet(r)
	if err != nil {
		fmt.Println(err)
		errorResp(w, err.Error(), 400)
		return
	}
	events, err := s.GetForMonth(event)
	if err != nil {
		fmt.Println(err)
		errorResp(w, err.Error(), 400)
		return
	}
	resultResp(w, events)
}

//GetForMonth получаем ивенты на месяц
func (s *Storage) GetForMonth(event *Event) ([]Event, error) {
	s.m.Lock()
	defer s.m.Unlock()
	result := make([]Event, 0)
	events, ok := s.events[event.UserID]
	if !ok {
		return nil, fmt.Errorf("User does not exists")
	}
	for _, v := range events {
		if v.Date.Year() == event.Date.Year() && v.Date.Month() == event.Date.Month() {
			result = append(result, v)
		}
	}
	return result, nil
}

//parsePost парсим параметры пост запроса
func parsePost(r *http.Request) (*Event, error) {
	event := Event{}

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

//parseGet пармис параметры гет запроса
func parseGet(r *http.Request) (*Event, error) {
	userID := r.FormValue("user_id")
	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		return nil, err
	}
	event := Event{
		UserID: userID,
		Date:   date,
	}
	return &event, nil
}

//Validate проверяем валидность полученных данных
func Validate(event *Event) error {
	if event.UserID == "" {
		return fmt.Errorf("Invalid user id: %s", event.UserID)
	}
	if event.ID == "" {
		return fmt.Errorf("Invalid event id: %s", event.ID)
	}
	return nil
}

//errorResp возвращаем ошибку
func errorResp(w http.ResponseWriter, e string, code int) {
	errorResponse := struct {
		Error string `json:"error"`
	}{e}

	json, _ := json.Marshal(errorResponse)
	http.Error(w, string(json), code)
}

//resultResp возвращаем результат обработки запроса
func resultResp(w http.ResponseWriter, r []Event) {
	resultResponse := struct {
		Result []Event `json:"result"`
	}{r}

	json, err := json.Marshal(resultResponse)
	if err != nil {
		errorResp(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(json)
}
