package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// sessionの作成
	// Must関数を使うことでエラーハンドリングもやってくれている
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "aws-go-test",
		SharedConfigState: session.SharedConfigEnable,
	}))

	// S3クライアント作成
	client := s3.New(sess)

	// バケットの一覧取得
	// result, _ := svc.ListBuckets(nil)
	bucketName := "aws-s3-my-picture"

	// ctx := context.Background()

	result, err := client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Contents {
		fmt.Println("Name", *item.Key)
	}

	/* 	// ファイルを開く
	   	filePath := "./test.txt"
	   	file, err := os.Open(filePath)
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	   	bucketName := ""
	   	key := ""

	   	// uploaderの作作成
	   	uploader := s3manager.NewUploader(sess)
	   	result, err := uploader.Upload(&s3manager.UploadInput{
	   		Bucket: aws.String(bucketName),
	   		Key:    aws.String(key),
	   		Body:   file,
	   	})
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	   	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	*/
	// downloaderの作成
	/* 	downloader := s3manager.NewDownloader(sess)
	   	n, err = downloader.Download(file, &s3.GetObjectInput{
	   		Bucket: aws.String(bucketName),
	   		Key:    aws.String(key),
	   	})
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	   	fmt.Printf(n) */
}

/*
func upload() {

}

func download() {

}

func (c *S3) GetObject(input *GetObjectInput) (*GetObjectOutput, error) {

}
*/
