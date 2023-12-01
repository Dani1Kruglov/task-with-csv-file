package storage

import (
	"csv-file/internal/model"
	"database/sql"
	"encoding/json"
)

func GetUserByNameFromDB(userName string, conn *sql.DB) ([]byte, error) {
	var user model.User

	err := conn.QueryRow("SELECT * FROM users WHERE name = $1 LIMIT 1", userName).Scan(
		&user.ID,
		&user.Name,
		&user.JobTitles,
		&user.Department,
		&user.FullOrPartTime,
		&user.SalaryOrHourly,
		&user.TypicalHours,
		&user.AnnualSalary,
		&user.HourlyRate,
	)
	if err != nil {
		return nil, err
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	return userJSON, nil
}
