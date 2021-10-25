package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type PromotionCodeReq struct {
	MaterialId string `json:"materialId"`
	SiteId     string `json:"siteId,omitempty"`
	PositionId uint64 `json:"positionId,omitempty"`
	SubUnionId string `json:"subUnionId,omitempty"`
	Ext1       string `json:"ext1,omitempty"`
	Pid        string `json:"pid,omitempty"`
	ChainType  uint   `json:"chainType,omitempty"`
	CouponUrl  string `json:"couponUrl,omitempty"`
}

type UnionPromotionCodeRequest struct {
	Request *sdk.Request
}

// create new request
func NewUnionPromotionCodeRequest() (req *UnionPromotionCodeRequest) {
	request := sdk.Request{MethodName: "jd.union.open.promotion.common.get", IsUnionGW: true, Params: make(map[string]interface{}, 1)}
	req = &UnionPromotionCodeRequest{
		Request: &request,
	}
	return
}

func (req *UnionPromotionCodeRequest) SetPromotionCodeReq(codeReq *PromotionCodeReq) {
	req.Request.Params["promotionCodeReq"] = codeReq
}

func (req *UnionPromotionCodeRequest) GetPromotionCodeReq() *PromotionCodeReq {
	codeReq, found := req.Request.Params["promotionCodeReq"]
	if found {
		return codeReq.(*PromotionCodeReq)
	}
	return nil
}
