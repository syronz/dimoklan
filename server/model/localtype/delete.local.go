package localtype

type Delete struct {
	PK string `dynamodbav:"PK"`
	SK string `dynamodbav:"SK"`
}
