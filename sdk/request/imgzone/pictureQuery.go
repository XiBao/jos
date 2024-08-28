package imgzone

import (
	"github.com/XiBao/jos/sdk"
)

type PictureQueryRequest struct {
	Request *sdk.Request
}

func NewPictureQueryRequest() (req *PictureQueryRequest) {
	request := sdk.Request{MethodName: "jingdong.imgzone.picture.query", Params: make(map[string]interface{}, 3)}
	req = &PictureQueryRequest{
		Request: &request,
	}
	return
}

func (req *PictureQueryRequest) SetPictureId(pictureId string) {
	req.Request.Params["picture_id"] = pictureId
}

func (req *PictureQueryRequest) GetPictureId() string {
	pictureId, found := req.Request.Params["picture_id"]
	if found {
		return pictureId.(string)
	}
	return ""
}

func (req *PictureQueryRequest) SetPictureCateId(pictureCateId int64) {
	req.Request.Params["picture_cate_id"] = pictureCateId
}

func (req *PictureQueryRequest) GetPictureCateId() int64 {
	pictureCateId, found := req.Request.Params["picture_cate_id"]
	if found {
		return pictureCateId.(int64)
	}
	return 0
}

func (req *PictureQueryRequest) SetPictureName(pictureName string) {
	req.Request.Params["picture_name"] = pictureName
}

func (req *PictureQueryRequest) GetPictureName() string {
	pictureName, found := req.Request.Params["picture_name"]
	if found {
		return pictureName.(string)
	}
	return ""
}

func (req *PictureQueryRequest) SetStartDate(startDate string) {
	req.Request.Params["start_date"] = startDate
}

func (req *PictureQueryRequest) GetStartDate() string {
	startDate, found := req.Request.Params["start_date"]
	if found {
		return startDate.(string)
	}
	return ""
}

func (req *PictureQueryRequest) SetEndDate(endDate string) {
	req.Request.Params["end_date"] = endDate
}

func (req *PictureQueryRequest) GetEndDate() string {
	endDate, found := req.Request.Params["end_date"]
	if found {
		return endDate.(string)
	}
	return ""
}

func (req *PictureQueryRequest) SetPageNum(pageNum int) {
	req.Request.Params["page_num"] = pageNum
}

func (req *PictureQueryRequest) GetPageNum() int {
	pageNum, found := req.Request.Params["page_num"]
	if found {
		return pageNum.(int)
	}
	return 0
}

func (req *PictureQueryRequest) SetPageSize(pageSize int) {
	req.Request.Params["page_size"] = pageSize
}

func (req *PictureQueryRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(int)
	}
	return 0
}
