package eco

import (
	"github.com/XiBao/jos/sdk"
)

type BizStreamFetchRequest struct {
	Request *sdk.Request
}

// create new request
func NewBizStreamFetchRequest() (req *BizStreamFetchRequest) {
	request := sdk.Request{MethodName: "jingdong.eco.data.biz.stream.fetch", Params: make(map[string]interface{}, 10)}
	req = &BizStreamFetchRequest{
		Request: &request,
	}
	return
}

func (req *BizStreamFetchRequest) SetTimeMin(timeMin string) {
	req.Request.Params["timeMin"] = timeMin
}

func (req *BizStreamFetchRequest) GetTimeMin() string {
	timeMin, found := req.Request.Params["timeMin"]
	if found {
		return timeMin.(string)
	}
	return ""
}

func (req *BizStreamFetchRequest) SetTimeMax(timeMax string) {
	req.Request.Params["timeMax"] = timeMax
}

func (req *BizStreamFetchRequest) GetTimeMax() string {
	timeMax, found := req.Request.Params["timeMax"]
	if found {
		return timeMax.(string)
	}
	return ""
}

func (req *BizStreamFetchRequest) SetServiceId(serviceId string) {
	req.Request.Params["serviceId"] = serviceId
}

func (req *BizStreamFetchRequest) GetServiceId() string {
	serviceId, found := req.Request.Params["serviceId"]
	if found {
		return serviceId.(string)
	}
	return ""
}

func (req *BizStreamFetchRequest) SetTime(time string) {
	req.Request.Params["TIME"] = time
}

func (req *BizStreamFetchRequest) GetTime() string {
	time, found := req.Request.Params["TIME"]
	if found {
		return time.(string)
	}
	return ""
}

func (req *BizStreamFetchRequest) SetTimestamp(timestamp string) {
	req.Request.Params["TIMESTAMP"] = timestamp
}

func (req *BizStreamFetchRequest) GetTimestamp() string {
	timestamp, found := req.Request.Params["TIMESTAMP"]
	if found {
		return timestamp.(string)
	}
	return ""
}

func (req *BizStreamFetchRequest) SetAdProStat(adProStat string) {
	req.Request.Params["ADPROSTAT"] = adProStat
}

func (req *BizStreamFetchRequest) GetAdProStat() string {
	adProStat, found := req.Request.Params["ADPROSTAT"]
	if found {
		return adProStat.(string)
	}
	return ""
}

func (req *BizStreamFetchRequest) SetAdProId(adProId string) {
	req.Request.Params["ADPROID"] = adProId
}

func (req *BizStreamFetchRequest) GetAdProId() string {
	adProId, found := req.Request.Params["ADPROID"]
	if found {
		return adProId.(string)
	}
	return ""
}

func (req *BizStreamFetchRequest) SetSku(sku string) {
	req.Request.Params["SKU"] = sku
}

func (req *BizStreamFetchRequest) GetSku() string {
	sku, found := req.Request.Params["SKU"]
	if found {
		return sku.(string)
	}
	return ""
}

func (req *BizStreamFetchRequest) SetAdType(adType string) {
	req.Request.Params["ADTYPE"] = adType
}

func (req *BizStreamFetchRequest) GetAdType() string {
	adType, found := req.Request.Params["ADTYPE"]
	if found {
		return adType.(string)
	}
	return ""
}

func (req *BizStreamFetchRequest) SetActEffectId(actEffectId string) {
	req.Request.Params["ACTEFFECTID"] = actEffectId
}

func (req *BizStreamFetchRequest) GetActEffectId() string {
	actEffectId, found := req.Request.Params["ACTEFFECTID"]
	if found {
		return actEffectId.(string)
	}
	return ""
}

func (req *BizStreamFetchRequest) SetTicketType(ticketType string) {
	req.Request.Params["TICKETTYPE"] = ticketType
}

func (req *BizStreamFetchRequest) GetTicketType() string {
	ticketType, found := req.Request.Params["TICKETTYPE"]
	if found {
		return ticketType.(string)
	}
	return ""
}
