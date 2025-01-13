package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AddressDetails struct {
	Street string
	City   string
	State  string
	Zip    string
}

type User struct {
	ID   uint
	Name string
	AddressDetails
}

func main() {
	src := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "RajeevRanjitha@123", "localhost", "3306", "ranju")
	db, err := gorm.Open(mysql.Open(src), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return
	}

	db.AutoMigrate(&User{})

	user := User{
		Name: "John Doe",
		AddressDetails: AddressDetails{
			Street: "123 Elm Street",
			City:   "Springfield",
			State:  "IL",
			Zip:    "62704",
		},
	}
	db.Create(&user)
	var retrievedUser User
	db.First(&retrievedUser, user.ID)
	fmt.Printf("Name: %s, Address: %s, %s, %s, %s\n",
		retrievedUser.Name,
		retrievedUser.Street,
		retrievedUser.City,
		retrievedUser.State,
		retrievedUser.Zip,
	)
}
