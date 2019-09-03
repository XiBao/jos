package report

import "github.com/XiBao/jos/sdk"

type QueryCampDailySumRequest struct {
	Request *sdk.Request
}

func NewQueryCampDailySumRequest() (req *QueryCampDailySumRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.report.queryCampDailySum", Params: make(map[string]interface{}, 10)}
	req = &QueryCampDailySumRequest{
		Request: &request,
	}
	return
}

func (req *QueryCampDailySumRequest) SetIsDaily(isDaily bool) {
	req.Request.Params["isDaily"] = isDaily
}

func (req *QueryCampDailySumRequest) GetIsDaily() bool {
	isDaily, found := req.Request.Params["isDaily"]
	if found {
		return isDaily.(bool)
	}
	return false
}

func (req *QueryCampDailySumRequest) SetCampaignId(campaignId int) {
	req.Request.Params["campaignId"] = campaignId
}

func (req *QueryCampDailySumRequest) GetCampaignId() int {
	campaignId, found := req.Request.Params["campaignId"]
	if found {
		return campaignId.(int)
	}
	return 0
}

func (req *QueryCampDailySumRequest) SetStartDay(startDay string) {
	req.Request.Params["startDay"] = startDay
}

func (req *QueryCampDailySumRequest) GetStartDay() string {
	startDay, found := req.Request.Params["startDay"]
	if found {
		return startDay.(string)
	}
	return ""
}

func (req *QueryCampDailySumRequest) SetEndDay(endDay string) {
	req.Request.Params["endDay"] = endDay
}

func (req *QueryCampDailySumRequest) GetEndDay() string {
	endDay, found := req.Request.Params["endDay"]
	if found {
		return endDay.(string)
	}
	return ""
}

func (req *QueryCampDailySumRequest) SetIsOrderOrClick(isOrderOrClick bool) {
	req.Request.Params["isOrderOrClick"] = isOrderOrClick
}

func (req *QueryCampDailySumRequest) GetIsOrderOrClick() bool {
	isOrderOrClick, found := req.Request.Params["isOrderOrClick"]
	if found {
		return isOrderOrClick.(bool)
	}
	return false
}

func (req *QueryCampDailySumRequest) SetIsTodayOr15Days(isTodayOr15Days bool) {
	req.Request.Params["isTodayOr15Days"] = isTodayOr15Days
}

func (req *QueryCampDailySumRequest) GetIsTodayOr15Days() bool {
	isTodayOr15Days, found := req.Request.Params["isTodayOr15Days"]
	if found {
		return isTodayOr15Days.(bool)
	}
	return false
}

func (req *QueryCampDailySumRequest) SetOrderStatusCategory(orderStatusCategory int) {
	req.Request.Params["orderStatusCategory"] = orderStatusCategory
}

func (req *QueryCampDailySumRequest) GetOrderStatusCategory() int {
	orderStatusCategory, found := req.Request.Params["orderStatusCategory"]
	if found {
		return orderStatusCategory.(int)
	}
	return 0
}

func (req *QueryCampDailySumRequest) SetPlatform(platform string) {
	req.Request.Params["platform"] = platform
}

func (req *QueryCampDailySumRequest) GetPlatform() string {
	platform, found := req.Request.Params["platform"]
	if found {
		return platform.(string)
	}
	return ""
}

func (req *QueryCampDailySumRequest) SetPageIndex(pageIndex int) {
	req.Request.Params["pageIndex"] = pageIndex
}

func (req *QueryCampDailySumRequest) GetPageIndex() int {
	pageIndex, found := req.Request.Params["pageIndex"]
	if found {
		return pageIndex.(int)
	}
	return 0
}

func (req *QueryCampDailySumRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *QueryCampDailySumRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(int)
	}
	return 0
}
