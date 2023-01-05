package svc

import (
	"live/internal/config"
	"live/internal/data/craw"
	"live/internal/data/craw/crawimpl"
)

type ServiceContext struct {
	Config config.Config
	HuYa   craw.HuYaCraw
}

func NewServiceContext(c config.Config) *ServiceContext {
	dataCraw, err := crawimpl.GetCreaFactoryOr(c.Craw)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		HuYa:   dataCraw.HuYaCraw(),
	}
}
