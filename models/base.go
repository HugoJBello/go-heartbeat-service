// models/setup.go

package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func init(){
	ConnectDataBase()
}

func GetDb() *gorm.DB{
	return DB
}

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&ServiceActivity{})

	DB = database
}