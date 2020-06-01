package areas

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/areas"
	"github.com/daviddengcn/ljson"
)

type CityGetRequest struct {
	api.BaseRequest
	ParentId uint64
}
type CityGetResponse struct {
	ErrorResp            *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	AreasCityGetResponse *AreasCityGetResponse `json:"jingdong_areas_city_get_responce,omitempty" codec:"jingdong_areas_city_get_responce,omitempty"`
}

type AreasCityGetResponse struct {
	Code                 string                    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc            string                    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	AreasServiceResponse *BaseAreasServiceResponse `json:"baseAreaServiceResponse,omitempty" codec:"baseAreaServiceResponse,omitempty"`
}

func CityGet(req *CityGetRequest) ([]*Result, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := areas.NewAreasCityGetRequest()
	r.SetParentId(req.ParentId)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no city result")
	}

	var response CityGetResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if response.AreasCityGetResponse.Code != "0" {
		return nil, errors.New(response.AreasCityGetResponse.ErrorDesc)
	}

	if response.AreasCityGetResponse.AreasServiceResponse == nil || response.AreasCityGetResponse.AreasServiceResponse.Data == nil {
		return nil, errors.New("no city result")
	}

	return response.AreasCityGetResponse.AreasServiceResponse.Data, nil
}
