package komoditas

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/taufiq-mohamad-41186/test-project/fetchapp/src/business/entity"
	"io/ioutil"
)

func (k *komoditas) getHTTPKomoditas(ctx context.Context) ([]entity.HTTPKomoditasResp, error) {
	var result []entity.HTTPKomoditasResp
	resp, err := k.httpClient.Get(k.opt.URL.BaseURL + k.opt.URL.GetKomoditas)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		// TODO handling http errors
		return result, errors.New("Something went wrong")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(b, &result); err != nil {
		return result, err
	}

	return result, nil
}
