package isv

import (
	"github.com/XiBao/jos/sdk"
)

type IsvUploadOrderLogRequest struct {
	Request *sdk.Request
}

// create new request
func NewIsvUploadOrderLogRequest() (req *IsvUploadOrderLogRequest) {
	request := sdk.Request{MethodName: "jingdong.isv.uploadOrderInfoLog", Params: make(map[string]interface{}, 11)}
	req = &IsvUploadOrderLogRequest{
		Request: &request,
	}
	return
}

func (req *IsvUploadOrderLogRequest) SetUserIp(userIp string) {
	req.Request.Params["user_ip"] = userIp
}

func (req *IsvUploadOrderLogRequest) GetUserIp() string {
	userIp, found := req.Request.Params["user_ip"]
	if found {
		return userIp.(string)
	}
	return ""
}

func (req *IsvUploadOrderLogRequest) SetAppName(appName string) {
	req.Request.Params["app_name"] = appName
}

func (req *IsvUploadOrderLogRequest) GetAppName() string {
	appName, found := req.Request.Params["app_name"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *IsvUploadOrderLogRequest) SetJosAppKey(josAppKey string) {
	req.Request.Params["josAppKey"] = josAppKey
}

func (req *IsvUploadOrderLogRequest) GetJosAppKey() string {
	josAppKey, found := req.Request.Params["josAppKey"]
	if found {
		return josAppKey.(string)
	}
	return ""
}

func (req *IsvUploadOrderLogRequest) SetJdId(jdId string) {
	req.Request.Params["jd_id"] = jdId
}

func (req *IsvUploadOrderLogRequest) GetJdId() string {
	jdId, found := req.Request.Params["jd_id"]
	if found {
		return jdId.(string)
	}
	return ""
}

func (req *IsvUploadOrderLogRequest) SetDeviceId(deviceId string) {
	req.Request.Params["device_id"] = deviceId
}

func (req *IsvUploadOrderLogRequest) GetDeviceId() string {
	deviceId, found := req.Request.Params["device_id"]
	if found {
		return deviceId.(string)
	}
	return ""
}

func (req *IsvUploadOrderLogRequest) SetUserId(userId string) {
	req.Request.Params["user_id"] = userId
}

func (req *IsvUploadOrderLogRequest) GetUserId() string {
	userId, found := req.Request.Params["user_id"]
	if found {
		return userId.(string)
	}
	return ""
}
func (req *IsvUploadOrderLogRequest) SetFileMd5(fileMd5 string) {
	req.Request.Params["file_md5"] = fileMd5
}

func (req *IsvUploadOrderLogRequest) GetFileMd5() string {
	fileMd5, found := req.Request.Params["file_md5"]
	if found {
		return fileMd5.(string)
	}
	return ""
}

func (req *IsvUploadOrderLogRequest) SetOrderIds(orderIds string) {
	req.Request.Params["order_ids"] = orderIds
}

func (req *IsvUploadOrderLogRequest) GetOrderIds() string {
	orderIds, found := req.Request.Params["order_ids"]
	if found {
		return orderIds.(string)
	}
	return ""
}

func (req *IsvUploadOrderLogRequest) SetOperation(operation int) {
	req.Request.Params["operation"] = operation
}

func (req *IsvUploadOrderLogRequest) GetOperation() int {
	operation, found := req.Request.Params["operation"]
	if found {
		return operation.(int)
	}
	return 0
}

func (req *IsvUploadOrderLogRequest) SetUrl(url string) {
	req.Request.Params["url"] = url
}

func (req *IsvUploadOrderLogRequest) GetUrl() string {
	url, found := req.Request.Params["url"]
	if found {
		return url.(string)
	}
	return ""
}

func (req *IsvUploadOrderLogRequest) SetTimestamp(timestamp uint64) {
	req.Request.Params["time_stamp"] = timestamp
}

func (req *IsvUploadOrderLogRequest) GetTimestamp() uint64 {
	timestamp, found := req.Request.Params["time_stamp"]
	if found {
		return timestamp.(uint64)
	}
	return 0
}
