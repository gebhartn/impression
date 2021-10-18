package store

import (
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	uuid "github.com/nu7hatch/gouuid"
)

var (
	bucket = os.Getenv("BUCKET_NAME")
	prefix = os.Getenv("PREFIX_NAME")
	cdn    = os.Getenv("CDN")
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

func (s *S3Store) UploadObject(id uint, f *multipart.FileHeader) (string, error) {
	p := getUserBucket(id)
	e := path.Ext(f.Filename)
	ct := f.Header.Get("Content-Type")

	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	fp := fmt.Sprintf("%s/%s%s", p, u, e)

	file, err := f.Open()
	if err != nil {
		return "", err
	}

	if _, err = s.u.Upload(&s3manager.UploadInput{
		Bucket:      &bucket,
		Key:         &fp,
		Body:        file,
		ContentType: &ct,
	}); err != nil {
		return "", err
	}

	res := fmt.Sprintf("%s/%s/%s%s", cdn, p, u, e)

	return res, nil
}

func getUserBucket(id uint) string {
	slug := strconv.FormatInt(int64(id), 10)
	return fmt.Sprintf("%s/%s", prefix, slug)
}
