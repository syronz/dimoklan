package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:root@tcp(127.001:3306)/dimo_basic"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error in opeining mysql connection; %v\n", err)
	}
	defer db.Close()

	// Create users table if not exists
	createTable := `
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    code VARCHAR(100) NOT NULL UNIQUE,
	color CHAR(6) NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
	language CHAR(2) NOT NULL DEFAULT 'en',
	status ENUM('active', 'inactive') DEFAULT 'active',
	reason VARCHAR(200) NOT NULL DEFAULT '',
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
	INDEX idx_code (code)
);
`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatalf("error in create user table; %v\n", err)
	}
}
