package sku

import (
	"github.com/XiBao/jos/sdk"
)

type GetUserCartSkuRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetUserCartSkuRequest() (req *GetUserCartSkuRequest) {
	request := sdk.Request{MethodName: "jingdong.market.bdp.userBehavior.getUserCartSku", Params: make(map[string]interface{}, 3)}
	req = &GetUserCartSkuRequest{
		Request: &request,
	}
	return
}

func (req *GetUserCartSkuRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *GetUserCartSkuRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *GetUserCartSkuRequest) SetOpenIdBuyer(openIdBuyer string) {
	req.Request.Params["open_id_buyer"] = openIdBuyer
}

func (req *GetUserCartSkuRequest) GetOpenIdBuyer() string {
	openIdBuyer, found := req.Request.Params["open_id_buyer"]
	if found {
		return openIdBuyer.(string)
	}
	return ""
}

func (req *GetUserCartSkuRequest) SetXidBuyer(xidBuyer string) {
	req.Request.Params["xid_buyer"] = xidIdBuyer
}

func (req *GetUserCartSkuRequest) GetXidIdBuyer() string {
	xidIdBuyer, found := req.Request.Params["xid_buyer"]
	if found {
		return xidIdBuyer.(string)
	}
	return ""
}
