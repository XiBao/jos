package follow

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/follow"
)

type IsFollowByPinAndVidRequest struct {
	api.BaseRequest
	Pin    string `json:"pin,omitempty" codec:"pin,omitempty"`         //
	ShopId uint64 `json:"shop_id,omitempty" codec:"shop_id,omitempty"` // 自定义返回字段
}

type IsFollowByPinAndVidResponse struct {
	ErrorResp *api.ErrorResponnse      `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *IsFollowByPinAndVidData `json:"jingdong_follow_vender_read_isFollowByPinAndVid_responce,omitempty" codec:"jingdong_follow_vender_read_isFollowByPinAndVid_responce,omitempty"`
}

func (r IsFollowByPinAndVidResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r IsFollowByPinAndVidResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type IsFollowByPinAndVidData struct {
	Code   string                     `json:"code,omitempty" codec:"code,omitempty"`
	Result *IsFollowByPinAndVidResult `json:"isfollowbypinandvid_result,omitempty" codec:"isfollowbypinandvid_result,omitempty"`
}

func (r IsFollowByPinAndVidData) IsError() bool {
	return r.Result == nil
}

func (r IsFollowByPinAndVidData) Error() string {
	return "no result data"
}

type IsFollowByPinAndVidResult struct {
	Data bool   `json:"data,omitempty" codec:"data,omitempty"`
	Code string `json:"code,omitempty" codec:"code,omitempty"`
}

func IsFollowByPinAndVid(req *IsFollowByPinAndVidRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := follow.NewIsFollowByPinAndVidRequest()
	r.SetPin(req.Pin)
	r.SetShopId(req.ShopId)

	var response IsFollowByPinAndVidResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Data.Result.Data, nil
}
