package db

import (
	"gorm.io/gorm"
)

type databaseType string

const (
	Mysql    databaseType = "mysql"
	Postgres databaseType = "postgres"
)

type MaoDB struct {
	DatabaseType databaseType
	*gorm.DB
}

func Paginate(page, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * (pageSize)
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
