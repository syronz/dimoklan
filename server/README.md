# dimoklan

Server side of game


## run

```bash
go run cmd/fullserver/main.go -cfg=./config.yaml
```

## database migration

```bash
go run script/dbmigration/dbmigration.go -action=down -region=us-west-2 -endpoint=http://127.0.0.1:8000
go run script/dbmigration/dbmigration.go -action=up -region=us-west-2 -endpoint=http://127.0.0.1:8000



```

## integration test

```bash
ginkgo ./...

# code coverage
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out



go test ./integration/mapgenerator/ -v

# run a specific test
go test -run TestA ./integration/mapgenerator/ -v
```

## run dynamodb

```bash
~/projects/docker-composes/dynamodb/01-daynamodb$ docker-compose up
```

## NoSQL Workbench

use NoSQL Workbench to see the dynamodb values
