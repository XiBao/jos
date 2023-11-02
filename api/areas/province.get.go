package areas

import (
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

func (r ProvinceGetResponse) IsError() bool {
	return r.ErrorResp != nil || r.AreasProvinceGetResponse == nil || r.AreasProvinceGetResponse.IsError()
}

func (r ProvinceGetResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.AreasProvinceGetResponse != nil {
		return r.AreasProvinceGetResponse.Error()
	}
	return "no result data"
}

type AreasProvinceGetResponse struct {
	Code                 string                    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc            string                    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	AreasServiceResponse *BaseAreasServiceResponse `json:"baseAreaServiceResponse,omitempty" codec:"baseAreaServiceResponse,omitempty"`
}

func (r AreasProvinceGetResponse) IsError() bool {
	return r.Code != "0"
}

func (r AreasProvinceGetResponse) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

func ProvinceGet(req *ProvinceGetRequest) ([]Result, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := areas.NewAreasProvinceGetRequest()

	var response ProvinceGetResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.AreasProvinceGetResponse.AreasServiceResponse.Data, nil
}
