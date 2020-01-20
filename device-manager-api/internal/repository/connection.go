package repository

import (
	"github.com/jclaudiocf/do-go/device-manager-api/internal/model"
	"github.com/jclaudiocf/do-go/device-manager-api/internal/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("postgres",
		"host=localhost port=5432 user=root password=admin dbname=device sslmode=disable")

	util.ErrorChecking(err, "connect database")

	return db
}

func InitMigrations() {
	db := GetConnection()
	db.AutoMigrate(&model.Device{})
	defer db.Close()
}
