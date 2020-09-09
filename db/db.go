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
