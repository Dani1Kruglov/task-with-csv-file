package work_with_file

import (
	"csv-file/internal/database"
	"encoding/csv"
	"gorm.io/gorm"
	"os"
)

func ReadCSVFileAndWriteInDB(db *gorm.DB) error {
	file, err := os.Open("data.csv")
	if err != nil {
		panic("Error opening file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		panic("Error skipping first row")
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				return err
			}
		}
		err = database.AddDataToDatabaseFromCSV(db, record)
		if err != nil {
			return err
		}
	}
	return nil
}
