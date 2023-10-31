package follow

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/follow"
)

type FollowProductRequest struct {
	api.BaseRequest
	Pin       string `json:"pin,omitempty" codec:"pin,omitempty"`             //加密pin
	ProductId uint64 `json:"productId,omitempty" codec:"productId,omitempty"` //skuid
}

type FollowProductResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FollowProductData  `json:"jingdong_follow_product_complex_write_follow_responce,omitempty" codec:"jingdong_follow_product_complex_write_follow_responce,omitempty"`
}

func (r FollowProductResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r FollowProductResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type FollowProductData struct {
	Result *FollowProductResult `json:"follow_result,omitempty" codec:"follow_result,omitempty"`
	Code   string               `json:"code"`
}

func (r FollowProductData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r FollowProductData) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type FollowProductResult struct {
	Msg  string `json:"msg,omitempty" codec:"msg,omitempty"`
	Code string `json:"code,omitempty" codec:"code,omitempty"` //状态码
	Data bool   `json:"data,omitempty" codec:"data,omitempty"` //是否成功
}

func (r FollowProductResult) IsError() bool {
	return r.Msg != ""
}

func (r FollowProductResult) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.Msg)
}

// TODO  通过pin将商品加入用户关注栏
func FollowProduct(req *FollowProductRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := follow.NewFollowProductComplexWriteFollow()
	r.SetPin(req.Pin)
	r.SetProductId(req.ProductId)

	var response FollowProductResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Data.Result.Data, nil
}
