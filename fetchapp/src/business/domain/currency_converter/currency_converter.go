package currency_converter

import (
	"context"
	"github.com/allegro/bigcache"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/entity"
	"net/http"
)

type DomainItf interface {
	Convert(ctx context.Context, cc entity.CacheControl) (entity.HTTPCurrencyConverterResp, error)
}

type converter struct {
	httpClient http.Client
	bCache     *bigcache.BigCache
	opt        Options
}

type Options struct {
	URL URLOptions
}

type URLOptions struct {
	BaseURL      string
	GetKomoditas string
}

func InitCurrencyConverterDomain(httpClient http.Client, bCache *bigcache.BigCache, opt Options) DomainItf {
	return &converter{
		httpClient: httpClient,
		bCache:     bCache,
		opt:        opt,
	}
}
