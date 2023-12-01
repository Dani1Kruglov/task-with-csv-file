package handler

import (
	"csv-file/internal/database"
)

func DataProcessing() error {
	db := database.ConnectToDatabase()
	err := database.StoreIndexByUser(db)
	if err != nil {
		return err
	}
	err = database.AddDataToDatabaseFromCSV(db)
	if err != nil {
		return err
	}
	return nil
}
