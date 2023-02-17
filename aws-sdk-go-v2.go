package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func main() {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("us-east-2"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			"CLAVE DE ACCESO",
			"CLAVE DE ACCESO SECRETA",
			"",
		)),
	)

	if err != nil {
		panic(err.Error())
	}

	client := sns.NewFromConfig(cfg)
	message := "Probando mensajes de texto..."
	phone := "+57"

	input := &sns.PublishInput{
		Message:     &message,
		PhoneNumber: &phone,
	}

	result, err := client.Publish(context.TODO(), input)

	if err != nil {
		panic(err.Error())
		return
	}

	fmt.Println(*result.MessageId)
}
