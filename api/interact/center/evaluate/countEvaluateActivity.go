package center

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api/util"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	center "github.com/XiBao/jos/sdk/request/interact/center/evaluate"
)

type CountEvaluateActivityRequest struct {
	api.BaseRequest
	AppName    string `json:"appName"` // 服务商名称
	Channel    uint8  `json:"channel"` // 请求渠道 (PC为1, APP为2, 任务中心为3,发现频道为4, 上海运营模板为5 , 微信为 6, QQ为 7, ISV为8)
	Name       string `json:"name"`
	Status     string `json:"status"`     // 活动状态0,1待开始 ，2进行中，3已结束，4, 6, -4代表已关闭 查询全部 不用传递
	StartTime  string `json:"startTime"`  // 活动开始时间
	EndTime    string `json:"endTime"`    // 活动结束时间
	PageNumber uint   `json:"pageNumber"` // 第几页
	PageSize   uint   `json:"pageSize"`   // 每一页的大小
}

type CountEvaluateActivityResponse struct {
	ErrorResp *api.ErrorResponnse        `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CountEvaluateActivityData `json:"jingdong_com_jd_interact_center_api_service_read_EvaluateActivityReadService_countByParams_responce,omitempty" codec:"jingdong_com_jd_interact_center_api_service_read_EvaluateActivityReadService_countByParams_responce,omitempty"`
}

type CountEvaluateActivityData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    uint   `json:"result,omitempty" codec:"result,omitempty"`
}

func CountEvaluateActivity(req *CountEvaluateActivityRequest) (uint, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewCountEvaluateActivityRequest()
	r.SetAppName(req.AppName)
	r.SetChannel(req.Channel)
	r.SetPageNumber(req.PageNumber)
	r.SetPageSize(req.PageSize)

	if req.Name != "" {
		r.SetName(req.Name)
	}

	if req.Status != "" {
		r.SetStatus(req.Status)
	}

	if req.StartTime != "" && req.EndTime != "" {
		r.SetStartTime(req.StartTime)
		r.SetEndTime(req.EndTime)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	result = util.RemoveJsonSpace(result)

	var response CountEvaluateActivityResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return 0, errors.New(response.Data.ErrorDesc)
	}

	return response.Data.Result, nil
}
