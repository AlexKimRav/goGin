package configs

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Connection without gorm with raw queries for test(No need)
func connectToMySql() {
	//Getenv is taking from system var with name MYSQL_PASSWORD
	fmt.Println("\nPASSWORD IS: " + os.Getenv("MYSQL_PASSWORD") + "\n")
	pw := os.Getenv("MYSQL_PASSWORD")

	db, err := sql.Open("mysql", "root:"+pw+"@/goschema")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
