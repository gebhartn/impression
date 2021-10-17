package store

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Store struct {
	svc *s3.S3
}

func NewS3Store(s *session.Session) *S3Store {
	svc := s3.New(s)
	return &S3Store{
		svc: svc,
	}
}

func (s *S3Store) ListBuckets() error {
	res, err := s.svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return err
	}

	fmt.Printf("%v", res)

	return nil
}
