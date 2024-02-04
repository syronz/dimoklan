package basstorage

import (
	"database/sql"
	"dimoklan/internal/config"
	"dimoklan/types"
	_ "embed"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BasMysql struct {
	cfg config.Core
}

func New(cfg config.Core) *BasMysql {
	return &BasMysql{
		cfg: cfg,
	}
}

func loadUserBit(db *sql.DB) int {
	getMaxBitQuery := `SELECT IFNULL(MAX(bit),1) FROM users;`
	var maxBit int

	err := db.QueryRow(getMaxBitQuery).Scan(&maxBit)
	if err != nil {
		log.Fatalf("error getting max bit: %v", err)
	}

	return maxBit
}

//go:embed queries/create_user.sql
var queryCreateUser string

func (bs *BasMysql) CreateUser(user types.User) error {
	_, err := bs.cfg.BasicMasterDB().Exec(
		queryCreateUser,
		user.Code,
		user.Name,
		user.Email,
		user.Username,
		user.Password,
		user.Color,
		user.Language,
		user.Status,
		user.Reason,
	)
	if err != nil {
		log.Printf("error in saving the user; err: %v", err)
		return err
	}

	return nil
}
