package infrastructure

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	//Tên biến và kiểu của biến
	DB *gorm.DB
}

//Hàm NewDatabase trả về stuct Database tại bên trên
func NewDatabase() Database {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)
	fmt.Printf("URL connection : %s", URL)
	db, err := gorm.Open(mysql.Open(URL))
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Database connection established!")
	return Database{
		DB: db,
	}
}
