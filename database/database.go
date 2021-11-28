package database

import (
	"cryptotest/settings"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func GetDBConnection() *gorm.DB {
	//dsn := "dsurush:dsurush@tcp(127.0.0.1:3306)/crypto?charset=utf8mb4&parseTime=True&loc=Local"
	//readSettings := settings.ReadSettings(`C:\Users\User\go\src\cryptotest\settings\settings-dev.json`)
	readSettings := settings.ReadSettings(`.\settings\settings-dev.json`)
	time.Sleep(5 *time.Second)
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", readSettings.Username, readSettings.Password,
		readSettings.Port, readSettings.NameDB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Can't open connection check err %err\n", err)
	}
	log.Println("DB connection is success")
	return db
}
