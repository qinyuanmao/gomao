package db

import (
	"context"

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

func (m *MaoDB) WithContext(ctx context.Context) *gorm.DB {
	d := m.DB.WithContext(ctx)
	return d
}
