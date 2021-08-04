package managers

import (
"fmt"
	"github.com/HugoJBello/go-heartbeat-service/models"
	"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func init(){
	fmt.Println("Connecting to database")
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

	database.AutoMigrate(&models.ServiceActivity{})

	DB = database
}
