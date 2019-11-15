package goods

import (
	"errors"
	"strconv"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/union/goods"
	"github.com/daviddengcn/ljson"
)

type GoodsLinkQueryRequest struct {
	api.BaseRequest
	Url        string `json:"url`         // 链接
	SubUnionId string `json:"subUnionId"` // 子联盟ID（需要联系运营开通权限才能拿到数据）
}

type GoodsLinkQueryResponse struct {
	ErrorResp *api.ErrorResponnse         `json:"error_response,omitempty"`
	Data      *GoodsLinkQueryResponseData `json:"jd_union_open_goods_link_query_response,omitempty"`
}

type GoodsLinkQueryResponseData struct {
	Result LinkQueryResult `json:"queryResult,omitempty"`
}

type LinkQueryResult struct {
	Code    int64           `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
	Data    []LinkGoodsResp `json:"data,omitempty"`
}

type LinkGoodsResp struct {
	SkuId     int64   `json:"skuId,omitempty"`     // skuId
	ProductId int64   `json:"productId,omitempty"` // productId
	Images    string  `json:"images,omitempty"`    // 图片集，逗号','分割，首张为主图
	SkuName   string  `json:"skuName,omitempty"`   // 商品名称
	Price     float64 `json:"price,omitempty"`     // 京东价，单位：元
	CosRatio  float64 `json:"cosRatio,omitempty"`  // 佣金比例，单位：%
	ShortUrl  string  `json:"shortUrl,omitempty"`  // 短链接
	ShopId    uint64  `json:"shopId,omitempty"`    // 店铺ID
	ShopName  string  `json:"shopName,omitempty"`  // 店铺名称
	Sales     uint    `json:"sales,omitempty"`     // 30天引单量
	IsSelf    string  `json:"isSelf,omitempty"`    // 是否自营，g：自营，p：pop
}

// 链接商品查询接口
func GoodsLinkQuery(req *PromotionGoodsInfoQueryRequest) ([]LinkGoodsResp, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := goods.NewGoodsLinkQueryRequest()
	goodsReq := &goods.LinkGoodsReq{
		Url:        req.Url,
		SubUnionId: req.SubUnionId,
	}
	r.SetGoodsReq(goodsReq)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	var response GoodsLinkQueryResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.Data == nil {
		return nil, errors.New("no result")
	}

	if response.Data.Result.Code != 200 {
		return nil, &api.ErrorResponnse{Code: strconv.FormatInt(response.Data.Result.Code, 10), ZhDesc: response.Data.Result.Message}
	}

	return response.Data.Result.Data, nil
}
