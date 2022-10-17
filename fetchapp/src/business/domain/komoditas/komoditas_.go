package komoditas

import (
	"context"
	"encoding/json"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/entity"
	"io/ioutil"
)

func (k *komoditas) GetKomoditas(ctx context.Context) ([]entity.HTTPKomoditasResp, error) {
	var result []entity.HTTPKomoditasResp
	resp, err := k.httpClient.Get(k.opt.URL.BaseURL + k.opt.URL.GetKomoditas)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(b, &result); err != nil {
		return result, err
	}

	return result, nil
}
