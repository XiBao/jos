package areas

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/areas"
)

type CityGetRequest struct {
	api.BaseRequest
	ParentId uint64
}
type CityGetResponse struct {
	ErrorResp            *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	AreasCityGetResponse *AreasCityGetResponse `json:"jingdong_areas_city_get_responce,omitempty" codec:"jingdong_areas_city_get_responce,omitempty"`
}

func (r CityGetResponse) IsError() bool {
	return r.ErrorResp != nil || r.AreasCityGetResponse == nil || r.AreasCityGetResponse.IsError()
}

func (r CityGetResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.AreasCityGetResponse != nil {
		return r.AreasCityGetResponse.Error()
	}
	return "no result data"
}

type AreasCityGetResponse struct {
	Code                 string                    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc            string                    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	AreasServiceResponse *BaseAreasServiceResponse `json:"baseAreaServiceResponse,omitempty" codec:"baseAreaServiceResponse,omitempty"`
}

func (r AreasCityGetResponse) IsError() bool {
	return r.Code != "0"
}

func (r AreasCityGetResponse) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

func CityGet(ctx context.Context, req *CityGetRequest) ([]Result, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := areas.NewAreasCityGetRequest()
	r.SetParentId(req.ParentId)

	var response CityGetResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.AreasCityGetResponse.AreasServiceResponse.Data, nil
}
