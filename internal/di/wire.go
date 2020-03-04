package di

import (
	"admin/internal/dao"
	"admin/internal/server/http"
	"admin/internal/server/thrift"
	"admin/internal/service"
	"admin/pkg/conf"
	"github.com/google/wire"
	"github.com/kataras/iris"
)

func config1() (*conf.Yaml, error) {
	return conf.New("configs/conf.yaml")
}

func InitApp1() (*iris.Application, func(), error) {
	panic(wire.Build(config, dao.New, dao.NewDao, service.New, dao.NewMongo, thrift.New, http.New))
}
