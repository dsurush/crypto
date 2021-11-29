package database

import (
	"cryptotest/settings"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetDBConnection() *sql.DB {
	//dsn := "dsurush:dsurush@tcp(127.0.0.1:3306)/crypto?charset=utf8mb4&parseTime=True&loc=Local"
	//readSettings := settings.ReadSettings(`C:\Users\User\go\src\cryptotest\settings\settings-dev.json`)
	readSettings := settings.ReadSettings(`.\settings\settings-dev.json`)
	dsn := fmt.Sprintf("%s:%s@/%s", readSettings.Username, readSettings.Password, readSettings.NameDB)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Can't open connection check err %err\n", err)
	}

	log.Println("DB connection is success")

	/*
		readSettings := settings.ReadSettings(`.\settings\settings-dev.json`)
		time.Sleep(5 *time.Second)
		dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", readSettings.Username, readSettings.Password,
			readSettings.Port, readSettings.NameDB)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Can't open connection check err %err\n", err)
		}
		log.Println("DB connection is success")
	*/
	return db
}
