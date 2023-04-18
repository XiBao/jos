package areas

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/areas"
)

type ProvinceGetRequest struct {
	api.BaseRequest
}
type ProvinceGetResponse struct {
	ErrorResp                *api.ErrorResponnse       `json:"error_response,omitempty" codec:"error_response,omitempty"`
	AreasProvinceGetResponse *AreasProvinceGetResponse `json:"jingdong_areas_province_get_responce,omitempty" codec:"jingdong_areas_province_get_responce,omitempty"`
}

type AreasProvinceGetResponse struct {
	Code                 string                    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc            string                    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	AreasServiceResponse *BaseAreasServiceResponse `json:"baseAreaServiceResponse,omitempty" codec:"baseAreaServiceResponse,omitempty"`
}

func ProvinceGet(req *ProvinceGetRequest) ([]*Result, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := areas.NewAreasProvinceGetRequest()
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no province result")
	}

	var response ProvinceGetResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if response.AreasProvinceGetResponse.Code != "0" {
		return nil, errors.New(response.AreasProvinceGetResponse.ErrorDesc)
	}

	if response.AreasProvinceGetResponse.AreasServiceResponse == nil || response.AreasProvinceGetResponse.AreasServiceResponse.Data == nil {
		return nil, errors.New("no province result")
	}

	return response.AreasProvinceGetResponse.AreasServiceResponse.Data, nil
}
