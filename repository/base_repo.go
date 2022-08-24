package repository

import (
	"github.com/devazizi/go-echo-media-service/infrastructure"
	"gorm.io/gorm"
)

type Connection struct {
	Store *gorm.DB
}

func NewDB(DB infrastructure.DB) Connection {
	return Connection{
		Store: DB.Connection,
	}
}
