package areas

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/areas"
	"github.com/daviddengcn/ljson"
)

type CountyGetRequest struct {
	api.BaseRequest
	ParentId uint64
}
type CountyGetResponse struct {
	ErrorResp              *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	AreasCountyGetResponse *AreasCountyGetResponse `json:"jingdong_areas_county_get_responce,omitempty" codec:"jingdong_areas_county_get_responce,omitempty"`
}

type AreasCountyGetResponse struct {
	Code                 string                    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc            string                    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	AreasServiceResponse *BaseAreasServiceResponse `json:"baseAreaServiceResponse,omitempty" codec:"baseAreaServiceResponse,omitempty"`
}

func CountyGet(req *CountyGetRequest) ([]*Result, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := areas.NewAreasCountyGetRequest()
	r.SetParentId(req.ParentId)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no county result")
	}

	var response CountyGetResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if response.AreasCountyGetResponse.Code != "0" {
		return nil, errors.New(response.AreasCountyGetResponse.ErrorDesc)
	}

	if response.AreasCountyGetResponse.AreasServiceResponse == nil || response.AreasCountyGetResponse.AreasServiceResponse.Data == nil {
		return nil, errors.New("no county result")
	}

	return response.AreasCountyGetResponse.AreasServiceResponse.Data, nil
}
