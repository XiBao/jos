package asset

import (
	"github.com/XiBao/jos/sdk"
)

type AccountBalanceQueryRequest struct {
	Request *sdk.Request
}

// create new request
func NewAccountBalanceQueryRequest() (req *AccountBalanceQueryRequest) {
	request := sdk.Request{MethodName: "jingdong.asset.account.balance.query", Params: make(map[string]interface{}, 1)}
	req = &AccountBalanceQueryRequest{
		Request: &request,
	}
	return
}

func (req *AccountBalanceQueryRequest) SetTypeId(typeId uint8) {
	req.Request.Params["type_id"] = typeId
}

func (req *AccountBalanceQueryRequest) GetTypeId() uint8 {
	typeId, found := req.Request.Params["type_id"]
	if found {
		return typeId.(uint8)
	}
	return 0
}
