package provider

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm(ctx context.Context) *gorm.DB {
	dbHost := ""
	dbUser := ""
	dbPass := ""
	dbName := ""
	dbPort := 5432
	sslMode := ""

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", dbHost, dbUser, dbPass, dbName, dbPort, sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("[gorm.Open]: %v", err)
	}

	return db
}
