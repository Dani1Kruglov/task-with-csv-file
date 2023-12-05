package work_with_file

import (
	"csv-file/internal/database"
	"encoding/csv"
	"gorm.io/gorm"
	"os"
	"sync"
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

	var wg sync.WaitGroup
	recordChan := make(chan []string, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go processRecords(db, &wg, recordChan)
	}
	for {
		record, err := reader.Read()
		if err != nil {
			close(recordChan)
			break
		}
		recordChan <- record
	}

	return nil
}

func processRecords(db *gorm.DB, wg *sync.WaitGroup, recordChan <-chan []string) {
	defer wg.Done()

	for record := range recordChan {
		err := database.AddDataToDatabaseFromCSV(db, record)
		if err != nil {
			panic(err)
		}
	}
}
