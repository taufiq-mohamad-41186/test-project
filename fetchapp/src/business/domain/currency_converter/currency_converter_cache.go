package currency_converter

import (
	"context"
	"encoding/json"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/entity"
)

func (c *converter) getCacheConverter(ctx context.Context, key string) (entity.HTTPCurrencyConverterResp, error) {
	var result entity.HTTPCurrencyConverterResp
	b, err := c.bCache.Get(key)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(b, &result); err != nil {
		return result, err
	}

	return result, nil
}

func (c *converter) createCacheConverter(ctx context.Context, key string, data entity.HTTPCurrencyConverterResp) (entity.HTTPCurrencyConverterResp, error) {
	var result entity.HTTPCurrencyConverterResp
	b, err := json.Marshal(data)
	if err != nil {
		return result, nil
	}
	if err := c.bCache.Set(key, b); err != nil {
		return result, err
	}

	return result, nil
}
