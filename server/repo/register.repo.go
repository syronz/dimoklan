package repo

import (
	"context"

	"dimoklan/consts/entity"
	"dimoklan/consts/hashtag"
	"dimoklan/model"
)

func (r *Repo) CreateRegister(ctx context.Context, register model.Register) error {

	return r.putUniqueItem(ctx, entity.Register, register.ToRepo())

	// item, err := attributevalue.MarshalMap(registerRepo)
	// if err != nil {
	// 	return fmt.Errorf("error in marshmap registerRepo; %w", err)
	// }

	// itemInput := &dynamodb.PutItemInput{
	// 	TableName:           table.Data(),
	// 	Item:                item,
	// 	ConditionExpression: aws.String("attribute_not_exists(PK)"),
	// }

	// _, err = r.core.DynamoDB().PutItem(ctx, itemInput)
	// if err != nil {
	// 	var conditionalCheckFailedErr *types.ConditionalCheckFailedException
	// 	if errors.As(err, &conditionalCheckFailedErr) {
	// 		return fmt.Errorf("email already exists")
	// 	}

	// 	return fmt.Errorf("error in register; err: %w", err)
	// }
	// return nil
}

func (r *Repo) ConfirmRegister(ctx context.Context, activationCode string) (model.Register, error) {
	registerRepo := model.RegisterRepo{}
	if err := r.getItem(ctx, entity.Register, &registerRepo, hashtag.Register+activationCode); err != nil {
		return model.Register{}, err
	}

	return registerRepo.ToAPI(), nil

	// activationCodeMarshaled, err := attributevalue.Marshal(activationCode)
	// if err != nil {
	// 	return model.RegisterRepo{}, fmt.Errorf("error in marshal activation_code; err: %w", err)
	// }

	// params := &dynamodb.GetItemInput{
	// 	TableName: table.Data(),
	// 	Key: map[string]types.AttributeValue{
	// 		"PK": activationCodeMarshaled,
	// 		"SK": activationCodeMarshaled,
	// 	},
	// }

	// resp, err := r.core.DynamoDB().GetItem(context.TODO(), params)
	// if err != nil {
	// 	return model.RegisterRepo{}, fmt.Errorf("error in getting register entity; err: %w", err)
	// }

	// registerRepo := model.RegisterRepo{}
	// err = attributevalue.UnmarshalMap(resp.Item, &registerRepo)
	// if err != nil {
	// 	return model.RegisterRepo{}, fmt.Errorf("binding registration data failed; err: %w", err)
	// }

	// return registerRepo, nil

}
