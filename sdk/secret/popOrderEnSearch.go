package secret

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/order"
)

type PopOrderEnSearchRequest struct {
	StartDate      string   `json:"start_date,omitempty" codec:"start_date,omitempty"`
	EndDate        string   `json:"end_date,omitempty" codec:"end_date,omitempty"`
	OrderState     []string `json:"order_state,omitempty" codec:"order_state,omitempty"`
	OptionalFields []string `json:"optional_fields,omitempty" codec:"optional_fields,omitempty"`
	Page           int      `json:"page,omitempty" codec:"page,omitempty"`
	PageSize       int      `json:"page_size,omitempty" codec:"page_size,omitempty"`
	SortType       uint8    `json:"sort_type,omitempty" codec:"sort_type,omitempty"`
	DateType       uint8    `json:"date_type,omitempty" codec:"date_type,omitempty"`
}

func (this *Client) PopOrderEnSearch(ctx context.Context, searchReq *PopOrderEnSearchRequest, decrypt bool) ([]order.OrderInfo, int, error) {
	req := &order.PopOrderEnSearchRequest{
		BaseRequest: api.BaseRequest{
			AnApiKey: &api.ApiKey{
				Key:    this.AppKey,
				Secret: this.AppSecret,
			},
			Session: this.AccessToken,
		},
		StartDate:      searchReq.StartDate,
		EndDate:        searchReq.EndDate,
		OrderState:     searchReq.OrderState,
		OptionalFields: searchReq.OptionalFields,
		Page:           searchReq.Page,
		PageSize:       searchReq.PageSize,
		SortType:       searchReq.SortType,
		DateType:       searchReq.DateType,
	}
	orders, total, err := order.PopOrderEnSearch(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	if decrypt {
		for idx, orderInfo := range orders {
			if err := this.DecryptOrderInfo(ctx, &orderInfo, false); err != nil {
				return nil, 0, err
			}
			orders[idx] = orderInfo
		}
	}

	return orders, total, nil
}
