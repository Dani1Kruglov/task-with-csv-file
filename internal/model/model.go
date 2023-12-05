package model

type JobTitle struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"type:varchar(255)"`
}

type Department struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"type:varchar(255)"`
}

type WorkingDay struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"type:varchar(1);unique"`
}

type Payment struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"type:varchar(10);unique"`
}

type Worker struct {
	ID               uint   `gorm:"primaryKey"`
	Name             string `gorm:"type:varchar(100)"`
	JobTitlesID      uint
	DepartmentID     uint
	FullOrPartTimeID uint
	SalaryOrHourlyID uint
	JobTitle         JobTitle   `gorm:"foreignKey:JobTitlesID"`
	Department       Department `gorm:"foreignKey:DepartmentID"`
	FullOrPartTime   WorkingDay `gorm:"foreignKey:FullOrPartTimeID"`
	SalaryOrHourly   Payment    `gorm:"foreignKey:SalaryOrHourlyID"`
}

type WorkerHourlyPayment struct {
	ID           uint `gorm:"primaryKey"`
	WorkerID     uint
	TypicalHours int
	HourlyRate   float64 `gorm:"type:double precision"`
	Worker       Worker  `gorm:"foreignKey:WorkerID"`
}

type WorkerSalaryPayment struct {
	ID           uint `gorm:"primaryKey"`
	WorkerID     uint
	AnnualSalary float64 `gorm:"type:double precision"`
	Worker       Worker  `gorm:"foreignKey:WorkerID"`
}

type FullWorker struct {
	ID                  uint
	Name                string
	JobTitlesID         uint
	DepartmentID        uint
	FullOrPartTimeID    uint
	SalaryOrHourlyID    uint
	WorkerHourlyPayment WorkerHourlyPayment `gorm:"foreignKey:WorkerID"`
	WorkerSalaryPayment WorkerSalaryPayment `gorm:"foreignKey:WorkerID"`
	JobTitle            JobTitle            `gorm:"foreignKey:JobTitlesID"`
	Department          Department          `gorm:"foreignKey:DepartmentID"`
	FullOrPartTime      WorkingDay          `gorm:"foreignKey:FullOrPartTimeID"`
	SalaryOrHourly      Payment             `gorm:"foreignKey:SalaryOrHourlyID"`
}
