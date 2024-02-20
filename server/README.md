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
```

## integration test

```bash
go test ./integration/mapgenerator/ -v

# run a specific test
go test -run TestA ./integration/mapgenerator/ -v
```

