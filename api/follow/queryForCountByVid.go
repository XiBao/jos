package follow

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/follow"
	"github.com/daviddengcn/ljson"
)

type QueryForCountByVidRequest struct {
	api.BaseRequest
	ShopId uint64
}

type QueryForCountByVidResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *QueryForCountByVidData `json:"jingdong_follow_vender_read_queryForCountByVid_responce,omitempty" codec:"jingdong_follow_vender_read_queryForCountByVid_responce,omitempty"`
}

type QueryForCountByVidData struct {
	Code   string                    `json:"code,omitempty" codec:"code,omitempty"`
	Result *QueryForCountByVidResult `json:"queryforcountbyvid_result,omitempty" codec:"queryforcountbyvid_result,omitempty"`
}

type QueryForCountByVidResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	Msg  string `json:"msg,omitempty" codec:"msg,omitempty"`
}

func QueryForCountByVid(req *QueryForCountByVidRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := follow.NewQueryForCountByVidRequest()

	r.SetShopId(req.ShopId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("no result.")
	}

	var response QueryForCountByVidResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}

	if response.Data.Code != "0" {
		return 0, errors.New(response.Data.Result.Msg)
	}

	return response.Data.Result.Data, nil
}
