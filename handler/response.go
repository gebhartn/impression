package handler

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gebhartn/impress/model"
	"github.com/gebhartn/impress/utils"
)

type userResponse struct {
	User struct {
		Username string `json:"username"`
		Token    string `json:"token"`
	} `json:"user"`
}

func newUserResponse(u *model.User) *userResponse {
	r := userResponse{}
	r.User.Username = u.Username
	r.User.Token = utils.GenerateJWT(u.ID)

	return &r
}

type imageUploadResponse struct {
	Image struct {
		Key string `json:"key"`
	} `json:"image"`
}

func newImageUploadResponse(i *model.Image) *imageUploadResponse {
	r := imageUploadResponse{}
	r.Image.Key = i.Key

	return &r
}

type userImage struct {
	URL          string    `json:"url"`
	LastModified time.Time `json:"last_modified"`
	// TODO: Add reference to uploader
}

type userImagesReponse struct {
	Images []*userImage `json:"images"`
}

func newUserImagesReponse(is *s3.ListObjectsV2Output) *userImagesReponse {
	r := userImagesReponse{Images: []*userImage{}}
	for _, i := range is.Contents {
		u := fmt.Sprintf("%s/%s", os.Getenv("CDN"), *i.Key)
		r.Images = append(r.Images, &userImage{
			URL:          u,
			LastModified: *i.LastModified,
		})
	}

	return &r
}
