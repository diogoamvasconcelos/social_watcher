// Reference: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-dynamodb-with-go-sdk.html

package lib

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
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

func DynamoDBPutItem(dynamodbsvc *dynamodb.DynamoDB, tableName string, item interface{}, conditionExpression string) string {
	dbItem, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Println("DynamoDB: Got error marshalling item:")
		log.Println(err.Error())
		return "ERROR"
	}

	input := dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      dbItem,
	}
	if conditionExpression != "" {
		input.ConditionExpression = aws.String(conditionExpression)
	}

	_, err = dynamodbsvc.PutItem(&input)
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

func DynamoDBGetItem(dynamodbsvc *dynamodb.DynamoDB, tableName string, key interface{}) (map[string]*dynamodb.AttributeValue, error) {
	dbKey, err := dynamodbattribute.MarshalMap(key)
	if err != nil {
		log.Println("DynamoDB: Got error marshalling key:")
		log.Println(err.Error())
		return nil, err
	}

	input := dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       dbKey,
	}

	result, err := dynamodbsvc.GetItem(&input)
	return result.Item, nil
}

// UnmarshalStreamImage converts events.DynamoDBAttributeValue to struct
// from: https://stackoverflow.com/questions/49129534/unmarshal-mapstringdynamodbattributevalue-into-a-struct
func UnmarshalStreamImage(attribute map[string]events.DynamoDBAttributeValue, out interface{}) error {
	dbAttrMap := make(map[string]*dynamodb.AttributeValue)
	for k, v := range attribute {
		var dbAttr dynamodb.AttributeValue
		bytes, marshalErr := v.MarshalJSON()
		if marshalErr != nil {
			return marshalErr
		}

		json.Unmarshal(bytes, &dbAttr)
		dbAttrMap[k] = &dbAttr
	}

	return dynamodbattribute.UnmarshalMap(dbAttrMap, out)
}
