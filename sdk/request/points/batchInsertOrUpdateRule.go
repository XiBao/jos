package points

import (
	"github.com/XiBao/jos/sdk"
)

type BatchInsertOrUpdateRuleRequest struct {
	Request *sdk.Request
}

func NewBatchInsertOrUpdateRuleRequest() (req *BatchInsertOrUpdateRuleRequest) {
	request := sdk.Request{MethodName: "jingdong.points.jos.batchInsertOrUpdateRule", Params: make(map[string]interface{}, 3)}
	req = &BatchInsertOrUpdateRuleRequest{
		Request: &request,
	}
	return
}

func (req *BatchInsertOrUpdateRuleRequest) SetMultiple(multiple float64) {
	req.Request.Params["multiple"] = multiple
}

func (req *BatchInsertOrUpdateRuleRequest) GetMultiple() float64 {
	multiple, found := req.Request.Params["multiple"]
	if found {
		return multiple.(float64)
	}
	return 0
}

func (req *BatchInsertOrUpdateRuleRequest) SetCreateTime(createTime string) {
	req.Request.Params["createTime"] = createTime
}

func (req *BatchInsertOrUpdateRuleRequest) GetCreateTime() string {
	createTime, found := req.Request.Params["createTime"]
	if found {
		return createTime.(string)
	}
	return ""
}

func (req *BatchInsertOrUpdateRuleRequest) SetModifyTime(modifyTime string) {
	req.Request.Params["modifyTime"] = modifyTime
}

func (req *BatchInsertOrUpdateRuleRequest) GetModifyTime() string {
	modifyTime, found := req.Request.Params["modifyTime"]
	if found {
		return modifyTime.(string)
	}
	return ""
}
