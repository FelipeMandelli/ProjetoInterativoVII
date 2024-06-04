package service

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pi.go/pkg/domain"
)

func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(createDBConnString(
		3306,
		"database-pi7.c9m0oesie62n.sa-east-1.rds.amazonaws.com",
		"admin",
		"devfelipe",
		"motor_monitoring",
	)), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not create connection: %w", err)
	}

	err = db.AutoMigrate(
		&domain.DataCollection{},
	)
	if err != nil {
		return nil, fmt.Errorf("could not execute auto migrate: %w", err)
	}

	return db, nil
}

func createDBConnString(port int, host, username, password, name string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", username, password, host, port, name)
}
