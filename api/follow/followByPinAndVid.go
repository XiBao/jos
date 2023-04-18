package follow

import (
	"encoding/json"
	"errors"

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

type FollowByPinAndVidData struct {
	Code   string                   `json:"code,omitempty" codec:"code,omitempty"`
	Result *FollowByPinAndVidResult `json:"followbypinandvid_result,omitempty" codec:"followbypinandvid_result,omitempty"`
}

type FollowByPinAndVidResult struct {
	Data bool   `json:"data,omitempty" codec:"data,omitempty"`
	Code string `json:"code,omitempty" codec:"code,omitempty"`
}

func FollowByPinAndVid(req *FollowByPinAndVidRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := follow.NewFollowByPinAndVidRequest()
	r.SetPin(req.Pin)
	r.SetShopId(req.ShopId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("No result info.")
	}
	var response FollowByPinAndVidResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}

	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	return response.Data.Result.Data, nil
}
