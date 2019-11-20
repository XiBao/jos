package secret

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/order"
)

func (this *Client) PopOrderEnGet(orderId uint64, orderState []string, optionalFields []string) (*order.OrderInfo, error) {
	req := &order.PopOrderEnGetRequest{
		BaseRequest: api.BaseRequest{
			AnApiKey: &api.ApiKey{
				Key:    this.AppKey,
				Secret: this.AppSecret,
			},
			Session: this.AccessToken,
		},
		OrderId:        orderId,
		OrderState:     orderState,
		OptionalFields: optionalFields,
	}
	orderInfo, err := order.PopOrderEnGet(req)
	if err != nil {
		return nil, err
	}
	if orderInfo.VatInfo != nil {
		if orderInfo.VatInfo.EncryptBankAccount != "" {
			if orderInfo.VatInfo.BankAccount, err = this.Decrypt(orderInfo.VatInfo.EncryptBankAccount, false); err != nil {
				return nil, err
			}
		}
		if orderInfo.VatInfo.EncryptUserAddress != "" {
			if orderInfo.VatInfo.UserAddress, err = this.Decrypt(orderInfo.VatInfo.EncryptUserAddress, false); err != nil {
				return nil, err
			}
		}
		if orderInfo.VatInfo.EncryptUserName != "" {
			if orderInfo.VatInfo.UserName, err = this.Decrypt(orderInfo.VatInfo.EncryptUserName, false); err != nil {
				return nil, err
			}
		}
	}
	if orderInfo.InvoiceEasyInfo != nil {
		if orderInfo.InvoiceEasyInfo.EncryptInvoiceTitle != "" {
			if orderInfo.InvoiceEasyInfo.InvoiceTitle, err = this.Decrypt(orderInfo.InvoiceEasyInfo.EncryptInvoiceTitle, false); err != nil {
				return nil, err
			}
		}
		if orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneeEmail != "" {
			if orderInfo.InvoiceEasyInfo.InvoiceConsigneeEmail, err = this.Decrypt(orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneeEmail, false); err != nil {
				return nil, err
			}
		}
		if orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneePhone != "" {
			if orderInfo.InvoiceEasyInfo.InvoiceConsigneePhone, err = this.Decrypt(orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneePhone, false); err != nil {
				return nil, err
			}
		}
	}
	if orderInfo.ConsigneeInfo != nil {
		if orderInfo.ConsigneeInfo.EncryptFullname != "" {
			if orderInfo.ConsigneeInfo.Fullname, err = this.Decrypt(orderInfo.ConsigneeInfo.EncryptFullname, false); err != nil {
				return nil, err
			}
		}
		if orderInfo.ConsigneeInfo.EncryptTelephone != "" {
			if orderInfo.ConsigneeInfo.Telephone, err = this.Decrypt(orderInfo.ConsigneeInfo.EncryptTelephone, false); err != nil {
				return nil, err
			}
		}
		if orderInfo.ConsigneeInfo.EncryptMobile != "" {
			if orderInfo.ConsigneeInfo.Mobile, err = this.Decrypt(orderInfo.ConsigneeInfo.EncryptMobile, false); err != nil {
				return nil, err
			}
		}
		if orderInfo.ConsigneeInfo.EncryptFullAddress != "" {
			if orderInfo.ConsigneeInfo.FullAddress, err = this.Decrypt(orderInfo.ConsigneeInfo.EncryptFullAddress, false); err != nil {
				return nil, err
			}
		}
	}
	return orderInfo, nil
}
