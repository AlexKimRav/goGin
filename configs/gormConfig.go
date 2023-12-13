package configs

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	//Getenv is taking from system var with name MYSQL_PASSWORD
	fmt.Println("\nPASSWORD IS: " + os.Getenv("MYSQL_PASSWORD") + "\n")
	pw := os.Getenv("MYSQL_PASSWORD")

	dsn := "root:" + pw + "@/goschema?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func GetDb() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		sleep := time.Duration(1)
		for dbase == nil {
			sleep = sleep * 2
			fmt.Println("db is unavailable, wait for %d sec.\n", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}
	}
	return dbase
}
