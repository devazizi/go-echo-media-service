package infrastructure

import (
	"github.com/devazizi/go-echo-media-service/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Connection *gorm.DB
}

func NewDB(dsn string) DB {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("fail to connect db")
	}

	var entities []any = []any{
		&entity.Media{},
	}

	if err := db.AutoMigrate(entities...); err != nil {
		panic("can not migrate database")
	}

	return DB{
		Connection: db,
	}
}
