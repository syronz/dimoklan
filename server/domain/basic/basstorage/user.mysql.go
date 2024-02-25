package basstorage

import (
	"database/sql"
	_ "embed"
	"fmt"

	"dimoklan/types"

	_ "github.com/go-sql-driver/mysql"
)

//go:embed queries/create_user.sql
var queryCreateUser string

func (bs *BasMysql) CreateUser(user types.User) error {
	_, err := bs.core.BasicMasterDB().Exec(
		queryCreateUser,
		user.Code,
		user.Name,
		user.Email,
		user.Password,
		user.Color,
		user.Language,
		user.Status,
		user.Reason,
	)
	if err != nil {
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

func (bs *BasMysql) CreateRegister(register types.Register) error {
	return nil
}

//go:embed queries/get_all_user_colors.sql
var queryGetAllColors string

func (bs *BasMysql) GetAllColors() (map[int]string, error) {
	rows, err := bs.core.BasicSlaveDB().Query(queryGetAllColors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mapColors := make(map[int]string)
	for rows.Next() {
		var userID int
		var color string
		if err := rows.Scan(&userID, &color); err != nil {
			return nil, fmt.Errorf("error in scanning row for get map_colors; %w", err)
		}

		mapColors[userID] = color
	}

	return mapColors, nil
}
