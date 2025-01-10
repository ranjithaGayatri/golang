package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Db *gorm.DB

func ConnectingToDatabase() {
	if err := godotenv.Load(); err != nil {
		fmt.Print("Error Loading EnviromentVariables")
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	fmt.Print("Hi Ranjtha Gayatri ", user, password, host, port, dbname, " Hello")

	src := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(src), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Open the Database")
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Error Converting Struct into the Tables")
	}
	Db = db
}
