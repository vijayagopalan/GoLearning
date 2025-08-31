package database

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
	var err error
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "Vijay123"
		database = "ECommerce"
	)
	connString := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, password, database)
	// Open connection
	if DB != nil {
		return DB, nil
	}
	DB, err := sql.Open("postgres", connString)
	if err != nil {
		return DB, err
	}

	// Verify DB connection
	err = DB.Ping()
	if err != nil {
		return DB, err
	}

	fmt.Println("✅ Connected to DB!")
	return DB, err
}

// func Close() error {
// 	if Db != nil {
// 		return Db.Close()
// 	}
// 	return nil
// }
