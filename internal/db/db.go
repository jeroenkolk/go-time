package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func SetupDB() {
	log.Printf("Setting up DB")
	db, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err2 := db.AutoMigrate(&Product{})
	if err2 != nil {
		log.Panicln("Failed to migrate", err)
	}

	var p []Product

	db.Find(&p)

	//userRepo := UserRepository(db)
}
