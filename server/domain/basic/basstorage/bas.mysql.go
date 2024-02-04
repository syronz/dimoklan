package basstorage

import (
	"database/sql"
	"dimoklan/internal/config"
	"dimoklan/types"
	_ "embed"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BasMysql struct {
	core config.Core
}

func New(cfg config.Core) *BasMysql {
	return &BasMysql{
		core: cfg,
	}
}

//go:embed queries/create_user.sql
var queryCreateUser string

func (bs *BasMysql) CreateUser(user types.User) error {
	_, err := bs.core.BasicMasterDB().Exec(
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

//go:embed queries/get_user_by_color.sql
var queryGetUserByColor string

func (bs *BasMysql) GetUserByColor(color string) (types.User, error) {
	var user types.User

	err := bs.core.BasicSlaveDB().QueryRow(queryGetUserByColor, color).Scan(
		&user.ID,
		&user.Code,
		&user.Color,
		&user.Status,
	)

	if err == sql.ErrNoRows {
		err = nil
	}

	if err != nil {
		return user, fmt.Errorf("error in getting user by color; %w", err)
	}

	return user, nil
}
