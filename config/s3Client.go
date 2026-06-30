package config

import (
	"FirstGo/models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func getS3Config() (string, string) {
	bucket := os.Getenv("AWS_S3_BUCKET")
	fileName := os.Getenv("AWS_S3_FILE_NAME")

	// Local development fallback defaults
	if bucket == "" {
		bucket = "gs-app-test"
	}
	if fileName == "" {
		fileName = "pack-details-go.json"
	}

	return bucket, fileName
}

func GetAvailablePackSizes() []int {
	bucket, file := getS3Config()
	fmt.Printf("Elastic Beanstalk: Fetching live data from Bucket: [%s], File: [%s]\n", bucket, file)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("AWS Configuration Load Error: %v\n", err)
		return []int{250, 500, 1000, 2000, 5000}
	}

	s3Client := s3.NewFromConfig(cfg)
	input := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &file,
	}
	output, err := s3Client.GetObject(context.TODO(), input)
	if err != nil {
		log.Printf("S3 Network Error: GetObject failed for bucket %s, file %s: %v\n", bucket, file, err)
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
