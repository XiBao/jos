package imgzone

import "github.com/XiBao/jos/sdk"

type CategoryQueryRequest struct {
	Request *sdk.Request
}

func NewCategoryQueryRequest() (req *CategoryQueryRequest) {
	request := sdk.Request{MethodName: "jingdong.imgzone.category.query", Params: make(map[string]interface{}, 3)}
	req = &CategoryQueryRequest{
		Request: &request,
	}
	return
}

func (req *CategoryQueryRequest) SetCateId(cateId int64) {
	req.Request.Params["cate_id"] = cateId
}

func (req *CategoryQueryRequest) GetCateId() int64 {
	cateId, found := req.Request.Params["cate_id"]
	if found {
		return cateId.(int64)
	}
	return 0
}

func (req *CategoryQueryRequest) SetCateName(cateName string) {
	req.Request.Params["cate_name"] = cateName
}

func (req *CategoryQueryRequest) GetCateName() string {
	cateName, found := req.Request.Params["cate_name"]
	if found {
		return cateName.(string)
	}
	return ""
}

func (req *CategoryQueryRequest) SetParentCateId(parentCatId int64) {
	req.Request.Params["parent_cate_id"] = parentCatId
}

func (req *CategoryQueryRequest) GetParentCateId() int64 {
	parentCateId, found := req.Request.Params["parent_cate_id"]
	if found {
		return parentCateId.(int64)
	}
	return 0
}
