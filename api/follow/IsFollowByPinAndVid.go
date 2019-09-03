package follow

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/follow"
	"github.com/daviddengcn/ljson"
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

type IsFollowByPinAndVidData struct {
	Code   string                     `json:"code,omitempty" codec:"code,omitempty"`
	Result *IsFollowByPinAndVidResult `json:"isfollowbypinandvid_result,omitempty" codec:"isfollowbypinandvid_result,omitempty"`
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("No result info.")
	}
	var response IsFollowByPinAndVidResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}

	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	return response.Data.Result.Data, nil
}
