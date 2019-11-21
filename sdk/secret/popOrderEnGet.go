package secret

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/order"
)

func (this *Client) PopOrderEnGet(orderId uint64, orderState []string, optionalFields []string, decrpyt bool) (*order.OrderInfo, error) {
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
	if !decrypt {
		return orderInfo, nil
	}
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
	return orderInfo, nil
}
