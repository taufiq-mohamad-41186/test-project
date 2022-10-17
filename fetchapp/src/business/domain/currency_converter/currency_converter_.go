package currency_converter

import (
	"context"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/entity"
)

func (c *converter) Convert(ctx context.Context, cc entity.CacheControl) (entity.HTTPCurrencyConverterResp, error) {
	if cc.MustRevalidate {
		result, err := c.getHTTPConvert(ctx)
		if err != nil {
			return entity.HTTPCurrencyConverterResp{}, err
		}
		return c.createCacheConverter(ctx, "USD_IDR", result)
	} else {
		// implement cache
		result, err := c.getCacheConverter(ctx, "USD_IDR")
		if err != nil {
			result, err := c.getHTTPConvert(ctx)
			if err != nil {
				return result, err
			}
			return c.createCacheConverter(ctx, "USD_IDR", result)
		}
		return result, nil
	}
}
