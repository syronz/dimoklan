package main

import (
	"dimoklan/internal/migration"
	_ "embed"
	"flag"
	"log"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	action = flag.String("action", "", "type of migration: up/down")
	dsn    = flag.String("dsn", "", "database dsn")
	steps  = flag.String("steps", "", "number accepted separated by comma")
)

func main() {
	flag.Parse()

	if *action == "" {
		log.Fatal("action is required")
	}

	if *dsn == "" {
		log.Fatal("dsh is required")
	}

	if *steps == "" {
		log.Fatal("steps is required")
	}

	var stepNums []int
	stepsArr := strings.Split(*steps, ",")
	for _, step := range stepsArr {
		stepNum, err := strconv.Atoi(step)
		if err != nil {
			log.Fatal("step needs be number")
		}
		stepNums = append(stepNums, stepNum)
	}

	log.Println("start migration for steps: ", stepNums)
	migration.MigrateDB(*dsn, *action, stepNums...)
}
