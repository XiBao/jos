package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionAddRequest struct {
	Request *sdk.Request
}

// add new request
func NewSellerPromotionAddRequest() (req *SellerPromotionAddRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.add", Params: make(map[string]interface{}, 9)}
	req = &SellerPromotionAddRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionAddRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *SellerPromotionAddRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *SellerPromotionAddRequest) SetType(pType uint8) {
	req.Request.Params["type"] = pType
}

func (req *SellerPromotionAddRequest) GetType() string {
	pType, found := req.Request.Params["type"]
	if found {
		return pType.(string)
	}
	return ""
}

func (req *SellerPromotionAddRequest) SetBeginTime(beginTime string) {
	req.Request.Params["begin_time"] = beginTime
}

func (req *SellerPromotionAddRequest) GetBeginTime() string {
	beginTime, found := req.Request.Params["begin_time"]
	if found {
		return beginTime.(string)
	}
	return ""
}

func (req *SellerPromotionAddRequest) SetEndTime(endTime string) {
	req.Request.Params["end_time"] = endTime
}

func (req *SellerPromotionAddRequest) GetEndTime() string {
	endTime, found := req.Request.Params["end_time"]
	if found {
		return endTime.(string)
	}
	return ""
}

func (req *SellerPromotionAddRequest) SetBound(bound uint8) {
	req.Request.Params["bound"] = bound
}

func (req *SellerPromotionAddRequest) GetBound() uint8 {
	bound, found := req.Request.Params["bound"]
	if found {
		return bound.(uint8)
	}
	return 0
}

func (req *SellerPromotionAddRequest) SetMember(member uint8) {
	req.Request.Params["member"] = member
}

func (req *SellerPromotionAddRequest) GetMember() uint8 {
	member, found := req.Request.Params["member"]
	if found {
		return member.(uint8)
	}
	return 0
}

func (req *SellerPromotionAddRequest) SetSlogan(slogan string) {
	req.Request.Params["slogan"] = slogan
}

func (req *SellerPromotionAddRequest) GetSlogan() string {
	slogan, found := req.Request.Params["slogan"]
	if found {
		return slogan.(string)
	}
	return ""
}

func (req *SellerPromotionAddRequest) SetComment(comment string) {
	req.Request.Params["comment"] = comment
}

func (req *SellerPromotionAddRequest) GetComment() string {
	comment, found := req.Request.Params["comment"]
	if found {
		return comment.(string)
	}
	return ""
}

func (req *SellerPromotionAddRequest) SetFavorMode(favorMode uint8) {
	req.Request.Params["favor_mode"] = favorMode
}

func (req *SellerPromotionAddRequest) GetFavorMode() uint8 {
	favorMode, found := req.Request.Params["favor_mode"]
	if found {
		return favorMode.(uint8)
	}
	return 0
}
