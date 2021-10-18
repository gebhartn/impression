package image

import "github.com/gebhartn/impress/model"

type Store interface {
	Create(*model.Image) error
}
