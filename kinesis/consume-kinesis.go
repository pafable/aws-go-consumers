package main

/*
	This lambda function will retrieve data from Kinesis Data Stream
*/

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/events"
	"github.com/aws/aws-sdk-go/lambda"
)

func main() {
	lambda.Start(kinesisHandler)
}

func kinesisHandler(ctx context.Context, kinesisEvent events.KinesisEvent) error {

	Record := kinesisEvent.Records[0]
	fmt.Println("Region: " + Record.AwsRegion)
	fmt.Println("Records:" + Record)

	return nil
}
