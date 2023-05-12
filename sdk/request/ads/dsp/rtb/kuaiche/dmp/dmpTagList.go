package dmp

import (
	"github.com/XiBao/jos/sdk"
)

type KuaicheDmpTagListRequest struct {
	Request *sdk.Request
}

// create new request
func NewKuaicheDmpTagListRequest() (req *KuaicheDmpTagListRequest) {
	request := sdk.Request{MethodName: "jingdong.dmp.new.tag.list.v1", Params: make(map[string]interface{}, 10)}
	req = &KuaicheDmpTagListRequest{
		Request: &request,
	}
	return
}

func (req *KuaicheDmpTagListRequest) SetAccessPin(accessPin string) {
	req.Request.Params["accessPin"] = accessPin
}

func (req *KuaicheDmpTagListRequest) GetAccessPin() string {
	accessPin, found := req.Request.Params["accessPin"]
	if found {
		return accessPin.(string)
	}
	return ""
}

func (req *KuaicheDmpTagListRequest) SetAuthType(authType string) {
	req.Request.Params["authType"] = authType
}

func (req *KuaicheDmpTagListRequest) GetAuthType() string {
	authType, found := req.Request.Params["authType"]
	if found {
		return authType.(string)
	}
	return ""
}

func (req *KuaicheDmpTagListRequest) SetCategoryId(categoryId string) {
	req.Request.Params["categoryId"] = categoryId
}

func (req *KuaicheDmpTagListRequest) GetCategoryId() string {
	categoryId, found := req.Request.Params["categoryId"]
	if found {
		return categoryId.(string)
	}
	return ""
}

func (req *KuaicheDmpTagListRequest) SetPageIndex(pageIndex int) {
	req.Request.Params["pageIndex"] = pageIndex
}

func (req *KuaicheDmpTagListRequest) GetPageIndex() int {
	pageIndex, found := req.Request.Params["pageIndex"]
	if found {
		return pageIndex.(int)
	}
	return 0
}

func (req *KuaicheDmpTagListRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *KuaicheDmpTagListRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(int)
	}
	return 0
}

func (req *KuaicheDmpTagListRequest) SetTagName(tagName string) {
	req.Request.Params["tagName"] = tagName
}

func (req *KuaicheDmpTagListRequest) GetTagName() string {
	tagName, found := req.Request.Params["tagName"]
	if found {
		return tagName.(string)
	}
	return ""
}

func (req *KuaicheDmpTagListRequest) SetSortType(sortType uint) {
	req.Request.Params["sortType"] = sortType
}

func (req *KuaicheDmpTagListRequest) GetSortType() uint {
	sortType, found := req.Request.Params["sortType"]
	if found {
		return sortType.(uint)
	}
	return 0
}

func (req *KuaicheDmpTagListRequest) SetIsFavorite(isFavorite int) {
	req.Request.Params["isFavorite"] = isFavorite
}

func (req *KuaicheDmpTagListRequest) GetIsFavorite() int {
	isFavorite, found := req.Request.Params["isFavorite"]
	if found {
		return isFavorite.(int)
	}
	return 0
}

func (req *KuaicheDmpTagListRequest) SetLevel(level uint) {
	req.Request.Params["level"] = level
}

func (req *KuaicheDmpTagListRequest) GetLevel() uint {
	level, found := req.Request.Params["level"]
	if found {
		return level.(uint)
	}
	return 0
}

func (req *KuaicheDmpTagListRequest) SetTagCategoryType(tagCategoryType uint) {
	req.Request.Params["tagCategoryType"] = tagCategoryType
}

func (req *KuaicheDmpTagListRequest) GetTagCategoryType() uint {
	tagCategoryType, found := req.Request.Params["tagCategoryType"]
	if found {
		return tagCategoryType.(uint)
	}
	return 0
}
