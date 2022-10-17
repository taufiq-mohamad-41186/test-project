package usecase

import (
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/domain"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/usecase/komoditas"
)

type Usecase struct {
	Komoditas komoditas.UsecaseItf
}

type Options struct {
	Komoditas komoditas.Options
}

func Init(opt Options, dom *domain.Domain) *Usecase {
	return &Usecase{
		Komoditas: komoditas.InitKomoditasUC(opt.Komoditas, dom.Komoditas, dom.Converter),
	}
}
