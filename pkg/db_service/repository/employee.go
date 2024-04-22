package repository

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ring0-rootkit/golang-staff-mngmnt/pkg/common/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type WorkEvent = string

const (
	WorkEventBegin WorkEvent = "B"
	WorkEventEnd   WorkEvent = "E"
)

type Employee struct {
	Id      int64
	Name    string
	Surname string
}

type Attendance struct {
	Id            int64
	EmployeeId    int64
	Time          time.Time
	WorkEventType WorkEvent
}

type Salary struct {
	SpH float64 `gorm:"column:salary_per_hour"`
}

var db *gorm.DB
var Log *log.Logger = logging.GetFor("db_service repository")

// TODO: rewrite using godotenv
func init() {
	dsn := "host=localhost user=postgres password=postgres " +
		"dbname=staff_mngmnt port=5432 sslmode=disable TimeZone=Europe/Minsk"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db = d
	if err != nil {
		panic(fmt.Sprintf("cant connect to database, err:%v", err.Error()))
	}
}

func GetHoursWorked(id int64) float64 {
	var attendance []Attendance = make([]Attendance, 0)
	db.Table("attendance").Where("employee_id = ?", id).Find(&attendance).Order("time")

	err := validateEmplAtt(attendance)
	if err != nil {
		Log.Printf("ERROR: employee attendance is not valid, "+
			"cannot get worked hours! id:%d; err:%s\n",
			id, err.Error())
		return -1
	}

	var hoursWorked float64 = 0
	for i := 0; i < len(attendance)-1; i += 2 {
		hoursWorked += attendance[i+1].Time.Sub(attendance[i].Time).Hours()
	}
	return hoursWorked
}

func GetSalaryPerHour(id int64) float64 {
	var s Salary
	db.Table("salary").Select("salary_per_hour").Where("employee_id = ?", id).First(&s)

	return s.SpH
}

func EmitWorkEvent(id int64, event WorkEvent) {
	a := Attendance{EmployeeId: id, Time: time.Now(), WorkEventType: event}
	db.Save(a)
}

func EmployeeIdByName(name string, surname string) int64 {
	var e Employee
	db.Table("employee").Where("name = ? AND surname = ?",
		strings.ToLower(name),
		strings.ToLower(surname)).First(&e)
	return e.Id
}

func validateEmplAtt(attendance []Attendance) error {
	if attendance[0].WorkEventType == WorkEventEnd {
		return errors.New("first work event is end of work")
	}
	for i := range len(attendance) - 1 {
		if attendance[i].WorkEventType == attendance[i+1].WorkEventType {
			return fmt.Errorf("two same events together\nevents:%v", attendance)
		}
	}
	return nil
}
