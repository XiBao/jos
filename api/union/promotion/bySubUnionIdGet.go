package promotion

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/union/promotion"
)

type UnionPromotionBySubUnionIdGetResponse struct {
	ErrorResp *api.ErrorResponnse             `json:"error_response,omitempty"`
	Data      *UnionPromotionCodeResponseData `json:"jd_union_open_promotion_bysubunionid_get_responce,omitempty"`
}

func (r *UnionPromotionBySubUnionIdGetResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.Result == ""
}

func (r *UnionPromotionBySubUnionIdGetResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

// 获取通用推广链接
func UnionPromotionBySubUnionIdGet(req *UnionPromotionCodeRequest) (*PromotionCodeResp, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewUnionPromotionBySubUnionIdRequest()
	codeReq := &promotion.PromotionCodeReq{
		MaterialId: req.MaterialId,
		SiteId:     req.SiteId,
		PositionId: req.PositionId,
		SubUnionId: req.SubUnionId,
		Ext1:       req.Ext1,
		Pid:        req.Pid,
		ChainType:  req.ChainType,
		CouponUrl:  req.CouponUrl,
	}
	r.SetPromotionCodeReq(codeReq)

	var response UnionPromotionBySubUnionIdGetResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}

	if response.Data == nil {
		return nil, errors.New("no data")
	}
	var ret UnionPromotioncodeResult
	if err := client.Logger().DecodeJSON([]byte(response.Data.Result), &ret); err != nil {
		return nil, err
	}
	if ret.IsError() {
		return nil, ret
	}

	var codeResp PromotionCodeResp
	if err := json.Unmarshal(ret.Data, &codeResp); err != nil {
		return nil, err
	}
	return &codeResp, nil
}
