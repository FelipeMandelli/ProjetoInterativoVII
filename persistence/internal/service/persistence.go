package service

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pi.go/pkg/domain"
)

func ConnectDatabase(provider *Provider) error {
	db, err := gorm.Open(mysql.Open(createDBConnString(
		3306,
		"",
		"",
		"",
		"",
	)), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("could not create connection: %w", err)
	}

	err = db.AutoMigrate(
		&domain.DataCollection{},
	)
	if err != nil {
		return fmt.Errorf("could not execute auto migrate: %w", err)
	}

	provider.DB = db

	return nil
}

func PersistData(p *Provider, data *domain.DataCollection) error {
	err := p.DB.Save(data).Error
	if err != nil {
		return fmt.Errorf("could not perform insert: %w", err)
	}

	return nil
}

func createDBConnString(port int, host, username, password, name string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", username, password, host, port, name)
}
