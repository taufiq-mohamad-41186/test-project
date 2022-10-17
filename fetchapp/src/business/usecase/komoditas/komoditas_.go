package komoditas

import (
	"context"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/entity"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/common"
	"strconv"
	"strings"
	"time"
)

func (k *komoditasUsecase) GetKomoditas(ctx context.Context) ([]entity.Komoditas, error) {
	var results []entity.Komoditas

	res, err := k.komoditasDom.GetKomoditas(ctx)
	if err != nil {
		return results, err
	}

	return k.cleanKomoditas(ctx, res)
}

func (k *komoditasUsecase) GetAggregate(ctx context.Context) ([]entity.AggregateKey, error) {
	results := []entity.AggregateKey{}
	res, err := k.komoditasDom.GetKomoditas(ctx)
	if err != nil {
		return results, err
	}

	komoditas, err := k.cleanKomoditas(ctx, res)
	if err != nil {
		return results, err
	}

	type RegisterAggregate struct {
		Key   entity.AggregateKey
		size  int
		Price int
	}
	registerAggregate := []RegisterAggregate{}
	for _, kom := range komoditas {
		registerAggregate = append(registerAggregate, RegisterAggregate{
			Key: entity.AggregateKey{
				AreaProvinsi: kom.AreaProvinsi,
				Week: func() int {
					_, week := kom.TglParsed.ISOWeek()
					return week
				}(),
			},
			size:  kom.Size,
			Price: int(kom.Price[0].Value),
		})
	}
	mapAggregateSize := make(map[entity.AggregateKey][]int)
	mapAggregatePrice := make(map[entity.AggregateKey][]int)
	//mapAggregate := make(map[entity.AggregateKey]entity.MapAggregate)
	for _, reg := range registerAggregate {
		mapAggregateSize[reg.Key] = append(mapAggregateSize[reg.Key], reg.size)
		mapAggregatePrice[reg.Key] = append(mapAggregatePrice[reg.Key], reg.Price)
	}
	for key, agg := range mapAggregateSize {
		results = append(results, entity.AggregateKey{
			AreaProvinsi: key.AreaProvinsi,
			Week:         key.Week,
			Aggregate: entity.Aggregate{
				Min: entity.AggregateData{
					Size:  common.FindMinValInt(agg),
					Price: common.FindMinValInt(mapAggregatePrice[key]),
				},
				Max: entity.AggregateData{
					Size:  common.FindMaxValueInt(agg),
					Price: common.FindMaxValueInt(mapAggregatePrice[key]),
				},
				Avg: entity.AggregateData{
					Size:  common.FindAvgValueInt(agg),
					Price: common.FindAvgValueInt(mapAggregatePrice[key]),
				},
				Median: entity.AggregateData{
					Size:  common.FindMedianValueInt(agg),
					Price: common.FindMedianValueInt(mapAggregatePrice[key]),
				},
			},
		})
	}

	return results, nil
}

func (k *komoditasUsecase) cleanKomoditas(ctx context.Context, res []entity.HTTPKomoditasResp) ([]entity.Komoditas, error) {
	var results []entity.Komoditas
	for _, r := range res {
		var _komoditas entity.Komoditas
		_komoditas.UUID = common.StringVal(r.UUID)
		_komoditas.Komoditas = common.StringVal(r.Komoditas)
		_komoditas.AreaProvinsi = strings.ToTitle(strings.TrimSpace(strings.ToLower(common.StringVal(r.AreaProvinsi))))
		_komoditas.AreaKota = common.StringVal(r.AreaKota)
		size, _ := strconv.Atoi(common.StringVal(r.Size))
		_komoditas.Size = size
		price, _ := strconv.Atoi(common.StringVal(r.Price))
		priceUsd, err := func(price int) (float64, error) {
			idrToUsd, err := k.converterDom.Convert(ctx, entity.CacheControl{MustRevalidate: false})
			if err != nil {
				return 0, err
			}
			return float64(price) * idrToUsd.IDRTOUSD, nil
		}(price)
		if err != nil {
			return []entity.Komoditas{}, nil
		}
		_komoditas.Price = []entity.Price{
			{
				UOM:   "IDR",
				Value: float64(price),
			},
			{
				UOM:   "USD",
				Value: priceUsd,
			},
		}
		if common.StringVal(r.TglParsed) != "" {
			custFormat := "2006-01-02 15:04:05.999+07:00"
			tglParsed, err := time.Parse(custFormat, common.StringVal(r.TglParsed))
			if err != nil {
				tglParsed, err = time.Parse(time.RFC3339, common.StringVal(r.TglParsed))
				if err != nil {
					return []entity.Komoditas{}, err
				}
			}
			_komoditas.TglParsed = tglParsed
		}

		//if common.StringVal(r.TimeStamp) != "" {
		//	unix, err := strconv.ParseFloat(common.StringVal(r.TimeStamp), 32)
		//	if err != nil {
		//		return []entity.Komoditas{}, err
		//	}
		//	timestamp := time.Unix(int64(unix), 0)
		//	_komoditas.TimeStamp = timestamp
		//}

		results = append(results, _komoditas)
	}

	return results, nil
}
