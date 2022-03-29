package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/TransportMall/payment-systems/internal/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //For postgres driver
)

var dbInstance *gorm.DB

//InitDB Initialize Database connection
func InitDB() *gorm.DB {

	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("DBPASSWORD"))
	dbInstance, err := gorm.Open("postgres", dbConnectionString)

	if err != nil {
		panic(fmt.Sprintf("Failed to connect database for %s, %s", dbConnectionString, err))
	} else {
		fmt.Println("Successfully connect to database")
	}
	//defer db.Close()

	// Run Migrations
	errors := dbInstance.AutoMigrate(&model.Admin{}, &model.Configurations{}, &model.Localization{}, &model.Users{}, &model.UsersAddress{}).Error
	if errors != nil {
		log.Println(errors.Error())
	}
	return dbInstance
}
