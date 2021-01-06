package crm

import (
	"github.com/XiBao/jos/sdk"
)

type ShopGiftCreateRequest struct {
	Request *sdk.Request
}

// create new request
func NewShopGiftCreateRequest() (req *ShopGiftCreateRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.crm.shopGift.createShopGift", Params: make(map[string]interface{}, 19)}
	req = &ShopGiftCreateRequest{
		Request: &request,
	}
	return
}

func (req *ShopGiftCreateRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *ShopGiftCreateRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetStartDate(startDate string) {
	req.Request.Params["startDate"] = startDate
}

func (req *ShopGiftCreateRequest) GetStartDate() string {
	startDate, found := req.Request.Params["startDate"]
	if found {
		return startDate.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetEndDate(endDate string) {
	req.Request.Params["endDate"] = endDate
}

func (req *ShopGiftCreateRequest) GetEndDate() string {
	endDate, found := req.Request.Params["endDate"]
	if found {
		return endDate.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetModelIds(modelIds string) {
	req.Request.Params["modelIds"] = modelIds
}

func (req *ShopGiftCreateRequest) GetModelIds() string {
	modelIds, found := req.Request.Params["modelIds"]
	if found {
		return modelIds.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetCreator(creator string) {
	req.Request.Params["creator"] = creator
}

func (req *ShopGiftCreateRequest) GetCreator() string {
	creator, found := req.Request.Params["creator"]
	if found {
		return creator.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetCreateChannel(createChannel string) {
	req.Request.Params["createChannel"] = createChannel
}

func (req *ShopGiftCreateRequest) GetCreateChannel() string {
	createChannel, found := req.Request.Params["createChannel"]
	if found {
		return createChannel.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetCloseLink(closeLink string) {
	req.Request.Params["closeLink"] = closeLink
}

func (req *ShopGiftCreateRequest) GetCloseLink() string {
	closeLink, found := req.Request.Params["closeLink"]
	if found {
		return closeLink.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetIsvName(isvName string) {
	req.Request.Params["isvName"] = isvName
}

func (req *ShopGiftCreateRequest) GetIsvName() string {
	isvName, found := req.Request.Params["isvName"]
	if found {
		return isvName.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetActivityLink(activityLink string) {
	req.Request.Params["activityLink"] = activityLink
}

func (req *ShopGiftCreateRequest) GetActivityLink() string {
	activityLink, found := req.Request.Params["activityLink"]
	if found {
		return activityLink.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetRequestPin(requestPin string) {
	req.Request.Params["requestPin"] = requestPin
}

func (req *ShopGiftCreateRequest) GetRequestPin() string {
	requestPin, found := req.Request.Params["requestPin"]
	if found {
		return requestPin.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetIsvValidate(isvValidate string) {
	req.Request.Params["isvValidate"] = isvValidate
}

func (req *ShopGiftCreateRequest) GetIsvValidate() string {
	isvValidate, found := req.Request.Params["isvValidate"]
	if found {
		return isvValidate.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetSendNumber(sendNumber string) {
	req.Request.Params["sendNumber"] = sendNumber
}

func (req *ShopGiftCreateRequest) GetSendNumber() string {
	sendNumber, found := req.Request.Params["sendNumber"]
	if found {
		return sendNumber.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetOpenIdSeller(openIdSeller string) {
	req.Request.Params["open_id_seller"] = openIdSeller
}

func (req *ShopGiftCreateRequest) GetOpenIdSeller() string {
	openIdSeller, found := req.Request.Params["open_id_seller"]
	if found {
		return openIdSeller.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetOpenIdIsv(openIdIsv string) {
	req.Request.Params["open_id_isv"] = openIdIsv
}

func (req *ShopGiftCreateRequest) GetOpenIdIsv() string {
	openIdIsv, found := req.Request.Params["open_id_isv"]
	if found {
		return openIdIsv.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetDiscount(discount string) {
	req.Request.Params["discount"] = discount
}

func (req *ShopGiftCreateRequest) GetDiscount() string {
	discount, found := req.Request.Params["discount"]
	if found {
		return discount.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetQuota(quota string) {
	req.Request.Params["quota"] = quota
}

func (req *ShopGiftCreateRequest) GetQuota() string {
	quota, found := req.Request.Params["quota"]
	if found {
		return quota.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetValidateDay(validateDay string) {
	req.Request.Params["validateDay"] = validateDay
}

func (req *ShopGiftCreateRequest) GetValidateDay() string {
	validateDay, found := req.Request.Params["validateDay"]
	if found {
		return validateDay.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetPrizeType(prizeType string) {
	req.Request.Params["prizeType"] = prizeType
}

func (req *ShopGiftCreateRequest) GetPrizeType() string {
	prizeType, found := req.Request.Params["prizeType"]
	if found {
		return prizeType.(string)
	}
	return ""
}

func (req *ShopGiftCreateRequest) SetSendCount(sendCount string) {
	req.Request.Params["sendCount"] = sendCount
}

func (req *ShopGiftCreateRequest) GetSendCount() string {
	sendCount, found := req.Request.Params["sendCount"]
	if found {
		return sendCount.(string)
	}
	return ""
}
