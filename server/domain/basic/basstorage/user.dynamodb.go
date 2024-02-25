package basstorage

import "dimoklan/types"

func (bd *BasDynamoDB) GetUserByColor(color string) (types.User, error) {

	return types.User{}, nil
}

func (bd *BasDynamoDB) CreateUser(user types.User) (error) {

	return nil
}

func (bd *BasDynamoDB) GetAllColors() (map[int]string, error) {

	return nil, nil
}
