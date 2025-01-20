package database

import (
	"log"

	"github.com/juliasilvamoura/gin-api-rest/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectDatabase() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Error conect database")
	}
	DB.AutoMigrate(&model.Aluno{})
}
