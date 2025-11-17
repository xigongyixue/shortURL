// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"shortURL/internal/config"
	"shortURL/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	ShortUrlModel model.ShortUrlMapModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	return &ServiceContext{
		Config: c,
		ShortUrlModel: model.NewShortUrlMapModel(conn),
	}
}
