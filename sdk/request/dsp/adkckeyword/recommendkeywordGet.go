package adkckeyword

import "github.com/XiBao/jos/sdk"

type RecommendkeywordGetRequest struct {
	Request *sdk.Request
}

func NewRecommendkeywordGetRequest() (req *RecommendkeywordGetRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.adkckeyword.recommendkeyword.get", Params: make(map[string]interface{}, 5)}
	req = &RecommendkeywordGetRequest{
		Request: &request,
	}
	return
}

func (req *RecommendkeywordGetRequest) SetSkuId(skuId uint64) {
	req.Request.Params["skuId"] = skuId
}

func (req *RecommendkeywordGetRequest) GetSkuId() uint64 {
	skuId, found := req.Request.Params["skuId"]
	if found {
		return skuId.(uint64)
	}
	return 0
}

func (req *RecommendkeywordGetRequest) SetSearchType(searchType uint8) {
	req.Request.Params["searchType"] = searchType
}

func (req *RecommendkeywordGetRequest) GetSearchType() uint8 {
	searchType, found := req.Request.Params["searchType"]
	if found {
		return searchType.(uint8)
	}
	return 0
}

func (req *RecommendkeywordGetRequest) SetOrder(order uint8) {
	req.Request.Params["order"] = order
}

func (req *RecommendkeywordGetRequest) GetOrder() uint8 {
	order, found := req.Request.Params["order"]
	if found {
		return order.(uint8)
	}
	return 0
}

func (req *RecommendkeywordGetRequest) SetSortType(sortType uint8) {
	req.Request.Params["sortType"] = sortType
}

func (req *RecommendkeywordGetRequest) GetSortType() uint8 {
	sortType, found := req.Request.Params["sortType"]
	if found {
		return sortType.(uint8)
	}
	return 0
}

func (req *RecommendkeywordGetRequest) SetKeyWordType(keyWordType uint8) {
	req.Request.Params["keyWordType"] = keyWordType
}

func (req *RecommendkeywordGetRequest) GetKeyWordType() uint8 {
	keyWordType, found := req.Request.Params["keyWordType"]
	if found {
		return keyWordType.(uint8)
	}
	return 0
}
