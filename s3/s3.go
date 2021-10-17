package s3

type Store interface {
	ListBuckets() error
}
