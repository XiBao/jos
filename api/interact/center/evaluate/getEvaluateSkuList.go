package center

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api/util"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	center "github.com/XiBao/jos/sdk/request/interact/center/evaluate"
)

type GetEvaluateSkuListRequest struct {
	api.BaseRequest
	AppName    string `json:"appName"`    // 服务商名称
	Channel    uint8  `json:"channel"`    // 请求渠道 (PC为1, APP为2, 任务中心为3,发现频道为4, 上海运营模板为5 , 微信为 6, QQ为 7, ISV为8)
	StartTime  string `json:"startTime"`  // 活动开始时间
	ActivityId uint64 `json:"activityId"` // 活动id
	EndTime    string `json:"endTime"`    // 活动结束时间
	PageNumber uint   `json:"pageNumber"` // 第几页
	PageSize   uint   `json:"pageSize"`   // 每一页的大小
}

type GetEvaluateSkuListResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetEvaluateSkuListData `json:"jingdong_com_jd_interact_center_api_service_read_EvaluateSkuReadService_getSkuListByPatams_responce,omitempty" codec:"jingdong_com_jd_interact_center_api_service_read_EvaluateSkuReadService_getSkuListByPatams_responce,omitempty"`
}

type GetEvaluateSkuListData struct {
	Code      string         `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string         `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    []*EvaluateSku `json:"result,omitempty" codec:"result,omitempty"`
}

type EvaluateSku struct {
	VenderId    uint64  `json:"vender_id"`
	SkuName     string  `json:"sku_name"`
	ActivityId  uint64  `json:"activity_id"`
	SkuId       uint64  `json:"sku_id"`
	WareId      uint64  `json:"ware_id"`
	BusinessId  uint64  `json:"business_id"`
	Price       float64 `json:"price"`
	Cate1stCode uint64  `json:"cate_1st_code"`
	Cate1stName string  `json:"cate_1st_name"`
	Cate2ndCode uint64  `json:"cate_2nd_code"`
	Cate2ndName string  `json:"cate_2nd_name"`
	Cate3rdCode uint64  `json:"cate_3rd_code"`
	Cate3rdName string  `json:"cate_3rd_name"`
	Logo        string  `json:"logo"`
	StockNum    int     `json:"stock_num"`
}

func GetEvaluateSkuList(req *GetEvaluateSkuListRequest) ([]*EvaluateSku, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewGetEvaluateSkuListRequest()
	r.SetActivityId(req.ActivityId)
	r.SetAppName(req.AppName)
	r.SetChannel(req.Channel)
	r.SetPageNumber(req.PageNumber)
	r.SetPageSize(req.PageSize)

	if req.StartTime != "" && req.EndTime != "" {
		r.SetStartTime(req.StartTime)
		r.SetEndTime(req.EndTime)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	result = util.RemoveJsonSpace(result)

	var response GetEvaluateSkuListResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}
	if response.Data.Result == nil {
		return nil, errors.New("No result.")
	}

	return response.Data.Result, nil
}
