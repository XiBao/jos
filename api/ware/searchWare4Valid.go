package ware

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ware"
)

type SearchWare4ValidRequest struct {
	api.BaseRequest
	WareId               string   `json:"wareId,omitempty" codec:"wareId,omitempty"`                             // 商品id列表,最多20个
	SearchKey            string   `json:"searchKey,omitempty" codec:"searchKey,omitempty"`                       // 商品搜索关键词,需要配合搜索域实现
	SearchField          string   `json:"searchField,omitempty" codec:"searchField,omitempty"`                   // 商品搜索域的范围,默认是商品名称.目前值范围[title]
	CategoryId           uint64   `json:"categoryId,omitempty" codec:"categoryId,omitempty"`                     // 商品叶子类目
	ShopCategoryIdLevel1 uint64   `json:"shopCategoryIdLevel1,omitempty" codec:"shopCategoryIdLevel1,omitempty"` // 一级店内分类id
	ShopCategoryIdLevel2 uint64   `json:"shopCategoryIdLevel2,omitempty" codec:"shopCategoryIdLevel2,omitempty"` // 二级店内分类id
	TemplateId           uint64   `json:"templateId,omitempty" codec:"templateId,omitempty"`                     // 关联板式ID
	PromiseId            uint64   `json:"promiseId,omitempty" codec:"promiseId,omitempty"`                       // 时效模板ID
	BrandId              uint64   `json:"brandId,omitempty" codec:"brandId,omitempty"`                           // 品牌ID
	FeatureKey           []string `json:"featureKey,omitempty" codec:"featureKey,omitempty"`                     // 商品的特殊属性key
	FeatureValue         []string `json:"featureValue,omitempty" codec:"featureValue,omitempty"`                 // 商品的特殊属性value
	WareStatusValue      []int    `json:"wareStatusValue,omitempty" codec:"wareStatusValue,omitempty"`           // 商品状态,多个值属于[或]操作 1:从未上架 2:自主下架 4:系统下架 8:上架 513:从未上架待审 514:自主下架待审 516:系统下架待审 520:上架待审核 1028:系统下架审核失败
	ItemNum              string   `json:"itemNum,omitempty" codec:"itemNum,omitempty"`                           // 商品货号
	BarCode              string   `json:"barCode,omitempty" codec:"barCode,omitempty"`                           // 商品条码
	ColType              int      `json:"colType,omitempty" codec:"colType,omitempty"`                           // 合作类型
	StartCreatedTime     int64    `json:"startCreatedTime,omitempty" codec:"startCreatedTime,omitempty"`         // 开始创建时间
	EndCreatedTime       int64    `json:"endCreatedTime,omitempty" codec:"endCreatedTime,omitempty"`             // 结束创建时间
	StartJdPrice         float64  `json:"startJdPrice,omitempty" codec:"startJdPrice,omitempty"`                 // 	开始京东价
	EndJdPrice           float64  `json:"endJdPrice,omitempty" codec:"endJdPrice,omitempty"`                     // 结束京东价
	StartOnlineTime      int64    `json:"startOnlineTime,omitempty" codec:"startOnlineTime,omitempty"`           // 开始上架时间
	EndOnlineTime        int64    `json:"endOnlineTime,omitempty" codec:"endOnlineTime,omitempty"`               // 结束上架时间
	StartModifiedTime    int64    `json:"startModifiedTime,omitempty" codec:"startModifiedTime,omitempty"`       // 开始修改时间
	EndModifiedTime      int64    `json:"endModifiedTime,omitempty" codec:"endModifiedTime,omitempty"`           // 结束修改时间
	StartOfflineTime     int64    `json:"startOfflineTime,omitempty" codec:"startOfflineTime,omitempty"`         // 开始下架时间
	EndOfflineTime       int64    `json:"endOfflineTime,omitempty" codec:"endOfflineTime,omitempty"`             // 结束下架时间
	StartStockNum        int      `json:"startStockNum,omitempty" codec:"startStockNum,omitempty"`               // 开始商品库存
	EndStockNum          int      `json:"endStockNum,omitempty" codec:"endStockNum,omitempty"`                   // 结束商品库存
	OrderField           []string `json:"orderField,omitempty" codec:"orderField,omitempty"`                     // 排序字段.值范围[wareId offlineTime onlineTime stockNum jdPrice modified]
	OrderType            []string `json:"orderType,omitempty" codec:"orderType,omitempty"`                       // 排序方式.值范围[asc desc]
	PageNo               int      `json:"pageNo,omitempty" codec:"pageNo,omitempty"`                             // 页码
	PageSize             int      `json:"pageSize,omitempty" codec:"pageSize,omitempty"`                         // 每页条数
	TransportId          uint64   `json:"transportId,omitempty" codec:"transportId,omitempty"`                   // 运费模板ID
	Claim                int      `json:"claim,omitempty" codec:"claim,omitempty"`                               // 是否认领
	GroupId              uint64   `json:"groupId,omitempty" codec:"groupId,omitempty"`                           // 分组ID
	MultiCategoryId      uint64   `json:"multiCategoryId,omitempty" codec:"multiCategoryId,omitempty"`           // 四级类目ID
	WarePropKey          string   `json:"warePropKey,omitempty" codec:"warePropKey,omitempty"`                   // 商品的类目属性搜索key
	WarePropValue        string   `json:"warePropValue,omitempty" codec:"warePropValue,omitempty"`               // 商品的类目属性搜索value
	Field                string   `json:"field,omitempty" codec:"field,omitempty"`                               // 可选的返回的字段
}

type SearchWare4ValidResponse struct {
	ErrorResp *api.ErrorResponnse          `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *SearchWare4ValidSubResponse `json:"jingdong_ware_read_searchWare4Valid_responce,omitempty" codec:"jingdong_ware_read_searchWare4Valid_responce,omitempty"`
}

type SearchWare4ValidSubResponse struct {
	Code      string                `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Page      *SearchWare4ValidPage `json:"page,omitempty" codec:"page,omitempty"`
}

type SearchWare4ValidPage struct {
	Data      []*Ware `json:"data,omitempty" codec:"data,omitempty"`
	PageNo    int     `json:"pageNo,omitempty" codec:"pageNo,omitempty"`
	PageSize  int     `json:"pageSize,omitempty" codec:"pageSize,omitempty"`
	TotalItem int     `json:"totalItem,omitempty" codec:"totalItem,omitempty"`
}

// 根据条件检索订单信息 （仅适用于SOP、LBP，SOPL类型，FBP类型请调取FBP订单检索 ）
func SearchWare4Valid(req *SearchWare4ValidRequest) (*SearchWare4ValidPage, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := ware.NewSearchWare4ValidRequest()
	if req.WareId != "" {
		r.SetWareId(req.WareId)
	}
	if req.Field != "" {
		r.SetField(req.Field)
	}
	if len(req.WareStatusValue) > 0 {
		r.SetWareStatusValue(req.WareStatusValue)
	}
	if req.PageNo > 0 {
		r.SetPageNo(req.PageNo)
	}
	if req.PageSize > 0 {
		r.SetPageSize(req.PageSize)
	}

	if req.SearchKey != "" {
		r.SetSearchKey(req.SearchKey)
	}
	if req.SearchField != "" {
		r.SetSearchField(req.SearchField)
	}
	if len(req.OrderField) > 0 {
		r.SetOrderField(req.OrderField)
	}
	if len(req.OrderType) > 0 {
		r.SetOrderType(req.OrderType)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No ware info.")
	}

	var response SearchWare4ValidResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}

	if len(response.Data.Page.Data) > 0 {
		for _, v := range response.Data.Page.Data {
			if v.OfflineTime > 0 {
				v.OfflineTimeStr = time.Unix(v.OfflineTime/1000, v.OfflineTime%1000).Format("2006-01-02 15:04:05")
			}
			if v.OnlineTime > 0 {
				v.OnlineTimeStr = time.Unix(v.OnlineTime/1000, v.OnlineTime%1000).Format("2006-01-02 15:04:05")
			}
		}
	}
	return response.Data.Page, nil
}
