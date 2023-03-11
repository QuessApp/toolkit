package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Credentials represents the AWS S3 credentials.
type S3Credentials struct {
	ID     string
	Secret string
	Token  string
}

// Configure configures AWS S3 client.
func Configure(region *string, S3credentials *S3Credentials) (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      region,
		Credentials: credentials.NewStaticCredentials(S3credentials.ID, S3credentials.Secret, S3credentials.Token),
	})

	if err != nil {
		return nil, err
	}

	return s3.New(sess), nil
}

// UploadFile uploads file to S3.
func UploadFile(client *s3.S3, fileName, bucketName string, file io.ReadSeeker) (*s3.PutObjectOutput, error) {
	return client.PutObject(&s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    aws.String(fileName),
		Body:   file,
	})
}
