package utils

import (
	"auth/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //Gorm postgres dialect interface
	"github.com/joho/godotenv"
	"log"
	"os"
)


func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("databaseUser")
	password := os.Getenv("databasePassword")
	databaseName := os.Getenv("databaseName")
	databaseHost := os.Getenv("databaseHost")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, username, databaseName, password)

	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}

	db.AutoMigrate(
		&models.UserReg{},
		&models.Favorite{},
		&models.History{})
	db.Model(&models.History{}).AddForeignKey("user_id", "user_regs(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Favorite{}).AddForeignKey("user_id", "user_regs(id)", "RESTRICT", "RESTRICT")

	fmt.Println("Successfully connected!", db)
	return db
}
