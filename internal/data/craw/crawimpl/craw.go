package crawimpl

import (
	"fmt"
	"live/internal/data/craw"
	"live/pkg/options"
	"net/http"
	"sync"
)

type dataCraw struct {
	httpClient *http.Client
}

// HuYaCraw implements craw.Factory
func (dc *dataCraw) HuYaCraw() craw.HuYaCraw {
	return newHuya(dc)
}

var (
	crawFactory craw.Factory
	once        sync.Once
)

func GetCreaFactoryOr(opts *options.CrawOptions) (craw.Factory, error) {
	if opts == nil && crawFactory == nil {
		return nil, fmt.Errorf("failed to get craw fatory")
	}
	var err error
	var httpClient *http.Client
	if err != nil {
		return nil, err
	}
	once.Do(func() {
		httpClient = &http.Client{
			Timeout: opts.TimeOut,
		}
		crawFactory = &dataCraw{httpClient}
	})
	if crawFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get craw fatory,crawFactory: %+v, error: %w", crawFactory, err)
	}
	return crawFactory, nil
}
