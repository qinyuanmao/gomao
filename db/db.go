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

func Paginate(page, pageSize *int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if *page == 0 {
			*page = 1
		}
		if *pageSize == 0 {
			*pageSize = 20
		}
		offset := (*page - 1) * (*pageSize)
		return db.Offset(int(offset)).Limit(int(*pageSize))
	}
}
