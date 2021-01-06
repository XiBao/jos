package crm

import "github.com/XiBao/jos/sdk"

type CustomerRequest struct {
    Request *sdk.Request
}

// create new request
func GetCustomerRequest() (req *CustomerRequest) {
    request := sdk.Request{MethodName: "jingdong.pop.crm.customer.getCustomer", Params: make(map[string]interface{})}
    req = &CustomerRequest{
        Request: &request,
    }
    return
}


func (req *CustomerRequest) SetCustomerPin(customerPin string) {
    req.Request.Params["customerPin"] = customerPin
}

func (req *CustomerRequest) GetCustomerPin() string {
    customerPin, found := req.Request.Params["customerPin"]
    if found {
        return customerPin.(string)
    }
    return ""
}