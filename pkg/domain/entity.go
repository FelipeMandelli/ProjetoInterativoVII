package domain

import (
	"time"

	"gorm.io/gorm"
)

type DataCollection struct {
	gorm.Model
	MotorID     string    `gorm:"primary_key;column:motor_identification"`
	Temperature float32   `gorm:"primary_key;column:temperature"`
	Sound       float32   `gorm:"primary_key;column:sound"`
	Current     float32   `gorm:"primary_key;column:current"`
	Vibration   string    `gorm:"primary_key;column:vibration"`
	DateTime    time.Time `gorm:"primary_key;column:datetime"`
}
