package ad

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/kc/ad"
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

func (r AdUpdateStatusResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r AdUpdateStatusResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type AdUpdateStatusData struct {
	Result *DspKcAdUpdateResult `json:"updatestatus_result"`
}

func (r AdUpdateStatusData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r AdUpdateStatusData) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type DspKcAdUpdateResult struct {
	Value      int    `json:"value,omitempty" codec:"value,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
}

func (r DspKcAdUpdateResult) IsError() bool {
	return !r.Success
}

func (r DspKcAdUpdateResult) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.ResultCode, r.ErrorMsg)
}

// 修改创意状态
func AdUpdateStatus(req *AdUpdateStatusRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := ad.NewAdUpdateStatusRequest()

	r.SetId(req.Id)
	r.SetStatus(req.Status)

	var response AdUpdateStatusResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return true, nil

}
