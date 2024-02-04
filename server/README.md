# dimoklan

Server side of game


## run

```bash
go run main.go -cfg=./config.yaml
```

## database migration

```bash
go run script/migration/0001_create_user_table.up.go
```

## integration test

```bash
go test ./integration/mapgenerator/ -v

# run a specific test
go test -run TestA ./integration/mapgenerator/ -v
```
