package repo

import (
	"dimoklan/model"
)

func (r *Repo) CreateAuth(auth model.Auth) error {
	/*
		auth.Email = consts.ParAuth + auth.Email
		auth.SK = auth.Email
		auth.EntityType = consts.AuthEntity
		av, err := dynamodbattribute.MarshalMap(auth)
		if err != nil {
			return fmt.Errorf("error in marshmap auth; err: %w", err)
		}

		input := &dynamodb.PutItemInput{
			Item:                av,
			TableName:           aws.String(consts.TableData),
			ConditionExpression: aws.String("attribute_not_exists(PK) AND attribute_not_exists(SK)"),
		}

		if _, err = r.core.DynamoDB().PutItem(input); err != nil {
			return fmt.Errorf("put_item_failed_for_auth; err:%w", err)
		}

		return nil
	*/
	return nil
}

func (r *Repo) DeleteAuth(authID string) error {
	/*
		input := &dynamodb.DeleteItemInput{
			TableName: aws.String(consts.TableData),
			Key: map[string]*dynamodb.AttributeValue{
				"PK": {S: aws.String(authID)},
				"SK": {S: aws.String(authID)},
			},
		}

		if _, err := r.core.DynamoDB().DeleteItem(input); err != nil {
			return fmt.Errorf("delete_item_failed_for_auth; err:%w", err)
		}

	*/
	return nil
}

func (r *Repo) GetAuthByEmail(email string) (model.Auth, error) {
	/*
		email = consts.ParAuth + email

		params := &dynamodb.GetItemInput{
			TableName: aws.String(consts.TableData),
			Key: map[string]*dynamodb.AttributeValue{
				"PK": {
					S: aws.String(email),
				},
				"SK": {
					S: aws.String(email),
				},
			},
		}

		resp, err := r.core.DynamoDB().GetItem(params)
		if err != nil {
			return model.Auth{}, fmt.Errorf("error in getting auth entity; err: %w", err)
		}

		auth := model.Auth{}
		err = dynamodbattribute.UnmarshalMap(resp.Item, &auth)
		if err != nil {
			return model.Auth{}, fmt.Errorf("binding auth data failed; err: %w", err)
		}

		return auth, nil
	*/
	return model.Auth{}, nil
}
