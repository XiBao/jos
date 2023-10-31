package goods

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/union/goods"
)

type GoodsLinkQueryRequest struct {
	api.BaseRequest
	Url        string `json:"url"`        // 链接
	SubUnionId string `json:"subUnionId"` // 子联盟ID（需要联系运营开通权限才能拿到数据）
}

type GoodsLinkQueryResponse struct {
	ErrorResp *api.ErrorResponnse         `json:"error_response,omitempty"`
	Data      *GoodsLinkQueryResponseData `json:"jd_union_open_goods_link_query_response,omitempty"`
}

func (r GoodsLinkQueryResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.Result.IsError()
}

func (r GoodsLinkQueryResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Result.Error()
	}
	return "no result data"
}

type GoodsLinkQueryResponseData struct {
	Result LinkQueryResult `json:"queryResult,omitempty"`
}

type LinkQueryResult struct {
	Code    int64           `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
	Data    []LinkGoodsResp `json:"data,omitempty"`
}

func (r LinkQueryResult) IsError() bool {
	return r.Code != 200
}

func (r LinkQueryResult) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", r.Code, r.Message)
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
func GoodsLinkQuery(req *GoodsLinkQueryRequest) ([]LinkGoodsResp, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := goods.NewGoodsLinkQueryRequest()
	goodsReq := &goods.LinkGoodsReq{
		Url:        req.Url,
		SubUnionId: req.SubUnionId,
	}
	r.SetGoodsReq(goodsReq)

	var response GoodsLinkQueryResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.Data, nil
}
