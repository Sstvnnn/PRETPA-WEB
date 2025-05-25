package database

import (
	"github.co/Sstvnnn/PRETPA-WEB/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dbURL := "postgres://postgres:admin@localhost:6969/PreTPAWEB"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil{
		panic(err)
	}

	err = db.AutoMigrate(
		&model.User{},
	)
	
	if err != nil{
		panic(err)
	}

	return db
}
