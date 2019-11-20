package secret

import (
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

func (this *Client) PopOrderEnSearch(searchReq *PopOrderEnSearchRequest) ([]*order.OrderInfo, int, error) {
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
	orders, total, err := order.PopOrderEnSearch(req)
	if err != nil {
		return nil, 0, err
	}
	for _, orderInfo := range orders {
		if orderInfo.VatInfo != nil {
			if orderInfo.VatInfo.EncryptBankAccount != "" {
				if orderInfo.VatInfo.BankAccount, err = this.Decrypt(orderInfo.VatInfo.EncryptBankAccount, false); err != nil {
					return nil, 0, err
				}
			}
			if orderInfo.VatInfo.EncryptUserAddress != "" {
				if orderInfo.VatInfo.UserAddress, err = this.Decrypt(orderInfo.VatInfo.EncryptUserAddress, false); err != nil {
					return nil, 0, err
				}
			}
			if orderInfo.VatInfo.EncryptUserName != "" {
				if orderInfo.VatInfo.UserName, err = this.Decrypt(orderInfo.VatInfo.EncryptUserName, false); err != nil {
					return nil, 0, err
				}
			}
		}
		if orderInfo.InvoiceEasyInfo != nil {
			if orderInfo.InvoiceEasyInfo.EncryptInvoiceTitle != "" {
				if orderInfo.InvoiceEasyInfo.InvoiceTitle, err = this.Decrypt(orderInfo.InvoiceEasyInfo.EncryptInvoiceTitle, false); err != nil {
					return nil, 0, err
				}
			}
			if orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneeEmail != "" {
				if orderInfo.InvoiceEasyInfo.InvoiceConsigneeEmail, err = this.Decrypt(orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneeEmail, false); err != nil {
					return nil, 0, err
				}
			}
			if orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneePhone != "" {
				if orderInfo.InvoiceEasyInfo.InvoiceConsigneePhone, err = this.Decrypt(orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneePhone, false); err != nil {
					return nil, 0, err
				}
			}
		}
		if orderInfo.ConsigneeInfo != nil {
			if orderInfo.ConsigneeInfo.EncryptFullname != "" {
				if orderInfo.ConsigneeInfo.Fullname, err = this.Decrypt(orderInfo.ConsigneeInfo.EncryptFullname, false); err != nil {
					return nil, 0, err
				}
			}
			if orderInfo.ConsigneeInfo.EncryptTelephone != "" {
				if orderInfo.ConsigneeInfo.Telephone, err = this.Decrypt(orderInfo.ConsigneeInfo.EncryptTelephone, false); err != nil {
					return nil, 0, err
				}
			}
			if orderInfo.ConsigneeInfo.EncryptMobile != "" {
				if orderInfo.ConsigneeInfo.Mobile, err = this.Decrypt(orderInfo.ConsigneeInfo.EncryptMobile, false); err != nil {
					return nil, 0, err
				}
			}
			if orderInfo.ConsigneeInfo.EncryptFullAddress != "" {
				if orderInfo.ConsigneeInfo.FullAddress, err = this.Decrypt(orderInfo.ConsigneeInfo.EncryptFullAddress, false); err != nil {
					return nil, 0, err
				}
			}
		}
	}
	return orders, total, nil
}
