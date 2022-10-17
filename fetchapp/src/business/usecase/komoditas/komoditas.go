package komoditas

import (
	"context"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/domain/currency_converter"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/domain/komoditas"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/entity"
)

type UsecaseItf interface {
	GetKomoditas(ctx context.Context) ([]entity.Komoditas, error)
	GetAggregate(ctx context.Context) ([]entity.AggregateKey, error)
}

type komoditasUsecase struct {
	opt          Options
	komoditasDom komoditas.DomainItf
	converterDom currency_converter.DomainItf
}

type Options struct{}

func InitKomoditasUC(opt Options, komoditasDom komoditas.DomainItf, converterDom currency_converter.DomainItf) UsecaseItf {
	return &komoditasUsecase{
		opt:          opt,
		komoditasDom: komoditasDom,
		converterDom: converterDom,
	}
}
