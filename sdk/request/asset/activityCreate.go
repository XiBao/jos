package asset

import (
	"github.com/XiBao/jos/sdk"
)

type ActivityCreateRequest struct {
	Request *sdk.Request
}

// create new request
func NewActivityCreateRequest() (req *ActivityCreateRequest) {
	request := sdk.Request{MethodName: "jingdong.asset.activity.create", Params: make(map[string]interface{}, 6)}
	req = &ActivityCreateRequest{
		Request: &request,
	}
	return
}

func (req *ActivityCreateRequest) SetActivityId(activityId string) {
	req.Request.Params["activity_id"] = activityId
}

func (req *ActivityCreateRequest) GetActivityId() string {
	activityId, found := req.Request.Params["activity_id"]
	if found {
		return activityId.(string)
	}
	return ""
}

func (req *ActivityCreateRequest) SetActivityName(activityName string) {
	req.Request.Params["activity_name"] = activityName
}

func (req *ActivityCreateRequest) GetActivityName() string {
	activityName, found := req.Request.Params["activity_name"]
	if found {
		return activityName.(string)
	}
	return ""
}

func (req *ActivityCreateRequest) SetBeginDate(beginDate string) {
	req.Request.Params["begin_date"] = beginDate
}

func (req *ActivityCreateRequest) GetBeginDate() string {
	beginDate, found := req.Request.Params["begin_date"]
	if found {
		return beginDate.(string)
	}
	return ""
}

func (req *ActivityCreateRequest) SetEndDate(endDate string) {
	req.Request.Params["end_date"] = endDate
}

func (req *ActivityCreateRequest) GetEndDate() string {
	endDate, found := req.Request.Params["end_date"]
	if found {
		return endDate.(string)
	}
	return ""
}

func (req *ActivityCreateRequest) SetTool(tool string) {
	req.Request.Params["tool"] = tool
}

func (req *ActivityCreateRequest) GetTool() string {
	tool, found := req.Request.Params["tool"]
	if found {
		return tool.(string)
	}
	return ""
}

func (req *ActivityCreateRequest) SetDetails(details string) {
	req.Request.Params["details"] = details
}

func (req *ActivityCreateRequest) GetDetails() string {
	details, found := req.Request.Params["details"]
	if found {
		return details.(string)
	}
	return ""
}
