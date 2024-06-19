package database

import (
	"database/sql"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(dataSourceName string) {
	var err error

	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := executeSQLFile("database/database.sql"); err != nil {
		log.Fatalf("Failed to execute database.sql: %v", err)
	}

	// データベース接続を確認
	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to ping databse: %v", err)
	}

	log.Println("Database connected successfully")
}

func executeSQLFile(filePath string) error {
	// sqlファイルを読み込む
	sqlBytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}

	queries := strings.Split(string(sqlBytes), ";")
	for _, query := range queries {
		if strings.TrimSpace(query) == "" {
			continue
		}
		_, err := DB.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}
