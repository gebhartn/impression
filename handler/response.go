package handler

import (
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
