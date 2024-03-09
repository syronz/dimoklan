package migration

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"dimoklan/consts/entity"
	"dimoklan/consts/hashtag"
	"dimoklan/consts/table"
	"dimoklan/model"
)

func (m Migration) AddUser() {
	// Add user
	userRepo := model.UserRepo{
		PK:            hashtag.User + "3224053",
		SK:            hashtag.User + "3224053",
		Color:         "3131f5",
		Email:         "sabina.diako@gmail.com",
		Kingdom:       "Malanda",
		Password:      "6b53d67e399b703b38c58fa4c9e25438478ca0372b190abc2e34579e5e3cfa83",
		Language:      "en",
		Suspend:       false,
		SuspendReason: "",
		Freeze:        false,
		FreezeReason:  "",
		CreatedAt:     1709064739,
		UpdatedAt:     1709064739,
		EntityType:    entity.User,
	}

	m.putItem(userRepo)

	// Add auth
	authRepo := model.AuthRepo{
		PK:            hashtag.Auth + "sabina.diako@gmail.com",
		SK:            hashtag.Auth + "sabina.diako@gmail.com",
		Password:      "6b53d67e399b703b38c58fa4c9e25438478ca0372b190abc2e34579e5e3cfa83",
		Suspend:       false,
		SuspendReason: "",
		UserID:        "3224053",
		EntityType:    entity.Auth,
	}

	m.putItem(authRepo)

}

func (m Migration) putItem(itemRepo any) {
	item, err := attributevalue.MarshalMap(itemRepo)
	if err != nil {
		log.Fatalf("error in muserRepoarshmap item; %v", err)
	}

	itemInput := &dynamodb.PutItemInput{
		TableName:           table.Data(),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(PK) AND attribute_not_exists(SK)"),
	}

	_, err = m.client.PutItem(context.TODO(), itemInput)
	if err != nil {
		var conditionalCheckFailedErr *types.ConditionalCheckFailedException
		if errors.As(err, &conditionalCheckFailedErr) {
			log.Fatalln("item already exists")
		}

		log.Fatalf("error in putting item; err: %v", err)
	}
}
