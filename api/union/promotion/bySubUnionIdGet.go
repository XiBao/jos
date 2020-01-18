package promotion

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/union/promotion"
	"github.com/daviddengcn/ljson"
)

// 获取通用推广链接
func UnionPromotionBySubUnionIdGet(req *UnionPromotionCodeRequest) (string, string, error) {
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return "", err
	}
	var response UnionPromotionCodeResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return "", err
	}

	if response.Data == nil {
		return "", errors.New("no data")
	}
	var ret UnionPromotioncodeResult
	err = ljson.Unmarshal([]byte(response.Data.Result), &ret)
	if err != nil {
		return "", err
	}

	if ret.Code != 200 {
		return "", &api.ErrorResponnse{Code: strconv.FormatInt(int64(ret.Code), 10), ZhDesc: ret.Message}
	}

	var codeResp PromotionCodeResp
	err = json.Unmarshal(ret.Data, &codeResp)
	if err != nil {
		return "", err
	}
	return codeResp.ClickURL, codeResp.ShortURL, nil
}
