package handler

import (
	"csv-file/internal/storage"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io"
	"net/http"
)

func GetUserName(db *gorm.DB, w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var user struct {
		Name string `json:"name"`
	}

	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(request.Body)
	workers, err := storage.GetWorkerDataByName(db, user.Name)
	if err != nil {
		panic(err)
	}

	for _, worker := range workers {
		fmt.Printf("Имя: %s\n", worker.Name)
		fmt.Printf("Должность: %s\n", worker.JobTitle.Title)
		fmt.Printf("Отдел: %s\n", worker.Department.Title)
		fmt.Printf("Тип занятости: %s\n", worker.FullOrPartTime.Title)
		fmt.Printf("Оплата: %s\n", worker.SalaryOrHourly.Title)
		switch worker.SalaryOrHourly.Title {
		case "HOURLY":
			fmt.Printf("Тип оплаты: Почасовая\nЧасов в неделю: %d\nПочасовая ставка: %f\n",
				worker.WorkerHourlyPayment.TypicalHours, worker.WorkerHourlyPayment.HourlyRate)
		case "SALARY":
			fmt.Printf("Тип оплаты: Помесячная\nГодовая зарплата: %f\n", worker.WorkerSalaryPayment.AnnualSalary)
		default:
			fmt.Println("Неизвестный тип оплаты")
		}
		fmt.Println("---------------------")
	}
}
