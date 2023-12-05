package database

import (
	"csv-file/internal/model"
	"gorm.io/gorm"
	"reflect"
	"strconv"
)

func CreateIndexForWorkersName(db *gorm.DB) {
	db.Exec("CREATE INDEX IF NOT EXISTS name_lower_idx ON workers ((lower(name)));")
}

func AddDataToDatabaseFromCSV(db *gorm.DB, record []string) error {
	jobTitleID, _ := addRecord(db, &model.JobTitle{}, record[1])
	departmentID, _ := addRecord(db, &model.Department{}, record[2])
	workingDayID, _ := addRecord(db, &model.WorkingDay{}, record[3])
	paymentID, _ := addRecord(db, &model.Payment{}, record[4])
	worker := model.Worker{
		Name:             record[0],
		JobTitlesID:      jobTitleID,
		DepartmentID:     departmentID,
		FullOrPartTimeID: workingDayID,
		SalaryOrHourlyID: paymentID,
	}
	db.Create(&worker)
	if record[4] == "SALARY" {
		salary, err := strconv.ParseFloat(record[6], 2)
		if err != nil {
			return err
		}
		db.Create(&model.WorkerSalaryPayment{
			WorkerID:     worker.ID,
			AnnualSalary: salary,
		})
	} else if record[4] == "HOURLY" {
		hours, err := strconv.Atoi(record[5])
		rate, err := strconv.ParseFloat(record[7], 2)
		if err != nil {
			return err
		}
		db.Create(&model.WorkerHourlyPayment{
			WorkerID:     worker.ID,
			TypicalHours: hours,
			HourlyRate:   rate,
		})
	}
	return nil
}

func addRecord(db *gorm.DB, model interface{}, title string) (uint, error) {
	record := map[string]interface{}{"title": title}
	result := db.FirstOrCreate(model, record)
	if result.Error != nil {
		return 0, result.Error
	}
	return uint(reflect.ValueOf(model).Elem().FieldByName("ID").Uint()), nil
}
