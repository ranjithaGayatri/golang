package main

// import (
// 	"database/sql/driver"
// 	"encoding/json"
// 	"fmt"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type JSONData map[string]interface{}

// func (j *JSONData) Scan(value interface{}) error {
// 	if value == nil {
// 		*j = nil
// 		return nil
// 	}
// 	switch v := value.(type) {
// 	case string:
// 		return json.Unmarshal([]byte(v), j)
// 	case []byte:
// 		return json.Unmarshal(v, j)
// 	default:
// 		return fmt.Errorf("invalid type for JSONData: %T", v)
// 	}
// }
// func (j JSONData) Value() (driver.Value, error) {
// 	if j == nil {
// 		return nil, nil
// 	}
// 	return json.Marshal(j)
// }

// type User struct {
// 	ID       uint
// 	Name     string
// 	Metadata JSONData `gorm:"type:string"`
// }

// func main() {
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "RajeevRanjitha@123", "localhost", "3306", "ranju")
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("failed to connect to database:", err)
// 		return
// 	}
// 	fmt.Print("Hi Ranjtiha Gayatri")
// 	if err := db.AutoMigrate(&User{}); err != nil {
// 		fmt.Print("Error raised")
// 	}
// 	fmt.Print("Hi Ranjtiha Gayatri")

// 	user := User{
// 		Name: "John Doe",
// 		Metadata: JSONData{
// 			"age":  30,
// 			"city": "New York",
// 		},
// 	}

// 	db.Create(&user)

// 	var retrievedUser User
// 	db.First(&retrievedUser, user.ID)

// 	fmt.Printf("Name: %s, Metadata: %v\n", retrievedUser.Name, retrievedUser.Metadata)
// }
