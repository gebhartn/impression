package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
)

func New() *session.Session {
	s := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsWest2RegionID),
	}))
	return s
}
