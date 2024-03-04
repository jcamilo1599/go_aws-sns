package main

import (
  "fmt"
  "os"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/sns"
)

func main() {
  sess, err := session.NewSession(&aws.Config{
    Region:      aws.String("us-east-2"),
    Credentials: credentials.NewStaticCredentials("KEY", "SECRET-KEY", ""),
  })

  if err != nil {
    panic(err.Error())
  }

  svc := sns.New(sess)
  phone := ""

  result, err := svc.Publish(&sns.PublishInput{
    Message:     aws.String("Probando mensajes de texto..."),
    PhoneNumber: aws.String("+57" + phone),
  })

  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  fmt.Println(*result.MessageId)
}
