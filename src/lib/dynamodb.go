package lib

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func NewDynamoClient() *dynamodb.DynamoDB {
	// Create AWS Session
	sess, err := AWSSessions()
	if err != nil {
		log.Println(err)
		return nil
	}
	dynamodbsvc := dynamodb.New(sess)
	// Return DynamoDB client
	return dynamodbsvc
}

func DynamoDBPutItem(dynamodbsvc *dynamodb.DynamoDB, tableName string, item interface{}) string {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Println("DynamoDB: Got error marshalling new movie item:")
		log.Println(err.Error())
		return "ERROR"
	}

	input := &dynamodb.PutItemInput{
		Item:                av,
		TableName:           aws.String(tableName),
		ConditionExpression: aws.String("attribute_not_exists(PK)"),
	}

	_, err = dynamodbsvc.PutItem(input)
	if err != nil {
		log.Println("DynamoDB: Got error calling PutItem:")
		log.Println(err.Error())

		if strings.HasPrefix(err.Error(), "ConditionalCheckFailedException") {
			return "CONDITION_FAILED"
		}
		return "ERROR"
	}

	return "OK"
}
