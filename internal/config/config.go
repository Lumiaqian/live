package config

import (
	"live/pkg/options"

	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Craw *options.CrawOptions
}
