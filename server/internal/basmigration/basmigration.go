package basmigration

import (
	"database/sql"
	_ "embed"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//go:embed 001_users.up.sql
var usersUp string

//go:embed 002_cells.up.sql
var cellsUp string

var upSteps = map[int]string{
	1: usersUp,
	2: cellsUp,
}

var downSteps = map[int]string{
}

func MigrateDB(dsn string, action string, steps ...int) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error in opeining mysql connection; %v\n", err)
	}
	defer db.Close()

	switch action {
	case "up":
		for _, step := range steps {
			_, err = db.Exec(upSteps[step])
			if err != nil {
				log.Fatalf("error in up migration; %v\n", err)
			}
		}
	case "down":
		for step := range steps {
			_, err = db.Exec(downSteps[step])
			if err != nil {
				log.Fatalf("error in down migration; %v\n", err)
			}
		}
	default:
		log.Fatalln("action should be up or down")
	}
}
