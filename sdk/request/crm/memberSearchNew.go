package crm

import (
	"github.com/XiBao/jos/sdk"
)

type MemberSearchNewRequest struct {
	Request *sdk.Request
}

// create new request
func NewMemberSearchNewRequest() (req *MemberSearchNewRequest) {
	request := sdk.Request{MethodName: "jingdong.crm.member.searchNew", Params: make(map[string]interface{}, 10)}
	req = &MemberSearchNewRequest{
		Request: &request,
	}
	return
}

func (req *MemberSearchNewRequest) SetPin(pin string) {
	req.Request.Params["customer_pin"] = pin
}

func (req *MemberSearchNewRequest) GetPin() string {
	pin, found := req.Request.Params["customer_pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *MemberSearchNewRequest) SetGrade(grade string) {
	req.Request.Params["grade"] = grade
}

func (req *MemberSearchNewRequest) GetGrade() string {
	grade, found := req.Request.Params["grade"]
	if found {
		return grade.(string)
	}
	return ""
}

func (req *MemberSearchNewRequest) SetMinLastTradeTime(minLastTradeTime string) {
	req.Request.Params["min_last_trade_time"] = minLastTradeTime
}

func (req *MemberSearchNewRequest) GetMinLastTradeTime() string {
	minLastTradeTime, found := req.Request.Params["min_last_trade_time"]
	if found {
		return minLastTradeTime.(string)
	}
	return ""
}

func (req *MemberSearchNewRequest) SetMaxLastTradeTime(maxLastTradeTime string) {
	req.Request.Params["max_last_trade_time"] = maxLastTradeTime
}

func (req *MemberSearchNewRequest) GetMaxLastTradeTime() string {
	maxLastTradeTime, found := req.Request.Params["max_last_trade_time"]
	if found {
		return maxLastTradeTime.(string)
	}
	return ""
}

func (req *MemberSearchNewRequest) SetMinTradeCount(minTradeCount int) {
	req.Request.Params["min_trade_count"] = minTradeCount
}

func (req *MemberSearchNewRequest) GetMinTradeCount() int {
	minTradeCount, found := req.Request.Params["min_trade_count"]
	if found {
		return minTradeCount.(int)
	}
	return 0
}

func (req *MemberSearchNewRequest) SetMaxTradeCount(maxTradeCount int) {
	req.Request.Params["max_trade_count"] = maxTradeCount
}

func (req *MemberSearchNewRequest) GetMaxTradeCount() int {
	maxTradeCount, found := req.Request.Params["max_trade_count"]
	if found {
		return maxTradeCount.(int)
	}
	return 0
}

func (req *MemberSearchNewRequest) SetAvgPrice(avgPrice int) {
	req.Request.Params["avg_price"] = avgPrice
}

func (req *MemberSearchNewRequest) GetAvgPrice() int {
	avgPrice, found := req.Request.Params["avg_price"]
	if found {
		return avgPrice.(int)
	}
	return 0
}

func (req *MemberSearchNewRequest) SetMinTradeAmount(minTradeAmount int) {
	req.Request.Params["min_trade_amount"] = minTradeAmount
}

func (req *MemberSearchNewRequest) GetMinTradeAmount() int {
	minTradeAmount, found := req.Request.Params["min_trade_amount"]
	if found {
		return minTradeAmount.(int)
	}
	return 0
}

func (req *MemberSearchNewRequest) SetPage(page int) {
	req.Request.Params["current_page"] = page
}

func (req *MemberSearchNewRequest) GetPage() int {
	page, found := req.Request.Params["current_page"]
	if found {
		return page.(int)
	}
	return 0
}

func (req *MemberSearchNewRequest) SetPageSize(pageSize int) {
	req.Request.Params["page_size"] = pageSize
}

func (req *MemberSearchNewRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(int)
	}
	return 0
}
