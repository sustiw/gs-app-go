package config

import (
	"FirstGo/models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	return ":" + port
}

func getS3Config() (string, string, string) {
	bucket := os.Getenv("AWS_S3_BUCKET")
	fileName := os.Getenv("AWS_S3_FILE_NAME")
	region := os.Getenv("AWS_REGION")

	if bucket == "" {
		bucket = "gs-app-test"
	}
	if fileName == "" {
		fileName = "pack-details-go.json"
	}
	if region == "" {
		region = "us-east-1"
	}

	return bucket, fileName, region
}

func GetAvailablePackSizes(ctx context.Context) []int {
	bucket, file, region := getS3Config()
	fmt.Printf("Elastic Beanstalk: Fetching live data from Bucket: [%s], File: [%s] in Region: [%s]\n", bucket, file, region)

	s3Ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	cfg, err := config.LoadDefaultConfig(s3Ctx,
		config.WithRegion(region),
	)
	if err != nil {
		log.Printf("AWS Configuration Load Error: %v. Using fallback data.\n", err)
		return []int{250, 500, 1000, 2000, 5000}
	}

	s3Client := s3.NewFromConfig(cfg)

	input := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &file,
	}

	output, err := s3Client.GetObject(s3Ctx, input)
	if err != nil {
		log.Printf("S3 Network Error: GetObject failed for bucket %s, file %s: %v. Using fallback data.\n", bucket, file, err)
		return []int{250, 500, 1000, 2000, 5000}
	}
	defer output.Body.Close()

	bodyBytes, err := io.ReadAll(output.Body)
	if err != nil {
		log.Printf("Streaming Error: Failed to parse body byte stream from S3: %v\n", err)
		return []int{250, 500, 1000, 2000, 5000}
	}

	var configData models.S3PackSizesConfig
	err = json.Unmarshal(bodyBytes, &configData)
	if err != nil {
		log.Printf("JSON Parse Error: S3 file layout doesn't match our struct keys: %v\n", err)
		return []int{250, 500, 1000, 2000, 5000}
	}

	return configData.PackSizes
}
