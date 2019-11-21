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

func (this *Client) PopOrderEnSearch(searchReq *PopOrderEnSearchRequest, decrypt bool) ([]*order.OrderInfo, int, error) {
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
	if !decrpyt {
		return orders, total, nil
	}
	for _, orderInfo := range orders {
		if orderInfo.VatInfo != nil {
			orderInfo.VatInfo.EncryptBankAccount = orderInfo.VatInfo.BankAccount
			if orderInfo.VatInfo.BankAccount, err = this.Decrypt(orderInfo.VatInfo.EncryptBankAccount, false); err != nil {
				return nil, err
			}
			orderInfo.VatInfo.EncryptUserAddress = orderInfo.VatInfo.UserAddress
			if orderInfo.VatInfo.UserAddress, err = this.Decrypt(orderInfo.VatInfo.EncryptUserAddress, false); err != nil {
				return nil, err
			}
			orderInfo.VatInfo.EncryptUserName = orderInfo.VatInfo.UserName
			if orderInfo.VatInfo.UserName, err = this.Decrypt(orderInfo.VatInfo.EncryptUserName, false); err != nil {
				return nil, err
			}
		}
		if orderInfo.InvoiceEasyInfo != nil {
			orderInfo.InvoiceEasyInfo.EncryptInvoiceTitle = orderInfo.InvoiceEasyInfo.InvoiceTitle
			if orderInfo.InvoiceEasyInfo.InvoiceTitle, err = this.Decrypt(orderInfo.InvoiceEasyInfo.EncryptInvoiceTitle, false); err != nil {
				return nil, err
			}
			orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneeEmail = orderInfo.InvoiceEasyInfo.InvoiceConsigneeEmail
			if orderInfo.InvoiceEasyInfo.InvoiceConsigneeEmail, err = this.Decrypt(orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneeEmail, false); err != nil {
				return nil, err
			}
			orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneePhone = orderInfo.InvoiceEasyInfo.InvoiceConsigneePhone
			if orderInfo.InvoiceEasyInfo.InvoiceConsigneePhone, err = this.Decrypt(orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneePhone, false); err != nil {
				return nil, err
			}
		}
		if orderInfo.ConsigneeInfo != nil {
			orderInfo.ConsigneeInfo.EncryptFullname = orderInfo.ConsigneeInfo.Fullname
			if orderInfo.ConsigneeInfo.Fullname, err = this.Decrypt(orderInfo.ConsigneeInfo.EncryptFullname, false); err != nil {
				return nil, err
			}
			orderInfo.ConsigneeInfo.EncryptTelephone = orderInfo.ConsigneeInfo.Telephone
			if orderInfo.ConsigneeInfo.Telephone, err = this.Decrypt(orderInfo.ConsigneeInfo.EncryptTelephone, false); err != nil {
				return nil, err
			}
			orderInfo.ConsigneeInfo.EncryptMobile = orderInfo.ConsigneeInfo.Mobile
			if orderInfo.ConsigneeInfo.Mobile, err = this.Decrypt(orderInfo.ConsigneeInfo.EncryptMobile, false); err != nil {
				return nil, err
			}
			orderInfo.ConsigneeInfo.EncryptFullAddress = orderInfo.ConsigneeInfo.FullAddress
			if orderInfo.ConsigneeInfo.FullAddress, err = this.Decrypt(orderInfo.ConsigneeInfo.EncryptFullAddress, false); err != nil {
				return nil, err
			}
		}
	}
	return orders, total, nil
}
