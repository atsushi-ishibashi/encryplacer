package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

var (
	s3svc s3iface.S3API
)

var (
	bucket     = flag.String("bucket", "", "bucket name")
	kmsID      = flag.String("kms", "", "KMS Key ID for encryption")
	region     = flag.String("region", "", "AWS region. priority: arg>AWS_REGION>AWS_DEFAULT_REGION")
	suffix     = flag.String("suffix", "", "suffix to filter objects")
	contain    = flag.String("contain", "", "contain to filter objects")
	concurrent = flag.Int("concurrent", 3, "concurrent")
)

func main() {
	flag.Parse()
	if *bucket == "" {
		log.Fatal("required bucket")
	}
	if *kmsID == "" {
		log.Fatal("required kms")
	}

	var regionStr string
	if os.Getenv("AWS_DEFAULT_REGION") != "" {
		regionStr = os.Getenv("AWS_DEFAULT_REGION")
	}
	if os.Getenv("AWS_REGION") != "" {
		regionStr = os.Getenv("AWS_REGION")
	}
	if *region != "" {
		regionStr = *region
	}
	s3svc = s3.New(session.New(), aws.NewConfig().WithRegion(regionStr))

	filter := Filter{
		Suffix:  *suffix,
		Contain: *contain,
	}

	targetKeys := make([]string, 0)

	input := &s3.ListObjectsV2Input{
		Bucket: bucket,
	}

	err := s3svc.ListObjectsV2Pages(input,
		func(page *s3.ListObjectsV2Output, lastPage bool) bool {
			for _, v := range page.Contents {
				key := *v.Key
				if !filter.match(key) {
					continue
				}
				targetKeys = append(targetKeys, key)
			}
			return page.IsTruncated != nil && !*page.IsTruncated
		})
	if err != nil {
		log.Println(err)
	}

	limit := make(chan struct{}, *concurrent)

	var wg sync.WaitGroup
	for _, key := range targetKeys {
		wg.Add(1)
		go func(k string) {
			limit <- struct{}{}
			defer wg.Done()
			if err := replaceEncryption(k); err != nil {
				log.Println(err)
			}
			<-limit
		}(key)
	}
	wg.Wait()
}

func replaceEncryption(key string) error {
	resp, err := s3svc.GetObject(&s3.GetObjectInput{
		Bucket: bucket,
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	buffer := make([]byte, *resp.ContentLength)
	resp.Body.Read(buffer)

	putinput := &s3.PutObjectInput{
		Bucket:               bucket,
		Body:                 bytes.NewReader(buffer),
		Key:                  aws.String(key),
		SSEKMSKeyId:          kmsID,
		ServerSideEncryption: aws.String("aws:kms"),
	}
	_, err = s3svc.PutObject(putinput)
	return err
}
