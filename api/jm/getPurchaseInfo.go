package jm

import (
	. "github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/jm"
)

type GetPurchaseInfoRequest struct {
	BaseRequest
}

type GetPurchaseInfoResponse struct {
	ErrorResp *ErrorResponnse             `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      GetPurchaseInfoResponseData `json:"jingdong_getPurchaseInfo_responce,omitempty" codec:"jingdong_getPurchaseInfo_responce,omitempty"`
}

func (r GetPurchaseInfoResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data.IsError()
}

func (r GetPurchaseInfoResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return r.Data.Error()
}

type GetPurchaseInfoResponseData struct {
	ReturnType GetPurchaseInfoReturnType `json:"returnType"`
}

func (r GetPurchaseInfoResponseData) IsError() bool {
	return r.ReturnType.IsError()
}

func (r GetPurchaseInfoResponseData) Error() string {
	return r.ReturnType.Error()
}

type GetPurchaseInfoReturnType struct {
	Success          bool           `json:"success"`
	Message          string         `json:"message,omitempty"`
	Code             int            `json:"errorCode"`
	PurchaseInfoList []PurchaseInfo `json:"purchaseInfoList,omitempty"`
}

func (r GetPurchaseInfoReturnType) IsError() bool {
	return r.Code != 0 && r.Code != 200
}

func (r GetPurchaseInfoReturnType) Error() string {
	return sdk.ErrorString(r.Code, r.Message)
}

func GetPurchaseInfo(req *GetPurchaseInfoRequest) ([]PurchaseInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := jm.NewGetPurchaseInfoRequest()

	var response GetPurchaseInfoResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.ReturnType.PurchaseInfoList, nil
}
