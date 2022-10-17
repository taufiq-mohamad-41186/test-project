package domain

import (
	"github.com/allegro/bigcache"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/domain/currency_converter"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/domain/komoditas"
	"net/http"
)

type Domain struct {
	Komoditas komoditas.DomainItf
	Converter currency_converter.DomainItf
}

type Options struct {
	Komoditas komoditas.Options
	Converter currency_converter.Options
}

func Init(httpClient http.Client, bCache *bigcache.BigCache, opt Options) *Domain {
	return &Domain{
		Komoditas: komoditas.InitKomoditasDomain(
			httpClient,
			opt.Komoditas,
		),
		Converter: currency_converter.InitCurrencyConverterDomain(
			httpClient,
			bCache,
			opt.Converter,
		),
	}
}
