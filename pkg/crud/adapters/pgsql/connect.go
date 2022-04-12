package pgsql

import (
	"fmt"
	"go-base/pkg/crud/adapters/pgsql/models"
	"go-base/pkg/helpers/configs"
	"go-base/pkg/helpers/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(config *configs.Config) *gorm.DB {
	cf := config.Postgresql
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", cf.Host,
		cf.Port, cf.User, cf.DbName, cf.SslMode, cf.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(err, "failed to connect database, db name:[%s]", cf.DbName)
	} else {
		log.Infof("connect to db :[%s]", cf.DbName)
	}
	if err = db.AutoMigrate(&models.Example{}); err != nil {
		log.Fatal(err, "auto migrate fail")
		return nil
	}
	return db
}
