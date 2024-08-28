package follow

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/follow"
)

type FollowByPinAndVidRequest struct {
	api.BaseRequest
	Pin    string `json:"pin,omitempty" codec:"pin,omitempty"`         //
	ShopId uint64 `json:"shop_id,omitempty" codec:"shop_id,omitempty"` // 自定义返回字段
}

type FollowByPinAndVidResponse struct {
	ErrorResp *api.ErrorResponnse    `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FollowByPinAndVidData `json:"jingdong_follow_vender_write_followByPinAndVid_responce,omitempty" codec:"jingdong_follow_vender_write_followByPinAndVid_responce,omitempty"`
}

func (r FollowByPinAndVidResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r FollowByPinAndVidResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type FollowByPinAndVidData struct {
	Code   string                   `json:"code,omitempty" codec:"code,omitempty"`
	Result *FollowByPinAndVidResult `json:"followbypinandvid_result,omitempty" codec:"followbypinandvid_result,omitempty"`
}

func (r FollowByPinAndVidData) IsError() bool {
	return r.Result == nil
}

func (r FollowByPinAndVidData) Error() string {
	return "no result data"
}

type FollowByPinAndVidResult struct {
	Data bool   `json:"data,omitempty" codec:"data,omitempty"`
	Code string `json:"code,omitempty" codec:"code,omitempty"`
}

func FollowByPinAndVid(ctx context.Context, req *FollowByPinAndVidRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := follow.NewFollowByPinAndVidRequest()
	r.SetPin(req.Pin)
	r.SetShopId(req.ShopId)

	var response FollowByPinAndVidResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Data.Result.Data, nil
}
