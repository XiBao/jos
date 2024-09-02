package imgzone

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/imgzone"
)

type PictureQueryRequest struct {
	api.BaseRequest
	PictureId     string `json:"picture_id,omitempty" codec:"picture_id,omitempty"`           // 图片ID
	PictureCateId int64  `json:"picture_cate_id,omitempty" codec:"picture_cate_id,omitempty"` // 图片所属分类ID
	PictureName   string `json:"picture_name,omitempty" codec:"picture_name,omitempty"`       // 图片名称，支持后模糊查询
	StartDate     string `json:"start_date,omitempty" codec:"start_date,omitempty"`           // 创建开始时间
	EndDate       string `json:"end_Date,omitempty" codec:"end_Date,omitempty"`               // 结束创建时间
	PageNum       int    `json:"page_num,omitempty" codec:"page_num,omitempty"`               // 页码值，对应分页结果页数，为空或非正整数时默认为1，超过最大页数返回空
	PageSize      int    `json:"page_size,omitempty" codec:"page_size,omitempty"`             // 每页条数，为空或非正整数时默认为20，最多返回100条
}

type PictureQueryResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *PictureQueryData   `json:"jingdong_imgzone_picture_query_responce,omitempty" codec:"jingdong_imgzone_picture_query_responce,omitempty"`
}

func (c PictureQueryResponse) IsError() bool {
	return c.ErrorResp != nil || c.Data == nil || c.Data.IsError()
}

func (c PictureQueryResponse) Error() string {
	if c.ErrorResp != nil {
		return c.ErrorResp.Error()
	}
	if c.Data != nil {
		return c.Data.Error()
	}
	return "no result data"
}

type PictureQueryData struct {
	ImgList    []Picture `json:"imgList,omitempty" codec:"imgList,omitempty"`
	Code       string    `json:"code,omitempty" codec:"code,omitempty"`
	Desc       string    `json:"desc1,omitempty" codec:"desc1,omitempty"`
	TotalNum   int       `json:"total_num,omitempty" codec:"total_num,omitempty"`
	ReturnCode int       `json:"return_code,omitempty" codec:"return_code,omitempty"`
}

func (c PictureQueryData) IsError() bool {
	return c.ReturnCode != 1
}

func (c PictureQueryData) Error() string {
	return sdk.ErrorString(c.ReturnCode, c.Desc)
}

// 查询图片信息接口，帮助获取图片信息
func PictureQuery(ctx context.Context, req *PictureQueryRequest) ([]Picture, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := imgzone.NewPictureQueryRequest()
	if req.PictureId != "" {
		r.SetPictureId(req.PictureId)
	}
	if req.PictureCateId > 0 {
		r.SetPictureCateId(req.PictureCateId)
	}
	if req.PictureName != "" {
		r.SetPictureName(req.PictureName)
	}
	if req.StartDate != "" {
		r.SetStartDate(req.StartDate)
	}
	if req.EndDate != "" {
		r.SetEndDate(req.EndDate)
	}
	if req.PageNum > 0 {
		r.SetPageNum(req.PageNum)
	}
	if req.PageSize > 0 {
		r.SetPageSize(req.PageSize)
	}

	var response PictureQueryResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.ImgList, nil
}
