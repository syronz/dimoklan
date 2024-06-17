package repo

import (
	"context"

	"dimoklan/consts/entity"
	"dimoklan/consts/hashtag"
	"dimoklan/model"
)

func (r *Repo) CreateUser(ctx context.Context, user model.User) error {
	return r.putUniqueItem(ctx, entity.User, user.ToRepo())
}

func (r *Repo) DeleteUser(ctx context.Context, userID string) error {
	pk := hashtag.User + userID
	return r.deleteItem(ctx, entity.User, pk)

	// userIDMarshaled, err := attributevalue.Marshal(userID)
	// if err != nil {
	// 	return fmt.Errorf("error in marshal userID; err: %w", err)
	// }

	// params := &dynamodb.DeleteItemInput{
	// 	TableName: table.Data(),
	// 	Key: map[string]types.AttributeValue{
	// 		"PK": userIDMarshaled,
	// 		"SK": userIDMarshaled,
	// 	},
	// }

	// if _, err := r.core.DynamoDB().DeleteItem(ctx, params); err != nil {
	// 	return fmt.Errorf("delete item failed for user; err:%w", err)
	// }

	// return nil
}

// func (r *Repo) GetUserByEmail(ctx context.Context, email string) (model.UserRepo, error) {

// 	emailMarshaled, err := attributevalue.Marshal(email)
// 	if err != nil {
// 		return model.UserRepo{}, fmt.Errorf("error in marshal email; err: %w", err)
// 	}

// 	params := &dynamodb.GetItemInput{
// 		TableName: table.Data(),
// 		Key: map[string]types.AttributeValue{
// 			"PK": emailMarshaled,
// 			"SK": emailMarshaled,
// 		},
// 	}
// 	resp, err := r.core.DynamoDB().GetItem(ctx, params)
// 	if err != nil {
// 		return model.UserRepo{}, fmt.Errorf("error in getting user entity; err: %w", err)
// 	}

// 	userRepo := model.UserRepo{}
// 	err = attributevalue.UnmarshalMap(resp.Item, &userRepo)
// 	if err != nil {
// 		return model.UserRepo{}, fmt.Errorf("binding userRepo data failed; err: %w", err)
// 	}

// 	return userRepo, nil

// }
