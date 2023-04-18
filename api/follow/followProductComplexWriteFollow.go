package follow

import (
	"encoding/json"
	"errors"

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

type FollowProductData struct {
	Result *FollowProductResult `json:"follow_result,omitempty" codec:"follow_result,omitempty"`
	Code   string               `json:"code"`
}

type FollowProductResult struct {
	Msg  string `json:"msg,omitempty" codec:"msg,omitempty"`
	Code string `json:"code,omitempty" codec:"code,omitempty"` //状态码
	Data bool   `json:"data,omitempty" codec:"data,omitempty"` //是否成功
}

// TODO  通过pin将商品加入用户关注栏
func FollowProduct(req *FollowProductRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := follow.NewFollowProductComplexWriteFollow()
	r.SetPin(req.Pin)
	r.SetProductId(req.ProductId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("No result info.")
	}
	var response FollowProductResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}

	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	if response.Data.Result.Msg != "" && !response.Data.Result.Data {
		return false, errors.New(response.Data.Result.Msg)
	}

	return response.Data.Result.Data, nil
}
