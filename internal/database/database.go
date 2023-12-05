package database

import (
	"csv-file/internal/config"
	"csv-file/internal/model"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	conn := config.Get().DatabaseDSN
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(1000)

	err = db.AutoMigrate(&model.JobTitle{}, &model.Department{}, &model.WorkingDay{}, &model.Payment{}, &model.Worker{}, &model.WorkerHourlyPayment{}, &model.WorkerSalaryPayment{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
