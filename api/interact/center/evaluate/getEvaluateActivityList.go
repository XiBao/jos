package center

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	center "github.com/XiBao/jos/sdk/request/interact/center/evaluate"
)

type GetEvaluateActivityListRequest struct {
	api.BaseRequest
	AppName    string `json:"appName"`    // 服务商名称
	Channel    uint8  `json:"channel"`    // 请求渠道 (PC为1, APP为2, 任务中心为3,发现频道为4, 上海运营模板为5 , 微信为 6, QQ为 7, ISV为8)
	Status     string `json:"status"`     // 活动状态0,1待开始 ，2进行中，3已结束，4, 6, -4代表已关闭 查询全部 不用传递
	StartTime  string `json:"startTime"`  // 活动开始时间
	EndTime    string `json:"endTime"`    // 活动结束时间
	PageNumber uint   `json:"pageNumber"` // 第几页
	PageSize   uint   `json:"pageSize"`   // 每一页的大小
}

type GetEvaluateActivityListResponse struct {
	ErrorResp *api.ErrorResponnse          `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetEvaluateActivityListData `json:"jingdong_com_jd_interact_center_api_service_read_EvaluateActivityReadService_getActivityListByParams_responce,omitempty" codec:"jingdong_com_jd_interact_center_api_service_read_EvaluateActivityReadService_getActivityListByParams_responce,omitempty"`
}

func (r GetEvaluateActivityListResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetEvaluateActivityListResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data.IsError() {
		return r.Data.Error()
	}
	return "no result data"
}

type GetEvaluateActivityListData struct {
	Code      string                          `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                          `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    []GetEvaluateActivityListResult `json:"GiftActivityResults,omitempty" codec:"GiftActivityResults,omitempty"`
}

func (r GetEvaluateActivityListData) IsError() bool {
	return r.Code != "0"
}

func (r GetEvaluateActivityListData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

type GetEvaluateActivityListResult struct {
	Id        uint64 `json:"id"`
	VenderId  uint64 `json:"vender_id"`
	EndTime   uint64 `json:"end_time"`
	StartTime uint64 `json:"start_time"`
	Type      string `json:"type"`
	Platform  int    `json:"platform"`
	Name      string `json:"name"`
	Status    uint8  `json:"status"`
}

func GetEvaluateActivityList(req *GetEvaluateActivityListRequest) ([]GetEvaluateActivityListResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewGetEvaluateActivityListRequest()
	r.SetAppName(req.AppName)
	r.SetChannel(req.Channel)
	r.SetPageNumber(req.PageNumber)
	r.SetPageSize(req.PageSize)

	if req.Status != "" {
		r.SetStatus(req.Status)
	}
	if req.StartTime != "" && req.EndTime != "" {
		r.SetStartTime(req.StartTime)
		r.SetEndTime(req.EndTime)
	}

	var response GetEvaluateActivityListResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result, nil
}
