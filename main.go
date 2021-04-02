package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
/* 	"github.com/labstack/echo" */
	pb "github.com/takahiro0530/picture_management_server/protocol/protocol/picture_management_proto"
)

type server struct{}

type Picture struct {
    PictureName string
}


func (s *server) ListPictures(ctx context.Context, req *pb.PicturesRequest) (*pb.PicturesResponce, error)) {
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

	pictureList := make([]*Picture, 0)

	for _, item := range result.Contents {
		if item != nil {
			picture := &Picture {
                PictureName: *item.Key
			}
			pictureList = append(pictureList, picture)
			fmt.Println("Name", *item.Key)
		}
	}

	res := pb.PicturesResponce{
		Picture: pictureList
	}
	
	return &res, nil 

}

func main() {
/* 
	e := echo.New()

	e.Logger.Fatal(e.Start(":8080"))
 */
	listenPort, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	pb.RegisterPictureManagemetServer(s, &server{})

	s.Serve(listenPort)

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
