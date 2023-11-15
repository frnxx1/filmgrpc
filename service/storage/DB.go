package storage

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	
	DatabaseConnection()
	
}


var DB *gorm.DB
var err error

type Film struct {
	ID        string `gorm:"primarykey"`
	Title     string
	Genre     string
	CreatedAt time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:false"`
}

func DatabaseConnection() {
	localhost := "localhost"
	
	db := "db"
	user := "user"
	pass := "pass"
	dsn := fmt.Sprintf("host=%s  user=%s dbname=%s password=%s sslmode=disable",
		localhost,
		
		user,
		db,
		pass ,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	DB.AutoMigrate(Film{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
}
