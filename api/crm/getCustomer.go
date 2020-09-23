package crm

import (
    "errors"
    "fmt"
    . "github.com/XiBao/jos/api"
    "github.com/XiBao/jos/sdk"
    "github.com/XiBao/jos/sdk/request/crm"
    "github.com/daviddengcn/ljson"
)

type GetCustomerRequest struct {
    BaseRequest
    CustomerPin string `json:"customerPin,omitempty" codec:"customerPin,omitempty"`
}

type GetCustomerResponse struct {
    ErrorResp *ErrorResponnse      `json:"error_response,omitempty" codec:"error_response,omitempty"`
    Response  GetCustomerResponse1 `json:"jingdong_pop_crm_customer_getCustomer_responce,omitempty" codec:"jingdong_pop_crm_customer_getCustomer_responce,omitempty"`
}

type GetCustomerResponse1 struct {
    Result GetCustomerResult `json:"returnResult,omitempty" codec:"returnResult,omitempty"`
}

type GetCustomerResult struct {
    Desc string     `json:"desc,omitempty" codec:"desc,omitempty"`
    Code string     `json:"code,omitempty" codec:"code,omitempty"`
    Data CardMember `json:"data,omitempty" codec:"data,omitempty"`
}

func GetCustomer(req GetCustomerRequest) (CardMember, error) {

    client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
    client.Debug = req.Debug
    r := crm.GetCustomerRequest()
    r.SetCustomerPin(req.CustomerPin)
    result, err := client.Execute(r.Request, req.Session)
    if err != nil {
        return CardMember{}, err
    }
    if len(result) == 0 {
        return CardMember{}, fmt.Errorf("no result info")
    }
    var response GetCustomerResponse
    err = ljson.Unmarshal(result, &response)
    if err != nil {
        return CardMember{}, err
    }
    if response.ErrorResp != nil {
        return CardMember{}, response.ErrorResp
    }

    if response.Response.Result.Code != "200" {
        return CardMember{}, errors.New(response.Response.Result.Desc)
    }

    return response.Response.Result.Data, nil
}
