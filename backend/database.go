package backend

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	m "trademarkia.com/animeAPI/model"
)

//DB variables declared
var DB *gorm.DB
var err error

//DNS points to local MySQL database connection
const DNS = "root:pass@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

//InitialMigration function to establish connection
func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	DB.AutoMigrate(&m.Anime{})
}

//GetAnime function to return Anime model from database
func GetAnime(id int) m.Anime {
	var anime m.Anime
	DB.Find(&anime, id)
	return anime
}

//SaveAnime function to save Anime to database
func SaveAnime(id int) {
	anime := GetByID(id)
	DB.Create(&anime)
}
