package migration

/*
func (m Migration) AddUser() {
	user := types.User{
		ID:            "u#3224053",
		SK:            "u#3224053",
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
		EntityType:    "user",
	}

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Fatalf("error in marshmap user; err: %v", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(consts.TableData),
	}

	if _, err = m.svc.PutItem(input); err != nil {
		log.Fatalf("error in creating user; err: %v", err)
	}

	fmt.Println("User added successfully")
}
*/
