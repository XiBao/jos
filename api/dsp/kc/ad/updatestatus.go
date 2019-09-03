package ad

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/kc/ad"
	"github.com/daviddengcn/ljson"
)

type AdUpdateStatusRequest struct {
	api.BaseRequest
	Id     string `json:"id,omitempty" codec:"id,omitempty"`
	Status int    `json:"status,omitempty" codec:"status,omitempty"`
}

type AdUpdateStatusResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *AdUpdateStatusData `json:"jingdong_dsp_kc_ad_updatestatus_response,omitempty" codec:"jingdong_dsp_kc_ad_updatestatus_response,omitempty"`
}

type AdUpdateStatusData struct {
	Result *DspKcAdUpdateResult `json:"updatestatus_result"`
}

type DspKcAdUpdateResult struct {
	Value      int    `json:"value,omitempty" codec:"value,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
}

// 修改创意状态
func AdUpdateStatus(req *AdUpdateStatusRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := ad.NewAdUpdateStatusRequest()

	r.SetId(req.Id)
	r.SetStatus(req.Status)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}
	var response AdUpdateStatusResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	if !response.Data.Result.Success {
		return false, errors.New(response.Data.Result.ErrorMsg)
	}

	return true, nil

}
