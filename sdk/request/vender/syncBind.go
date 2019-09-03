package vender

import (
	"github.com/XiBao/jos/sdk"
)

type VenderSyncBindRequest struct {
	Request *sdk.Request
}

// create new request
func NewVenderSyncBindRequest() (req *VenderSyncBindRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.vender.syncBind", Params: make(map[string]interface{}, 13)}
	req = &VenderSyncBindRequest{
		Request: &request,
	}
	return
}

func (req *VenderSyncBindRequest) SetBirthday(birthday string) {
	req.Request.Params["birthday"] = birthday
}

func (req *VenderSyncBindRequest) GetBirthday() string {
	birthday, found := req.Request.Params["birthday"]
	if found {
		return birthday.(string)
	}
	return ""
}

func (req *VenderSyncBindRequest) SetGender(gender string) {
	req.Request.Params["gender"] = gender
}

func (req *VenderSyncBindRequest) GetGender() string {
	gender, found := req.Request.Params["gender"]
	if found {
		return gender.(string)
	}
	return ""
}

func (req *VenderSyncBindRequest) SetCity(city string) {
	req.Request.Params["city"] = city
}

func (req *VenderSyncBindRequest) GetCity() string {
	city, found := req.Request.Params["city"]
	if found {
		return city.(string)
	}
	return ""
}

func (req *VenderSyncBindRequest) SetChannel(channel uint16) {
	req.Request.Params["channel"] = channel
}

func (req *VenderSyncBindRequest) GetChannel() uint16 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint16)
	}
	return 0
}

func (req *VenderSyncBindRequest) SetCardNo(cardNo string) {
	req.Request.Params["cardNo"] = cardNo
}

func (req *VenderSyncBindRequest) GetCardNo() string {
	cardNo, found := req.Request.Params["cardNo"]
	if found {
		return cardNo.(string)
	}
	return ""
}

func (req *VenderSyncBindRequest) SetPhoneNo(phoneNo string) {
	req.Request.Params["phoneNo"] = phoneNo
}

func (req *VenderSyncBindRequest) GetPhoneNo() string {
	phoneNo, found := req.Request.Params["phoneNo"]
	if found {
		return phoneNo.(string)
	}
	return ""
}

func (req *VenderSyncBindRequest) SetCustomerLevel(customerLevel uint8) {
	req.Request.Params["customerLevel"] = customerLevel
}

func (req *VenderSyncBindRequest) GetCustomerLevel() uint8 {
	customerLevel, found := req.Request.Params["customerLevel"]
	if found {
		return customerLevel.(uint8)
	}
	return 0
}

func (req *VenderSyncBindRequest) SetExtend(extend string) {
	req.Request.Params["extend"] = extend
}

func (req *VenderSyncBindRequest) GetExtend() string {
	extend, found := req.Request.Params["extend"]
	if found {
		return extend.(string)
	}
	return ""
}

func (req *VenderSyncBindRequest) SetCustomerType(customerType uint8) {
	req.Request.Params["customerType"] = customerType
}

func (req *VenderSyncBindRequest) GetCustomerType() uint8 {
	customerType, found := req.Request.Params["customerType"]
	if found {
		return customerType.(uint8)
	}
	return 0
}

func (req *VenderSyncBindRequest) SetProvince(province string) {
	req.Request.Params["province"] = province
}

func (req *VenderSyncBindRequest) GetProvince() string {
	province, found := req.Request.Params["province"]
	if found {
		return province.(string)
	}
	return ""
}

func (req *VenderSyncBindRequest) SetStreet(street string) {
	req.Request.Params["street"] = street
}

func (req *VenderSyncBindRequest) GetStreet() string {
	street, found := req.Request.Params["street"]
	if found {
		return street.(string)
	}
	return ""
}

func (req *VenderSyncBindRequest) SetCustomerPin(customerPin string) {
	req.Request.Params["customerPin"] = customerPin
}

func (req *VenderSyncBindRequest) GetCustomerPin() string {
	customerPin, found := req.Request.Params["customerPin"]
	if found {
		return customerPin.(string)
	}
	return ""
}

func (req *VenderSyncBindRequest) SetStatus(status uint8) {
	req.Request.Params["status"] = status
}

func (req *VenderSyncBindRequest) GetStatus() uint8 {
	status, found := req.Request.Params["status"]
	if found {
		return status.(uint8)
	}
	return 0
}
