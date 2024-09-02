package helper

import (
	"context"
	"time"

	"github.com/XiBao/jos/api"
	sellerVender "github.com/XiBao/jos/api/seller/vender"
	"github.com/XiBao/jos/api/vender"
)

type GetShopVenderInfoRequest struct {
	api.BaseRequest
}

func GetShopVenderInfo(ctx context.Context, req *GetShopVenderInfoRequest) (*vender.ShopInfo, error) {
	shop, err := vender.ShopQuery(ctx, &vender.ShopQueryRequest{
		BaseRequest: api.BaseRequest{
			Session:  req.Session,
			AnApiKey: req.AnApiKey,
		},
	})
	if err != nil {
		return nil, err
	}
	venderInfo, err := sellerVender.InfoGet(ctx, &sellerVender.InfoGetRequest{
		BaseRequest: api.BaseRequest{
			Session:  req.Session,
			AnApiKey: req.AnApiKey,
		},
	})
	if err != nil {
		return nil, err
	}

	shop.ColType = venderInfo.ColType
	shop.OpenTimeStr = time.Unix(shop.OpenTime/1000, shop.OpenTime%1000).Format("2006-01-02 15:04:05")

	return shop, nil
}
