package connection

import (
	"todo-app/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

var (
	// DB variable for usage in controllers/models
	DB               *gorm.DB
	connectionString = "root:root@tcp(localhost:3306)/todoapp?charset=utf8&parseTime=true"
)

//InitConnection setup DB var
func InitConnection() {
	con, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	// check if connection was really opened
	er := con.DB().Ping()
	if er != nil {
		panic(er)
	}

	//con.DropTableIfExists(&models.Task{}, &models.User{})
	con.AutoMigrate(&models.Task{}, &models.User{})

	DB = con

	migrateUser()
}

func migrateUser() {
	user := models.User{
		Name:     "Test User",
		Username: "admin",
		Password: "admin"}

	if res := DB.Create(&user); res.Error != nil {
		panic(res.Error)
	}
}
