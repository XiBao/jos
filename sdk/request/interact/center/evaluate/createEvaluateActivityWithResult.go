package center

import (
	"github.com/XiBao/jos/sdk"
)

type CreateEvaluateActivityWithResultRequest struct {
	Request *sdk.Request
}

// create new request
func NewCreateEvaluateActivityWithResultRequest() (req *CreateEvaluateActivityWithResultRequest) {
	request := sdk.Request{MethodName: "jingdong.com.jd.interact.center.api.write.EvaluateActivityWriteService.createActivityWithResult", Params: make(map[string]interface{})}
	req = &CreateEvaluateActivityWithResultRequest{
		Request: &request,
	}
	return
}

func (req *CreateEvaluateActivityWithResultRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *CreateEvaluateActivityWithResultRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *CreateEvaluateActivityWithResultRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *CreateEvaluateActivityWithResultRequest) SetSkuName(skuName string) {
	req.Request.Params["skuName"] = skuName
}

func (req *CreateEvaluateActivityWithResultRequest) GetSkuName() string {
	skuName, found := req.Request.Params["skuName"]
	if found {
		return skuName.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetCate3rdName(cate3rdName string) {
	req.Request.Params["cate3rdName"] = cate3rdName
}

func (req *CreateEvaluateActivityWithResultRequest) GetCate3rdName() string {
	cate3rdName, found := req.Request.Params["cate3rdName"]
	if found {
		return cate3rdName.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetCate1stName(cate1stName string) {
	req.Request.Params["cate1stName"] = cate1stName
}

func (req *CreateEvaluateActivityWithResultRequest) GetCate1stName() string {
	cate1stName, found := req.Request.Params["cate1stName"]
	if found {
		return cate1stName.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetCate3rdCode(cate3rdCode string) {
	req.Request.Params["cate3rdCode"] = cate3rdCode
}

func (req *CreateEvaluateActivityWithResultRequest) GetCate3rdCode() string {
	cate3rdCode, found := req.Request.Params["cate3rdCode"]
	if found {
		return cate3rdCode.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetPrice(price string) {
	req.Request.Params["price"] = price
}

func (req *CreateEvaluateActivityWithResultRequest) GetPrice() string {
	price, found := req.Request.Params["price"]
	if found {
		return price.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetCate2ndName(cate2ndName string) {
	req.Request.Params["cate2ndName"] = cate2ndName
}

func (req *CreateEvaluateActivityWithResultRequest) GetCate2ndName() string {
	cate2ndName, found := req.Request.Params["cate2ndName"]
	if found {
		return cate2ndName.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetSkuId(skuId string) {
	req.Request.Params["skuId"] = skuId
}

func (req *CreateEvaluateActivityWithResultRequest) GetSkuId() string {
	skuId, found := req.Request.Params["skuId"]
	if found {
		return skuId.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetWareId(wareId string) {
	req.Request.Params["wareId"] = wareId
}

func (req *CreateEvaluateActivityWithResultRequest) GetWareId() string {
	wareId, found := req.Request.Params["wareId"]
	if found {
		return wareId.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetCate1stCode(cate1stCode string) {
	req.Request.Params["cate1stCode"] = cate1stCode
}

func (req *CreateEvaluateActivityWithResultRequest) GetCate1stCode() string {
	cate1stCode, found := req.Request.Params["cate1stCode"]
	if found {
		return cate1stCode.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetCate2ndCode(cate2ndCode string) {
	req.Request.Params["cate2ndCode"] = cate2ndCode
}

func (req *CreateEvaluateActivityWithResultRequest) GetCate2ndCode() string {
	cate2ndCode, found := req.Request.Params["cate2ndCode"]
	if found {
		return cate2ndCode.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetSupplierCode(supplierCode string) {
	req.Request.Params["supplierCode"] = supplierCode
}

func (req *CreateEvaluateActivityWithResultRequest) GetSupplierCode() string {
	supplierCode, found := req.Request.Params["supplierCode"]
	if found {
		return supplierCode.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetEndTime(endTime string) {
	req.Request.Params["endTime"] = endTime
}

func (req *CreateEvaluateActivityWithResultRequest) GetEndTime() string {
	endTime, found := req.Request.Params["endTime"]
	if found {
		return endTime.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetModifier(modifier string) {
	req.Request.Params["modifier"] = modifier
}

func (req *CreateEvaluateActivityWithResultRequest) GetModifier() string {
	modifier, found := req.Request.Params["modifier"]
	if found {
		return modifier.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetStartTime(startTime string) {
	req.Request.Params["startTime"] = startTime
}

func (req *CreateEvaluateActivityWithResultRequest) GetStartTime() string {
	startTime, found := req.Request.Params["startTime"]
	if found {
		return startTime.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetPictureRequirement(pictureRequirement uint) {
	req.Request.Params["pictureRequirement"] = pictureRequirement
}

func (req *CreateEvaluateActivityWithResultRequest) GetPictureRequirement() uint {
	pictureRequirement, found := req.Request.Params["pictureRequirement"]
	if found {
		return pictureRequirement.(uint)
	}
	return 0
}

func (req *CreateEvaluateActivityWithResultRequest) SetShopName(shopName string) {
	req.Request.Params["shopName"] = shopName
}

func (req *CreateEvaluateActivityWithResultRequest) GetShopName() string {
	shopName, found := req.Request.Params["shopName"]
	if found {
		return shopName.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetValidateDay(validateDay string) {
	req.Request.Params["validateDay"] = validateDay
}

func (req *CreateEvaluateActivityWithResultRequest) GetValidateDay() string {
	validateDay, found := req.Request.Params["validateDay"]
	if found {
		return validateDay.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetAssetItemId(assetItemId string) {
	req.Request.Params["assetItemId"] = assetItemId
}

func (req *CreateEvaluateActivityWithResultRequest) GetAssetItemId() string {
	assetItemId, found := req.Request.Params["assetItemId"]
	if found {
		return assetItemId.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetType(eType string) {
	req.Request.Params["type"] = eType
}

func (req *CreateEvaluateActivityWithResultRequest) GetType() string {
	eType, found := req.Request.Params["type"]
	if found {
		return eType.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetDiscount(discount string) {
	req.Request.Params["discount"] = discount
}

func (req *CreateEvaluateActivityWithResultRequest) GetDiscount() string {
	discount, found := req.Request.Params["discount"]
	if found {
		return discount.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetAwardType(awardType string) {
	req.Request.Params["awardType"] = awardType
}

func (req *CreateEvaluateActivityWithResultRequest) GetAwardType() string {
	awardType, found := req.Request.Params["awardType"]
	if found {
		return awardType.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetQuota(quota string) {
	req.Request.Params["quota"] = quota
}

func (req *CreateEvaluateActivityWithResultRequest) GetQuota() string {
	quota, found := req.Request.Params["quota"]
	if found {
		return quota.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetRulePrice(rulePrice string) {
	req.Request.Params["rulePrice"] = rulePrice
}

func (req *CreateEvaluateActivityWithResultRequest) GetRulePrice() string {
	rulePrice, found := req.Request.Params["rulePrice"]
	if found {
		return rulePrice.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetFloatRatio(floatRatio string) {
	req.Request.Params["floatRatio"] = floatRatio
}

func (req *CreateEvaluateActivityWithResultRequest) GetFloatRatio() string {
	floatRatio, found := req.Request.Params["floatRatio"]
	if found {
		return floatRatio.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetNums(nums string) {
	req.Request.Params["nums"] = nums
}

func (req *CreateEvaluateActivityWithResultRequest) GetNums() string {
	nums, found := req.Request.Params["nums"]
	if found {
		return nums.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetBatchKey(batchKey string) {
	req.Request.Params["batchKey"] = batchKey
}

func (req *CreateEvaluateActivityWithResultRequest) GetBatchKey() string {
	batchKey, found := req.Request.Params["batchKey"]
	if found {
		return batchKey.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetExpireType(expireType string) {
	req.Request.Params["expireType"] = expireType
}

func (req *CreateEvaluateActivityWithResultRequest) GetExpireType() string {
	expireType, found := req.Request.Params["expireType"]
	if found {
		return expireType.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *CreateEvaluateActivityWithResultRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *CreateEvaluateActivityWithResultRequest) SetVedioRequirement(vedioRequirement uint) {
	req.Request.Params["vedioRequirement"] = vedioRequirement
}

func (req *CreateEvaluateActivityWithResultRequest) GetVedioRequirement() uint {
	vedioRequirement, found := req.Request.Params["vedioRequirement"]
	if found {
		return vedioRequirement.(uint)
	}
	return 0
}

func (req *CreateEvaluateActivityWithResultRequest) SetWordRequirement(wordRequirement uint) {
	req.Request.Params["wordRequirement"] = wordRequirement
}

func (req *CreateEvaluateActivityWithResultRequest) GetWordRequirement() uint {
	wordRequirement, found := req.Request.Params["wordRequirement"]
	if found {
		return wordRequirement.(uint)
	}
	return 0
}
