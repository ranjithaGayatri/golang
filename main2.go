package main

// import (
// 	"database/sql/driver"
// 	"encoding/json"
// 	"errors"
// 	"fmt"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type Preferences struct {
// 	Theme  string `json:"theme"`
// 	Layout string `json:"layout"`
// }

// func (p *Preferences) Scan(value interface{}) error {
// 	bytes, ok := value.([]byte)
// 	if !ok {
// 		return errors.New("type assertion to []byte failed")
// 	}
// 	return json.Unmarshal(bytes, p)
// }

// func (p Preferences) Value() (driver.Value, error) {
// 	return json.Marshal(p)
// }

// type User struct {
// 	ID          uint
// 	Name        string
// 	Preferences Preferences `gorm:"type:string"`
// }

// func main() {
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "RajeevRanjitha@123", "localhost", "3306", "ranju")
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("failed to connect to the database:", err)
// 		return
// 	}

// 	db.AutoMigrate(&User{})

// 	user := User{
// 		Name: "Alice",
// 		Preferences: Preferences{
// 			Theme:  "dark",
// 			Layout: "grid",
// 		},
// 	}
// 	db.Create(&user)

// 	var retrievedUser User
// 	db.First(&retrievedUser, user.ID)

// 	fmt.Printf("Name: %s, Preferences: %+v\n", retrievedUser.Name, retrievedUser.Preferences)
// }
