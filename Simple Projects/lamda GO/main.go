package main

import (
	"fmt"

	models "github.com/arnavmahajan630/Learn-Go/Simple-Projects/lamda-GO/Models"
	"github.com/aws/aws-lambda-go/lambda"
)




func HandleLambdaEvent(event models.MyEvent) (models.Myresponse, error) {
	return models.Myresponse{Message: fmt.Sprintf("%s is %d years old", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}