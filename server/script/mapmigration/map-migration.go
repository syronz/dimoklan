package main

import (
	_ "embed"
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"dimoklan/internal/mapmigration"
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

	migrationActor := mapmigration.New(*region, *endpoint)

	switch *action {
	case "up":
		migrationActor.CreateMapTable()
	case "down":
		migrationActor.DeleteMapTable()
	default:
		log.Fatal("action is invalid")
	}

}
