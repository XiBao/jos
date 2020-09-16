package center

import "github.com/XiBao/jos/sdk"

type WriteWritePersonInfoRequest struct{
    Request *sdk.Request
}


// create new request
func NewWriteWritePersonInfoRequest() (req *WriteWritePersonInfoRequest) {
    request := sdk.Request{MethodName: "jingdong.interact.center.api.service.write.writePersonInfo", Params: make(map[string]interface{})}
    req = &WriteWritePersonInfoRequest{
        Request: &request,
    }
    return
}


func (req *WriteWritePersonInfoRequest) SetAppName(AppName string) {
    req.Request.Params["appName"] = AppName
}

func (req *WriteWritePersonInfoRequest) GetAppName() string {
    AppName, found := req.Request.Params["appName"]
    if found {
        return AppName.(string)
    }
    return ""
}


func (req *WriteWritePersonInfoRequest) SetChannel(Channel uint8) {
    req.Request.Params["channel"] = Channel
}

func (req *WriteWritePersonInfoRequest) GetChannel() uint8 {
    Channel, found := req.Request.Params["channel"]
    if found {
        return Channel.(uint8)
    }
    return 0
}


func (req *WriteWritePersonInfoRequest) SetPin(Pin string) {
    req.Request.Params["pin"] = Pin
}

func (req *WriteWritePersonInfoRequest) GetPin() string {
    Pin, found := req.Request.Params["pin"]
    if found {
        return Pin.(string)
    }
    return ""
}


func (req *WriteWritePersonInfoRequest) SetOpenIdBuyer(OpenIdBuyer string) {
    req.Request.Params["open_id_buyer"] = OpenIdBuyer
}

func (req *WriteWritePersonInfoRequest) GetOpenIdBuyer() string {
    OpenIdBuyer, found := req.Request.Params["open_id_buyer"]
    if found {
        return OpenIdBuyer.(string)
    }
    return ""
}


func (req *WriteWritePersonInfoRequest) SetProfileUrl(ProfileUrl string) {
    req.Request.Params["profileUrl"] = ProfileUrl
}

func (req *WriteWritePersonInfoRequest) GetProfileUrl() string {
    ProfileUrl, found := req.Request.Params["profileUrl"]
    if found {
        return ProfileUrl.(string)
    }
    return ""
}


func (req *WriteWritePersonInfoRequest) SetActivityId(ActivityId string) {
    req.Request.Params["activityId"] = ActivityId
}

func (req *WriteWritePersonInfoRequest) GetActivityId() uint64 {
    ActivityId, found := req.Request.Params["activityId"]
    if found {
        return ActivityId.(uint64)
    }
    return 0
}


func (req *WriteWritePersonInfoRequest) SetCreated(Created string) {
    req.Request.Params["created"] = Created
}

func (req *WriteWritePersonInfoRequest) GetCreated() string {
    Created, found := req.Request.Params["created"]
    if found {
        return Created.(string)
    }
    return ""
}


func (req *WriteWritePersonInfoRequest) SetStartTime(StartTime string) {
    req.Request.Params["startTime"] = StartTime
}

func (req *WriteWritePersonInfoRequest) GetStartTime() string {
    StartTime, found := req.Request.Params["startTime"]
    if found {
        return StartTime.(string)
    }
    return ""
}


func (req *WriteWritePersonInfoRequest) SetEndTime(EndTime string) {
    req.Request.Params["endTime"] = EndTime
}

func (req *WriteWritePersonInfoRequest) GetEndTime() string {
    EndTime, found := req.Request.Params["endTime"]
    if found {
        return EndTime.(string)
    }
    return ""
}


func (req *WriteWritePersonInfoRequest) SetId(Id uint64) {
    req.Request.Params["id"] = Id
}

func (req *WriteWritePersonInfoRequest) GetId() uint64 {
    Id, found := req.Request.Params["id"]
    if found {
        return Id.(uint64)
    }
    return 0
}


func (req *WriteWritePersonInfoRequest) SetType(Type string) {
    req.Request.Params["type"] = Type
}

func (req *WriteWritePersonInfoRequest) GetType() uint8 {
    Type, found := req.Request.Params["type"]
    if found {
        return Type.(uint8)
    }
    return 0
}


func (req *WriteWritePersonInfoRequest) SetActionType(ActionType string) {
    req.Request.Params["actionType"] = ActionType
}

func (req *WriteWritePersonInfoRequest) GetActionType() uint8 {
    ActionType, found := req.Request.Params["actionType"]
    if found {
        return ActionType.(uint8)
    }
    return 0
}

