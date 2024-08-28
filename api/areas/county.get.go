package areas

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/areas"
)

type CountyGetRequest struct {
	api.BaseRequest
	ParentId uint64
}
type CountyGetResponse struct {
	ErrorResp              *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	AreasCountyGetResponse *AreasCountyGetResponse `json:"jingdong_areas_county_get_responce,omitempty" codec:"jingdong_areas_county_get_responce,omitempty"`
}

func (r CountyGetResponse) IsError() bool {
	return r.ErrorResp != nil || r.AreasCountyGetResponse == nil || r.AreasCountyGetResponse.IsError()
}

func (r CountyGetResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.AreasCountyGetResponse != nil {
		return r.AreasCountyGetResponse.Error()
	}
	return "no result data"
}

type AreasCountyGetResponse struct {
	Code                 string                    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc            string                    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	AreasServiceResponse *BaseAreasServiceResponse `json:"baseAreaServiceResponse,omitempty" codec:"baseAreaServiceResponse,omitempty"`
}

func (r AreasCountyGetResponse) IsError() bool {
	return r.Code != "0"
}

func (r AreasCountyGetResponse) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

func CountyGet(ctx context.Context, req *CountyGetRequest) ([]Result, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := areas.NewAreasCountyGetRequest()
	r.SetParentId(req.ParentId)

	var response CountyGetResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.AreasCountyGetResponse.AreasServiceResponse.Data, nil
}
