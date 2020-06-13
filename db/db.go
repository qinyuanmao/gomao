package db

import (
	"context"

	"github.com/jinzhu/gorm"
	otgorm "github.com/smacker/opentracing-gorm"
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
	d := otgorm.SetSpanToGorm(ctx, m.DB)
	return d
}

// TODO  添加通用查询方法
