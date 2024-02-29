package main

import (
	_ "embed"
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"dimoklan/internal/basmigration"
)

var (
	action   = flag.String("action", "", "type of migration: up/down")
	region   = flag.String("region", "", "aws region")
	endpoint = flag.String("endpoint", "", "dynamodb endpoint")
)

func main() {
	flag.Parse()

	if *action == "" {
		log.Fatal("action is required")
	}

	if *region == "" {
		log.Fatal("dsh is required")
	}

	if *endpoint == "" {
		log.Fatal("endpoint is required")
	}

	migrationActor := basmigration.New(*region, *endpoint)

	switch *action {
	case "up":
		migrationActor.CreateRegisterTable()
		migrationActor.CreateUserTable()
		migrationActor.AddUser()
		migrationActor.CreateDataTable()
	case "down":
		migrationActor.DeleteRegisterTable()
		migrationActor.DeleteUserTable()
		migrationActor.DeleteDataTable()
	default:
		log.Fatal("action is invalid")
	}

}
