package initialise

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"github.com/SaaranshShandilya/go-crud/models"
)

func ConnectToDb() *gorm.DB{
	dsn := "host=satao.db.elephantsql.com user=bykxchwg password=bTKf5wrtfQvAsFhNPzVSrJEXk9bFD0n7 dbname=bykxchwg port=5432 sslmode=disable"
  	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err !=nil{
		log.Fatal("Not able to connect to db", err)
	}

	DB.AutoMigrate(&models.Post{})

	return DB
}