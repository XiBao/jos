package center

import "github.com/XiBao/jos/sdk"

type WritePersonInfoRequest struct {
	Request *sdk.Request
}

// create new request
func NewWritePersonInfoRequest() (req *WritePersonInfoRequest) {
	request := sdk.Request{MethodName: "jingdong.interact.center.api.service.write.writePersonInfo", Params: make(map[string]interface{})}
	req = &WritePersonInfoRequest{
		Request: &request,
	}
	return
}

func (req *WritePersonInfoRequest) SetAppName(AppName string) {
	req.Request.Params["appName"] = AppName
}

func (req *WritePersonInfoRequest) GetAppName() string {
	AppName, found := req.Request.Params["appName"]
	if found {
		return AppName.(string)
	}
	return ""
}

func (req *WritePersonInfoRequest) SetChannel(Channel uint8) {
	req.Request.Params["channel"] = Channel
}

func (req *WritePersonInfoRequest) GetChannel() uint8 {
	Channel, found := req.Request.Params["channel"]
	if found {
		return Channel.(uint8)
	}
	return 0
}

func (req *WritePersonInfoRequest) SetPin(Pin string) {
	req.Request.Params["pin"] = Pin
}

func (req *WritePersonInfoRequest) GetPin() string {
	Pin, found := req.Request.Params["pin"]
	if found {
		return Pin.(string)
	}
	return ""
}

func (req *WritePersonInfoRequest) SetOpenIdBuyer(OpenIdBuyer string) {
	req.Request.Params["open_id_buyer"] = OpenIdBuyer
}

func (req *WritePersonInfoRequest) GetOpenIdBuyer() string {
	OpenIdBuyer, found := req.Request.Params["open_id_buyer"]
	if found {
		return OpenIdBuyer.(string)
	}
	return ""
}

func (req *WritePersonInfoRequest) SetXIdBuyer(xIdBuyer string) {
	req.Request.Params["xid_buyer"] = xIdBuyer
}

func (req *WritePersonInfoRequest) GetXIdBuyer() string {
	xIdBuyer, found := req.Request.Params["xid_buyer"]
	if found {
		return xIdBuyer.(string)
	}
	return ""
}

func (req *WritePersonInfoRequest) SetProfileUrl(ProfileUrl string) {
	req.Request.Params["profileUrl"] = ProfileUrl
}

func (req *WritePersonInfoRequest) GetProfileUrl() string {
	ProfileUrl, found := req.Request.Params["profileUrl"]
	if found {
		return ProfileUrl.(string)
	}
	return ""
}

func (req *WritePersonInfoRequest) SetActivityId(ActivityId string) {
	req.Request.Params["activityId"] = ActivityId
}

func (req *WritePersonInfoRequest) GetActivityId() uint64 {
	ActivityId, found := req.Request.Params["activityId"]
	if found {
		return ActivityId.(uint64)
	}
	return 0
}

func (req *WritePersonInfoRequest) SetCreated(Created string) {
	req.Request.Params["created"] = Created
}

func (req *WritePersonInfoRequest) GetCreated() string {
	Created, found := req.Request.Params["created"]
	if found {
		return Created.(string)
	}
	return ""
}

func (req *WritePersonInfoRequest) SetStartTime(StartTime string) {
	req.Request.Params["startTime"] = StartTime
}

func (req *WritePersonInfoRequest) GetStartTime() string {
	StartTime, found := req.Request.Params["startTime"]
	if found {
		return StartTime.(string)
	}
	return ""
}

func (req *WritePersonInfoRequest) SetEndTime(EndTime string) {
	req.Request.Params["endTime"] = EndTime
}

func (req *WritePersonInfoRequest) GetEndTime() string {
	EndTime, found := req.Request.Params["endTime"]
	if found {
		return EndTime.(string)
	}
	return ""
}

func (req *WritePersonInfoRequest) SetId(Id uint64) {
	req.Request.Params["id"] = Id
}

func (req *WritePersonInfoRequest) GetId() uint64 {
	Id, found := req.Request.Params["id"]
	if found {
		return Id.(uint64)
	}
	return 0
}

func (req *WritePersonInfoRequest) SetType(Type string) {
	req.Request.Params["type"] = Type
}

func (req *WritePersonInfoRequest) GetType() uint8 {
	Type, found := req.Request.Params["type"]
	if found {
		return Type.(uint8)
	}
	return 0
}

func (req *WritePersonInfoRequest) SetActionType(ActionType string) {
	req.Request.Params["actionType"] = ActionType
}

func (req *WritePersonInfoRequest) GetActionType() uint8 {
	ActionType, found := req.Request.Params["actionType"]
	if found {
		return ActionType.(uint8)
	}
	return 0
}
