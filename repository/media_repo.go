package repository

import "github.com/devazizi/go-echo-media-service/entity"

type MediaRepositoryInterface interface {
	StoreMedia(media entity.Media) entity.Media
}

func (c Connection) StoreMedia(media entity.Media) entity.Media {

	c.Store.Create(&media)

	return media
}
