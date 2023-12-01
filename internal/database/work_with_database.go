package database

import (
	"database/sql"
	"path/filepath"
)

func AddDataToDatabaseFromCSV(conn *sql.DB) error {
	absPath, err := filepath.Abs("data.csv")
	_, err = conn.Exec(`COPY users(name, job_titles, department, full_or_part_time, salary_or_hourly, typical_hours, annual_salary, hourly_rate) 
FROM '` + absPath + `' WITH CSV HEADER DELIMITER ',';`)
	if err != nil {
		return err
	}
	return nil
}

func StoreIndexByUser(conn *sql.DB) error {
	_, err := conn.Exec(`CREATE INDEX IF NOT EXISTS name_lower_idx ON users ((lower(name)))`)
	if err != nil {
		return err
	}
	return nil
}
