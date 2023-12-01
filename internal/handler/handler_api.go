package handler

import (
	"csv-file/internal/database"
	"csv-file/internal/model"
	"csv-file/internal/storage"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUserName(w http.ResponseWriter, request *http.Request) {
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
	db := database.ConnectToDatabase()

	userJSON, err := storage.GetUserByNameFromDB(user.Name, db)
	if err != nil {
		panic(err)
	}

	var userData model.User
	err = json.Unmarshal(userJSON, &userData)
	if err != nil {
		panic(err)
	}

	fmt.Println(userData)
}
