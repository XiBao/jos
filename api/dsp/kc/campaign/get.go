package campaign

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/kc/campaign"
	"github.com/daviddengcn/ljson"
)

type GetRequest struct {
	api.BaseRequest
	CampaignId uint64 `json:"campaign_id,omitempty" codec:"campaign_id,omitempty"` // 计划Id
}

type GetResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetData            `json:"jingdong_dsp_kc_campain_get_responce,omitempty" codec:"jingdong_dsp_kc_campain_get_responce,omitempty"`
}

type GetData struct {
	Result *GetResult `json:"findcampaignbyid_result"`
}

type GetResult struct {
	ResultCode string    `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string    `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Success    bool      `json:"success,omitempty" codec:"success,omitempty"`
	Value      *GetValue `json:"value,omitempty" codec:"value,omitempty"`
}

type GetValue struct {
	StartTime          string  `json:"startTime,omitempty" codec:"startTime,omitempty"`                   // 开始时间
	CampaignId         uint64  `json:"campaignId,omitempty" codec:"campaignId,omitempty"`                 // 计划Id
	Status             uint8   `json:"status,omitempty" codec:"status,omitempty"`                         // 状态
	PutType            int8    `json:"putType,omitempty" codec:"putType,omitempty"`                       // 推广类型
	Name               string  `json:"name,omitempty" codec:"name,omitempty"`                             // 姓名
	TimeRangePriceCoef string  `json:"timeRangePriceCoef,omitempty" codec:"timeRangePriceCoef,omitempty"` // 推广时间
	Yn                 int8    `json:"yn,omitempty" codec:"yn,omitempty"`                                 // 是否有效
	EndTime            string  `json:"endTime,omitempty" codec:"endTime,omitempty"`                       // 结束时间
	DayBudget          float64 `json:"dayBudget,omitempty" codec:"dayBudget,omitempty"`                   // 日限额
}

// 查询.快车.计划信息（指定计划ID）
func Get(req *GetRequest) (*GetValue, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewCampainGetRequest()
	r.SetCampaignId(req.CampaignId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result info")
	}
	var response GetResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if !response.Data.Result.Success {
		return nil, errors.New(response.Data.Result.ErrorMsg)
	}

	return response.Data.Result.Value, nil

}
