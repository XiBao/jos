package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type UnionPromotionBySubUnoinIdRequest struct {
	Request *sdk.Request
}

// create new request
func NewUnionPromotionBySubUnionIdRequest() (req *UnionPromotionBySubUnoinIdRequest) {
	request := sdk.Request{MethodName: "jd.union.open.promotion.bysubunionid.get", IsUnionGW: true, Params: make(map[string]interface{}, 1)}
	req = &UnionPromotionBySubUnoinIdRequest{
		Request: &request,
	}
	return
}

func (req *UnionPromotionBySubUnoinIdRequest) SetPromotionCodeReq(codeReq *PromotionCodeReq) {
	req.Request.Params["promotionCodeReq"] = codeReq
}

func (req *UnionPromotionBySubUnoinIdRequest) GetPromotionCodeReq() *PromotionCodeReq {
	codeReq, found := req.Request.Params["promotionCodeReq"]
	if found {
		return codeReq.(*PromotionCodeReq)
	}
	return nil
}
