package main

import (
	"csv-file/internal/handler"
	command "csv-file/internal/work_with_command"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	err := command.WorkWithCSVByCommand()
	if err != nil {
		panic(err)
	}
	fmt.Println("Данные в бд добавлены, индекс по имени создан")
	fmt.Println("http://localhost:8080/user - по ссылке в теле json объектом введите имя (работа через Postman)")
	router := mux.NewRouter()
	router.HandleFunc("/user", handler.GetUserName).Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
