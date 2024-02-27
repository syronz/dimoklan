# dimoklan

Server side of game


## run

```bash
go run cmd/basserver/main.go -cfg=./config.yaml
go run cmd/mapserver/main.go -cfg=./config.yaml
```

## database migration

```bash
go run script/basmigration/main.go -action=up -dsn="root:root@tcp(127.001:3306)/dimo_basic" -steps=1


-- map domain
go run script/mapmigration/map-migration.go -action=down -region=us-west-2 -endpoint=http://127.0.0.1:8000
go run script/mapmigration/map-migration.go -action=up -region=us-west-2 -endpoint=http://127.0.0.1:8000 

-- bas domain
go run script/basmigration/basmigration-dynamodb.go -action=down -region=us-west-2 -endpoint=http://127.0.0.1:8000
go run script/basmigration/basmigration-dynamodb.go -action=up -region=us-west-2 -endpoint=http://127.0.0.1:8000



```

## integration test

```bash
go test ./integration/mapgenerator/ -v

# run a specific test
go test -run TestA ./integration/mapgenerator/ -v
```

