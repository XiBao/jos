package crm

import (
	"github.com/XiBao/jos/sdk"
)

type MemberScanRequest struct {
	Request *sdk.Request
}

// create new request
func NewMemberScanRequest() (req *MemberScanRequest) {
	request := sdk.Request{MethodName: "jingdong.crm.member.scan", Params: make(map[string]interface{}, 10)}
	req = &MemberScanRequest{
		Request: &request,
	}
	return
}

func (req *MemberScanRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *MemberScanRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *MemberScanRequest) SetGrade(grade string) {
	req.Request.Params["grade"] = grade
}

func (req *MemberScanRequest) GetGrade() string {
	grade, found := req.Request.Params["grade"]
	if found {
		return grade.(string)
	}
	return ""
}

func (req *MemberScanRequest) SetMinLastTradeTime(minLastTradeTime string) {
	req.Request.Params["min_last_trade_time"] = minLastTradeTime
}

func (req *MemberScanRequest) GetMinLastTradeTime() string {
	minLastTradeTime, found := req.Request.Params["min_last_trade_time"]
	if found {
		return minLastTradeTime.(string)
	}
	return ""
}

func (req *MemberScanRequest) SetMaxLastTradeTime(maxLastTradeTime string) {
	req.Request.Params["max_last_trade_time"] = maxLastTradeTime
}

func (req *MemberScanRequest) GetMaxLastTradeTime() string {
	maxLastTradeTime, found := req.Request.Params["max_last_trade_time"]
	if found {
		return maxLastTradeTime.(string)
	}
	return ""
}

func (req *MemberScanRequest) SetMinTradeCount(minTradeCount int) {
	req.Request.Params["min_trade_count"] = minTradeCount
}

func (req *MemberScanRequest) GetMinTradeCount() int {
	minTradeCount, found := req.Request.Params["min_trade_count"]
	if found {
		return minTradeCount.(int)
	}
	return 0
}

func (req *MemberScanRequest) SetMaxTradeCount(maxTradeCount int) {
	req.Request.Params["max_trade_count"] = maxTradeCount
}

func (req *MemberScanRequest) GetMaxTradeCount() int {
	maxTradeCount, found := req.Request.Params["max_trade_count"]
	if found {
		return maxTradeCount.(int)
	}
	return 0
}

func (req *MemberScanRequest) SetAvgPrice(avgPrice int) {
	req.Request.Params["avg_price"] = avgPrice
}

func (req *MemberScanRequest) GetAvgPrice() int {
	avgPrice, found := req.Request.Params["avg_price"]
	if found {
		return avgPrice.(int)
	}
	return 0
}

func (req *MemberScanRequest) SetMinTradeAmount(minTradeAmount int) {
	req.Request.Params["min_trade_amount"] = minTradeAmount
}

func (req *MemberScanRequest) GetMinTradeAmount() int {
	minTradeAmount, found := req.Request.Params["min_trade_amount"]
	if found {
		return minTradeAmount.(int)
	}
	return 0
}

func (req *MemberScanRequest) SetPageSize(pageSize int) {
	req.Request.Params["page_size"] = pageSize
}

func (req *MemberScanRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(int)
	}
	return 0
}

func (req *MemberScanRequest) SetScrollId(scrollId string) {
	req.Request.Params["scroll_id"] = scrollId
}

func (req *MemberScanRequest) GetScrollId() string {
	scrollId, found := req.Request.Params["scroll_id"]
	if found {
		return scrollId.(string)
	}
	return ""
}
