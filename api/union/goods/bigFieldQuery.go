package goods

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/union/goods"
)

type BidFieldQueryRequest struct {
	api.BaseRequest
	SkuIds []uint64 `json:"skuIds"` // skuId集合
	Fields []string `json:"fields"` // 查询域集合，不填写则查询全部 ('categoryInfo','imageInfo','baseBigFieldInfo','bookBigFieldInfo','videoBigFieldInfo')
}

type BidFieldQueryResponse struct {
	ErrorResp *api.ErrorResponnse        `json:"error_response,omitempty"`
	Data      *BidFieldQueryResponseData `json:"jd_union_open_goods_bigfield_query_response,omitempty"`
}

type BidFieldQueryResponseData struct {
	Result string `json:"queryResult,omitempty"`
}

type BidFieldQueryResult struct {
	Code    int64               `json:"code,omitempty"`
	Message string              `json:"message,omitempty"`
	Data    []BidFieldQueryResp `json:"data,omitempty"`
}

type BidFieldQueryResp struct {
	SkuId    int64              `json:"skuId,omitempty"`             // skuId
	SkuName  string             `json:"skuName,omitempty"`           // 商品名称
	Category *CategoryInfo      `json:"categoryInfo,omitempty"`      // 目录信息
	Image    *ImageInfo         `json:"imageInfo,omitempty"`         // 图片信心
	Base     *BaseBigFieldInfo  `json:"baseBigFieldInfo,omitempty"`  // 基础大字段信息
	Book     *BookBigFieldInfo  `json:"bookBigFieldInfo,omitempty"`  // 图书大字段信息
	Video    *VideoBigFieldInfo `json:"videoBigFieldInfo,omitempty"` // 影音大字段信息
}

type CategoryInfo struct {
	Cid1     uint64 `json:"cid1"`     // 一级类目ID
	Cid1Name string `json:"cid1Name"` // 一级类目名称
	Cid2     uint64 `json:"cid2"`     // 二级类目ID
	Cid2Name string `json:"cid2Name"` // 二级类目名称
	Cid3     uint64 `json:"cid3"`     // 三级类目ID
	Cid3Name string `json:"cid3Name"` // 三级类目名称
}

type ImageInfo struct {
	List []UrlInfo `json:"imageList,omitempty"` // 图片合集
}

type UrlInfo struct {
	Info Url `json:"urlInfo"`
}

type Url struct {
	Url string `json:"url"` // 图片链接地址，第一个图片链接为主图链接
}

type BaseBigFieldInfo struct {
	Wdis     string  `json:"wdis,omitempty"`     // 商品介绍
	PropCode string  `json:"propCode,omitempty"` // 规格参数
	WareQD   float64 `json:"wareQD,omitempty"`   // 包装清单(仅自营商品)
}

type BookBigFieldInfo struct {
	Comments        string `json:"comments,omitempty"`        // 媒体评论
	Image           string `json:"image,omitempty"`           // 精彩文摘与插图(插图)
	ContentDesc     string `json:"contentDesc,omitempty"`     // 内容摘要(内容简介)
	RelatedProducts string `json:"relatedProducts,omitempty"` // 产品描述(相关商品)
	EditerDesc      string `json:"editerDesc,omitempty"`      // 编辑推荐
	Catalogue       string `json:"catalogue,omitempty"`       // 目录
	BookAbstract    string `json:"bookAbstract,omitempty"`    // 精彩摘要(精彩书摘)
	AuthorDesc      string `json:"authorDesc,omitempty"`      // 作者简介
	Introduction    string `json:"introduction,omitempty"`    // 前言(前言/序言)
	ProductFeatures string `json:"productFeatures,omitempty"` // 产品特色
}

type VideoBigFieldInfo struct {
	Comments            string `json:"comments,omitempty"`             // 媒体评论
	Image               string `json:"image,omitempty"`                // 精彩文摘与插图(插图)
	ContentDesc         string `json:"contentDesc,omitempty"`          // 内容摘要(内容简介)
	RelatedProducts     string `json:"relatedProducts,omitempty"`      // 产品描述(相关商品)
	EditerDesc          string `json:"editerDesc,omitempty"`           // 编辑推荐
	Catalogue           string `json:"catalogue,omitempty"`            // 目录
	BoxContents         string `json:"box_Contents,omitempty"`         // 包装清单
	MaterialDescription string `json:"material_Description,omitempty"` // 特殊说明
	Manual              string `json:"manual,omitempty"`               // 说明书
	ProductFeatures     string `json:"productFeatures,omitempty"`      // 产品特色
}

// 大字段商品查询接口
func BidFieldQuery(req *BidFieldQueryRequest) ([]BidFieldQueryResp, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := goods.NewBigFieldQueryRequest()
	goodsReq := &goods.BigFieldGoodsReq{
		SkuIds: req.SkuIds,
		Fields: req.Fields,
	}
	r.SetGoodsReq(goodsReq)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	var response BidFieldQueryResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.Data == nil {
		return nil, errors.New("no result")
	}

	var ret BidFieldQueryResult
	err = json.Unmarshal([]byte(response.Data.Result), &ret)
	if err != nil {
		return nil, err
	}

	if ret.Code != 200 {
		return nil, &api.ErrorResponnse{Code: strconv.FormatInt(ret.Code, 10), ZhDesc: ret.Message}
	}

	return ret.Data, nil
}
