package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var lock = &sync.Mutex{}

var sqlSingleton *sql.DB

func GetInstance() *sql.DB {
	if sqlSingleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if sqlSingleton == nil {
			fmt.Println("Creating single instance now.")
			mysqlUser := os.Getenv("MYSQL_USER")
			mysqlPassword := os.Getenv("MYSQL_PASSWORD")
			mysqlDB := os.Getenv("MYSQL_DATABASE")
			connectionString := fmt.Sprintf("%s:%s@tcp(db)/%s?parseTime=true", mysqlUser, mysqlPassword, mysqlDB)
			conn, err := sql.Open("mysql", connectionString)
			if err != nil {
				log.Fatal("MySQL connection fail")
			}
			sqlSingleton = conn
		}
	}

	return sqlSingleton
}
