package s3

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/service/s3"
)

type Store interface {
	ListBuckets() (*s3.ListBucketsOutput, error)
	ListObjects() (*s3.ListObjectsV2Output, error)
	ListObjectsById(id uint) (*s3.ListObjectsV2Output, error)
	UploadObject(id uint, f *multipart.FileHeader) (string, error)
}
