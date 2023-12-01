package model

import "database/sql"

type User struct {
	ID             int64           `json:"id"`
	Name           string          `json:"name"`
	JobTitles      string          `json:"job_titles"`
	Department     string          `json:"department"`
	FullOrPartTime string          `json:"full_or_part_time"`
	SalaryOrHourly string          `json:"salary_or_hourly"`
	TypicalHours   sql.NullInt64   `json:"typical_hours"`
	AnnualSalary   sql.NullFloat64 `json:"annual_salary"`
	HourlyRate     sql.NullFloat64 `json:"hourly_rate"`
}
