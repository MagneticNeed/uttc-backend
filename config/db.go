package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(user, password, database, host string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", user, password, host, database)
	var db *sql.DB
	var err error

	// リトライロジックを実装
	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Could not open db: %v", err)
		} else if err = db.Ping(); err == nil {
			break
		}

		log.Printf("Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return fmt.Errorf("could not establish database connection: %w", err)
	}

	DB = db
	log.Println("Database connection established")
	return nil
}

func GetDB() *sql.DB {
	return DB
}
