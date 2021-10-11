package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var sess = session.Must(session.NewSession())
var s3svc = s3iface.S3API(s3.New(sess))
var uploader = s3manager.NewUploader(sess)

type MenuItem struct {
	DishName string  `json:"dishName"`
	Price    float32 `json:"price"`
}

type Restaurants struct {
	CashBalance    float32    `json:"cashBalance"`
	Menu           []MenuItem `json:"menu"`
	OpeningHours   string     `json:"openingHours"`
	RestaurantName string     `json:"restaurantName"`
}

func handler(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key

		log.Printf("Bucket: %s, key:%s\n", bucket, key)

		obj, err := s3svc.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})

		if err != nil {
			log.Printf("Error in downloading %s from S3: %s\n", key, err)
			continue
		}

		var res []Restaurants
		body, err := ioutil.ReadAll(obj.Body)
		if err != nil {
			log.Printf("Error in reading file %s: %s\n", key, err)
			continue
		}

		err = json.Unmarshal(body, &res)
		if err != nil {
			log.Printf("Error in unmarshalling request file %s: %s\n", key, err)
			continue
		}

		fmt.Println(res)
	}
}

func main() {
	lambda.Start(handler)
}
