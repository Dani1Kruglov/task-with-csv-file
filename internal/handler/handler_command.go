package handler

import (
	"csv-file/internal/database"
	"csv-file/internal/work_with_file"
	"fmt"
	"gorm.io/gorm"
)

func DataProcessing(db *gorm.DB) error {
	fmt.Println("Происходит загрузка данных в бд")
	database.CreateIndexForWorkersName(db)
	err := work_with_file.ReadCSVFileAndWriteInDB(db)
	if err != nil {
		return err
	}
	return nil
}
