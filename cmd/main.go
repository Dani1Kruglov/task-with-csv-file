package main

import (
	"csv-file/internal/database"
	"csv-file/internal/handler"
	command "csv-file/internal/work_with_command"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	db, err := database.Connect()
	fmt.Println("Выберите, что хотите сделать: \n 1 - Загрузить данные в бд и сделать поиск \n 2 - Сделать поиск")
	var com string
	_, err = fmt.Scan(&com)
	if err != nil {
		panic(err)
	}
	switch com {
	case "1":
		err = command.WorkWithCSVByCommand(db)
		if err != nil {
			panic(err)
		}
		fmt.Println("Данные в бд добавлены, индекс по имени создан")
		startingServer(db)
		break
	case "2":
		startingServer(db)
		break
	default:
		fmt.Println("Такой команды нет")

	}
}

func startingServer(db *gorm.DB) {
	fmt.Println("http://localhost:8080/user - по ссылке в теле json объектом введите имя (работа через Postman)")
	router := mux.NewRouter()
	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		handler.GetUserName(db, w, r)
	}).Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
