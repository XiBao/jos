package report

import (
	"github.com/XiBao/jos/sdk"
)

type DspReportQueryAccountReportRequest struct {
	Request *sdk.Request
}

// create new request
func NewDspReportQueryAccountReportRequest() (req *DspReportQueryAccountReportRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.report.queryAccountReport", Params: make(map[string]interface{}, 9)}
	req = &DspReportQueryAccountReportRequest{
		Request: &request,
	}
	return
}

func (req *DspReportQueryAccountReportRequest) SetStartDay(startDay string) {
	req.Request.Params["startDay"] = startDay
}

func (req *DspReportQueryAccountReportRequest) GetStartDay() string {
	startDay, found := req.Request.Params["startDay"]
	if found {
		return startDay.(string)
	}
	return ""
}

func (req *DspReportQueryAccountReportRequest) SetEndDay(endDay string) {
	req.Request.Params["endDay"] = endDay
}

func (req *DspReportQueryAccountReportRequest) GetEndDay() string {
	endDay, found := req.Request.Params["endDay"]
	if found {
		return endDay.(string)
	}
	return ""
}

func (req *DspReportQueryAccountReportRequest) SetIsDaily(isDaily string) {
	req.Request.Params["isDaily"] = isDaily
}

func (req *DspReportQueryAccountReportRequest) GetIsDaily() string {
	isDaily, found := req.Request.Params["isDaily"]
	if found {
		return isDaily.(string)
	}
	return ""
}

func (req *DspReportQueryAccountReportRequest) SetPlatform(platform string) {
	req.Request.Params["platform"] = platform
}

func (req *DspReportQueryAccountReportRequest) GetPlatform() string {
	platform, found := req.Request.Params["platform"]
	if found {
		return platform.(string)
	}
	return ""
}

func (req *DspReportQueryAccountReportRequest) SetClickOrOrderDay(clickOrOrderDay uint8) {
	req.Request.Params["clickOrOrderDay"] = clickOrOrderDay
}

func (req *DspReportQueryAccountReportRequest) GetClickOrOrderDay() uint8 {
	clickOrOrderDay, found := req.Request.Params["clickOrOrderDay"]
	if found {
		return clickOrOrderDay.(uint8)
	}
	return 0
}

func (req *DspReportQueryAccountReportRequest) SetClickOrOrderCaliber(clickOrOrderCaliber uint8) {
	req.Request.Params["clickOrOrderCaliber"] = clickOrOrderCaliber
}

func (req *DspReportQueryAccountReportRequest) GetClickOrOrderCaliber() uint8 {
	clickOrOrderCaliber, found := req.Request.Params["clickOrOrderCaliber"]
	if found {
		return clickOrOrderCaliber.(uint8)
	}
	return 0
}

func (req *DspReportQueryAccountReportRequest) SetOrderStatusCategory(orderStatusCategory uint8) {
	req.Request.Params["orderStatusCategory"] = orderStatusCategory
}

func (req *DspReportQueryAccountReportRequest) GetOrderStatusCategory() uint8 {
	orderStatusCategory, found := req.Request.Params["orderStatusCategory"]
	if found {
		return orderStatusCategory.(uint8)
	}
	return 0
}

func (req *DspReportQueryAccountReportRequest) SetPageIndex(pageIndex uint16) {
	req.Request.Params["pageIndex"] = pageIndex
}

func (req *DspReportQueryAccountReportRequest) GetPageIndex() uint16 {
	pageIndex, found := req.Request.Params["pageIndex"]
	if found {
		return pageIndex.(uint16)
	}
	return 0
}

func (req *DspReportQueryAccountReportRequest) SetPageSize(pageSize uint8) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *DspReportQueryAccountReportRequest) GetPageSize() uint8 {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(uint8)
	}
	return 0
}
