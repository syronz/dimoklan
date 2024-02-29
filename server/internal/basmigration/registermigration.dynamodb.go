package basmigration

/*
func (m Migration) CreateRegisterTable() {
	// Check if the table exists
	describeTableInput := &dynamodb.DescribeTableInput{
		TableName: aws.String(consts.TableRegister),
	}
	_, err := m.svc.DescribeTable(describeTableInput)
	if err == nil {
		log.Fatalf("table already exist: %v, err:%v", consts.TableRegister, err)
	}

	fmt.Println("Table doesn't exist. Creating table...", consts.TableRegister)
	createTableInput := &dynamodb.CreateTableInput{
		TableName: aws.String(consts.TableRegister),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("email"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},
			{
				AttributeName: aws.String("activation_code"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("email"),
				KeyType:       aws.String(dynamodb.KeyTypeHash),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("activation_code_index"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("activation_code"),
						KeyType:       aws.String(dynamodb.KeyTypeHash),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String(dynamodb.ProjectionTypeAll),
				},
			},
		},
	}

	_, err = m.svc.CreateTable(createTableInput)
	if err != nil {
		log.Fatalf("Error creating table: %v; %v", consts.TableRegister, err)
	}

	fmt.Println("Table created successfully:", consts.TableRegister)

	// Wait for table creation to complete (Optional)
	// Note: This step is optional and depends on your use case
	err = m.svc.WaitUntilTableExists(&dynamodb.DescribeTableInput{
		TableName: aws.String(consts.TableRegister),
	})
	if err != nil {
		log.Fatalf("Error waiting for table: %v; %v", consts.TableRegister, err)
	}

	// Enable TTL on the table
	ttlSpecification := &dynamodb.TimeToLiveSpecification{
		AttributeName: aws.String("ttl"),
		Enabled:       aws.Bool(true),
	}

	updateInput := &dynamodb.UpdateTimeToLiveInput{
		TableName:               aws.String(consts.TableRegister),
		TimeToLiveSpecification: ttlSpecification,
	}

	_, err = m.svc.UpdateTimeToLive(updateInput)
	if err != nil {
		log.Fatalf("Error updating ttl in table: %v; %v", consts.TableRegister, err)
	}

	fmt.Println("TTL enabled successfully")

}

func (m Migration) DeleteRegisterTable() {

	input := &dynamodb.DeleteTableInput{
		TableName: aws.String(consts.TableRegister),
	}

	// Delete the table
	_, err := m.svc.DeleteTable(input)
	if err != nil {
		log.Printf("Error deleting table: %v; %v", consts.TableRegister, err)
		return
	}
}
*/
