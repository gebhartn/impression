package store

import (
	"fmt"
	"mime/multipart"
	"path"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	uuid "github.com/nu7hatch/gouuid"
)

var (
	bucket = "impression-int"
	prefix = "Users"
)

type S3Store struct {
	svc *s3.S3
	u   *s3manager.Uploader
}

func NewS3Store(s *session.Session) *S3Store {
	svc := s3.New(s)
	u := s3manager.NewUploader(s)
	return &S3Store{
		svc: svc,
		u:   u,
	}
}

func (s *S3Store) ListBuckets() (*s3.ListBucketsOutput, error) {
	res, err := s.svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *S3Store) ListObjects() (*s3.ListObjectsV2Output, error) {
	res, err := s.svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: &bucket,
		Prefix: &prefix,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *S3Store) ListObjectsById(id uint) (*s3.ListObjectsV2Output, error) {
	p := getUserBucket(id)

	res, err := s.svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: &bucket,
		Prefix: &p,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *S3Store) UploadObject(id uint, f *multipart.FileHeader) (*s3manager.UploadOutput, error) {
	p := getUserBucket(id)

	key, err := getKeyName(f.Filename, p)
	if err != nil {
		return nil, err
	}

	file, err := f.Open()
	if err != nil {
		return nil, err
	}

	res, err := s.u.Upload(&s3manager.UploadInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   file,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getKeyName(f string, p string) (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	e := path.Ext(f)
	key := fmt.Sprintf("%s/%s%s", p, u, e)
	return key, nil
}

func getUserBucket(id uint) string {
	slug := strconv.FormatInt(int64(id), 10)
	return fmt.Sprintf("%s/%s", prefix, slug)
}
