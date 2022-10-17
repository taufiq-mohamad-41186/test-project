package komoditas

import (
	"context"
	"net/http"

	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/entity"
)

type DomainItf interface {
	GetKomoditas(ctx context.Context) ([]entity.HTTPKomoditasResp, error)
}

type komoditas struct {
	httpClient http.Client
	opt        Options
}

type Options struct {
	URL URLOptions
}

type URLOptions struct {
	BaseURL      string
	GetKomoditas string
}

func InitKomoditasDomain(httpClient http.Client, opt Options) DomainItf {
	return &komoditas{
		httpClient: httpClient,
		opt:        opt,
	}
}
