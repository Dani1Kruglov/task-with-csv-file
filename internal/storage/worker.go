package storage

import (
	"csv-file/internal/model"
	"fmt"
	"gorm.io/gorm"
)

func GetWorkerDataByName(db *gorm.DB, workerName string) ([]model.FullWorker, error) {
	var workers []model.FullWorker
	err := db.Table("workers").Preload("JobTitle").
		Preload("Department").
		Preload("FullOrPartTime").
		Preload("SalaryOrHourly").
		Preload("WorkerHourlyPayment").
		Preload("WorkerSalaryPayment").
		Where("workers.name LIKE ?", "%"+workerName+"%").
		Find(&workers).Error
	if err != nil {
		return nil, err
	}

	if len(workers) == 0 {
		fmt.Println("Рабочие не найдены")
		return nil, nil
	}

	return workers, nil
}
