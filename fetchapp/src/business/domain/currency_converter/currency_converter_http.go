package currency_converter

import (
	"context"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/entity"
)

func (c *converter) getHTTPConvert(ctx context.Context) (entity.HTTPCurrencyConverterResp, error) {
	//resp, err := c.httpClient.Get(c.opt.URL.BaseURL + c.opt.URL.GetKomoditas)
	//if err != nil {
	//	return err
	//}
	//defer resp.Body.Close()
	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return err
	//}

	//if err := json.Unmarshal(b, &result); err != nil {
	//	return result, err
	//}

	// mock response
	return entity.HTTPCurrencyConverterResp{
		IDRTOUSD: 0.000065,
	}, nil
}
