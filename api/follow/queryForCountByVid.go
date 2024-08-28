package follow

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/follow"
)

type QueryForCountByVidRequest struct {
	api.BaseRequest
	ShopId uint64
}

type QueryForCountByVidResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *QueryForCountByVidData `json:"jingdong_follow_vender_read_queryForCountByVid_responce,omitempty" codec:"jingdong_follow_vender_read_queryForCountByVid_responce,omitempty"`
}

func (r QueryForCountByVidResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r QueryForCountByVidResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type QueryForCountByVidData struct {
	Code   string                    `json:"code,omitempty" codec:"code,omitempty"`
	Result *QueryForCountByVidResult `json:"queryforcountbyvid_result,omitempty" codec:"queryforcountbyvid_result,omitempty"`
}

func (r QueryForCountByVidData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r QueryForCountByVidData) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type QueryForCountByVidResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	Msg  string `json:"msg,omitempty" codec:"msg,omitempty"`
}

func (r QueryForCountByVidResult) IsError() bool {
	return r.Code != "0" && r.Code != "200"
}

func (r QueryForCountByVidResult) Error() string {
	return sdk.ErrorString(r.Code, r.Msg)
}

func QueryForCountByVid(ctx context.Context, req *QueryForCountByVidRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := follow.NewQueryForCountByVidRequest()

	r.SetShopId(req.ShopId)

	var response QueryForCountByVidResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Data.Result.Data, nil
}
