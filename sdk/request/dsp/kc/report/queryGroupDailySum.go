package report

import "github.com/XiBao/jos/sdk"

type QueryGroupDailySumRequest struct {
	Request *sdk.Request
}

func NewQueryGroupDailySumRequest() (req *QueryGroupDailySumRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.report.queryGroupDailySum", Params: make(map[string]interface{}, 10)}
	req = &QueryGroupDailySumRequest{
		Request: &request,
	}
	return
}

func (req *QueryGroupDailySumRequest) SetIsDaily(isDaily bool) {
	req.Request.Params["isDaily"] = isDaily
}

func (req *QueryGroupDailySumRequest) GetIsDaily() bool {
	isDaily, found := req.Request.Params["isDaily"]
	if found {
		return isDaily.(bool)
	}
	return false
}

func (req *QueryGroupDailySumRequest) SetCampaignId(campaignId int) {
	req.Request.Params["campaignId"] = campaignId
}

func (req *QueryGroupDailySumRequest) GetCampaignId() int {
	campaignId, found := req.Request.Params["campaignId"]
	if found {
		return campaignId.(int)
	}
	return 0
}

func (req *QueryGroupDailySumRequest) SetStartDay(startDay string) {
	req.Request.Params["startDay"] = startDay
}

func (req *QueryGroupDailySumRequest) GetStartDay() string {
	startDay, found := req.Request.Params["startDay"]
	if found {
		return startDay.(string)
	}
	return ""
}

func (req *QueryGroupDailySumRequest) SetEndDay(endDay string) {
	req.Request.Params["endDay"] = endDay
}

func (req *QueryGroupDailySumRequest) GetEndDay() string {
	endDay, found := req.Request.Params["endDay"]
	if found {
		return endDay.(string)
	}
	return ""
}

func (req *QueryGroupDailySumRequest) SetIsOrderOrClick(isOrderOrClick bool) {
	req.Request.Params["isOrderOrClick"] = isOrderOrClick
}

func (req *QueryGroupDailySumRequest) GetIsOrderOrClick() bool {
	isOrderOrClick, found := req.Request.Params["isOrderOrClick"]
	if found {
		return isOrderOrClick.(bool)
	}
	return false
}

func (req *QueryGroupDailySumRequest) SetIsTodayOr15Days(isTodayOr15Days bool) {
	req.Request.Params["isTodayOr15Days"] = isTodayOr15Days
}

func (req *QueryGroupDailySumRequest) GetIsTodayOr15Days() bool {
	isTodayOr15Days, found := req.Request.Params["isTodayOr15Days"]
	if found {
		return isTodayOr15Days.(bool)
	}
	return false
}

func (req *QueryGroupDailySumRequest) SetOrderStatusCategory(orderStatusCategory int) {
	req.Request.Params["orderStatusCategory"] = orderStatusCategory
}

func (req *QueryGroupDailySumRequest) GetOrderStatusCategory() int {
	orderStatusCategory, found := req.Request.Params["orderStatusCategory"]
	if found {
		return orderStatusCategory.(int)
	}
	return 0
}

func (req *QueryGroupDailySumRequest) SetPlatform(platform string) {
	req.Request.Params["platform"] = platform
}

func (req *QueryGroupDailySumRequest) GetPlatform() string {
	platform, found := req.Request.Params["platform"]
	if found {
		return platform.(string)
	}
	return ""
}

func (req *QueryGroupDailySumRequest) SetPageIndex(pageIndex int) {
	req.Request.Params["pageIndex"] = pageIndex
}

func (req *QueryGroupDailySumRequest) GetPageIndex() int {
	pageIndex, found := req.Request.Params["pageIndex"]
	if found {
		return pageIndex.(int)
	}
	return 0
}

func (req *QueryGroupDailySumRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *QueryGroupDailySumRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(int)
	}
	return 0
}
