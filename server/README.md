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

## Get the aws data

```bash
# get the all data in local dynamodb
aws dynamodb scan --table-name data --endpoint-url http://localhost:8000

# get a value by PK and SK
aws dynamodb get-item --table-name data \
--key '{"PK": {"S": "f#1:1"}, "SK": {"S": "c#2:6"}}' \
--endpoint-url http://127.0.0.1:8000 \
--output text

# query a value just by PK
aws dynamodb query --table-name data \
--key-condition-expression "PK = :pk" \
--expression-attribute-values '{":pk": {"S": "f#1:1"}}' \
--endpoint-url http://localhost:8000 \
--output text

# get a marshal: PK=userID SK=marshalID
aws dynamodb get-item --table-name data \
--key '{"PK": {"S": "u#3224053"}, "SK": {"S": "m#3224053:1"}}' \
--endpoint-url http://127.0.0.1:8000 \
--output text


```
