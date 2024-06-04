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
		"database-pi7.c9m0oesie62n.sa-east-1.rds.amazonaws.com",
		"admin",
		"devfelipe",
		"motor_monitoring",
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

// GetDataCollections retorna os registros com base nos filtros
func GetDataCollections(p *Provider, filters map[string]string) ([]domain.DataCollection, error) {
	var dataCollections []domain.DataCollection
	query := p.DB

	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	if err := query.Find(&dataCollections).Error; err != nil {
		return nil, err
	}

	return dataCollections, nil
}

func createDBConnString(port int, host, username, password, name string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", username, password, host, port, name)
}
