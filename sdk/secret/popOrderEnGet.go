package secret

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/order"
)

func (this *Client) PopOrderEnGet(ctx context.Context, orderId uint64, orderState []string, optionalFields []string, decrypt bool) (*order.OrderInfo, error) {
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
	orderInfo, err := order.PopOrderEnGet(ctx, req)
	if err != nil {
		return nil, err
	}
	if decrypt {
		err = this.DecryptOrderInfo(ctx, orderInfo, false)
		if err != nil {
			return nil, err
		}
	}

	return orderInfo, nil
}

func (this *Client) DecryptOrderInfo(ctx context.Context, orderInfo *order.OrderInfo, usePrivateKey bool) (err error) {
	if orderInfo.VatInfo != nil {
		orderInfo.VatInfo.EncryptBankAccount = orderInfo.VatInfo.BankAccount
		if orderInfo.VatInfo.BankAccount, err = this.Decrypt(ctx, orderInfo.VatInfo.EncryptBankAccount, usePrivateKey); err != nil {
			return err
		}
		orderInfo.VatInfo.EncryptUserAddress = orderInfo.VatInfo.UserAddress
		if orderInfo.VatInfo.UserAddress, err = this.Decrypt(ctx, orderInfo.VatInfo.EncryptUserAddress, usePrivateKey); err != nil {
			return err
		}
		orderInfo.VatInfo.EncryptUserName = orderInfo.VatInfo.UserName
		if orderInfo.VatInfo.UserName, err = this.Decrypt(ctx, orderInfo.VatInfo.EncryptUserName, usePrivateKey); err != nil {
			return err
		}
	}
	if orderInfo.InvoiceEasyInfo != nil {
		orderInfo.InvoiceEasyInfo.EncryptInvoiceTitle = orderInfo.InvoiceEasyInfo.InvoiceTitle
		if orderInfo.InvoiceEasyInfo.InvoiceTitle, err = this.Decrypt(ctx, orderInfo.InvoiceEasyInfo.EncryptInvoiceTitle, usePrivateKey); err != nil {
			return err
		}
		orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneeEmail = orderInfo.InvoiceEasyInfo.InvoiceConsigneeEmail
		if orderInfo.InvoiceEasyInfo.InvoiceConsigneeEmail, err = this.Decrypt(ctx, orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneeEmail, usePrivateKey); err != nil {
			return err
		}
		orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneePhone = orderInfo.InvoiceEasyInfo.InvoiceConsigneePhone
		if orderInfo.InvoiceEasyInfo.InvoiceConsigneePhone, err = this.Decrypt(ctx, orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneePhone, usePrivateKey); err != nil {
			return err
		}
	}
	if orderInfo.ConsigneeInfo != nil {
		orderInfo.ConsigneeInfo.EncryptFullname = orderInfo.ConsigneeInfo.Fullname
		if orderInfo.ConsigneeInfo.Fullname, err = this.Decrypt(ctx, orderInfo.ConsigneeInfo.EncryptFullname, usePrivateKey); err != nil {
			return err
		}
		orderInfo.ConsigneeInfo.EncryptTelephone = orderInfo.ConsigneeInfo.Telephone
		if orderInfo.ConsigneeInfo.Telephone, err = this.Decrypt(ctx, orderInfo.ConsigneeInfo.EncryptTelephone, usePrivateKey); err != nil {
			return err
		}
		orderInfo.ConsigneeInfo.EncryptMobile = orderInfo.ConsigneeInfo.Mobile
		if orderInfo.ConsigneeInfo.Mobile, err = this.Decrypt(ctx, orderInfo.ConsigneeInfo.EncryptMobile, usePrivateKey); err != nil {
			return err
		}
		orderInfo.ConsigneeInfo.EncryptFullAddress = orderInfo.ConsigneeInfo.FullAddress
		if orderInfo.ConsigneeInfo.FullAddress, err = this.Decrypt(ctx, orderInfo.ConsigneeInfo.EncryptFullAddress, usePrivateKey); err != nil {
			return err
		}
	}
	return nil
}

func (this *Client) EncryptOrderInfo(ctx context.Context, orderInfo *order.OrderInfo, usePrivateKey bool) (err error) {
	if orderInfo.VatInfo != nil {
		if orderInfo.VatInfo.EncryptBankAccount, err = this.Encrypt(ctx, orderInfo.VatInfo.BankAccount, usePrivateKey); err != nil {
			return err
		}
		if orderInfo.VatInfo.EncryptUserAddress, err = this.Decrypt(ctx, orderInfo.VatInfo.UserAddress, usePrivateKey); err != nil {
			return err
		}
		if orderInfo.VatInfo.EncryptUserName, err = this.Decrypt(ctx, orderInfo.VatInfo.UserName, usePrivateKey); err != nil {
			return err
		}
	}
	if orderInfo.InvoiceEasyInfo != nil {
		if orderInfo.InvoiceEasyInfo.EncryptInvoiceTitle, err = this.Decrypt(ctx, orderInfo.InvoiceEasyInfo.InvoiceTitle, usePrivateKey); err != nil {
			return err
		}
		if orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneeEmail, err = this.Decrypt(ctx, orderInfo.InvoiceEasyInfo.InvoiceConsigneeEmail, usePrivateKey); err != nil {
			return err
		}
		if orderInfo.InvoiceEasyInfo.EncryptInvoiceConsigneePhone, err = this.Decrypt(ctx, orderInfo.InvoiceEasyInfo.InvoiceConsigneePhone, usePrivateKey); err != nil {
			return err
		}
	}
	if orderInfo.ConsigneeInfo != nil {
		if orderInfo.ConsigneeInfo.EncryptFullname, err = this.Decrypt(ctx, orderInfo.ConsigneeInfo.Fullname, usePrivateKey); err != nil {
			return err
		}
		if orderInfo.ConsigneeInfo.EncryptTelephone, err = this.Decrypt(ctx, orderInfo.ConsigneeInfo.Telephone, usePrivateKey); err != nil {
			return err
		}
		if orderInfo.ConsigneeInfo.EncryptMobile, err = this.Decrypt(ctx, orderInfo.ConsigneeInfo.Mobile, usePrivateKey); err != nil {
			return err
		}
		if orderInfo.ConsigneeInfo.EncryptFullAddress, err = this.Decrypt(ctx, orderInfo.ConsigneeInfo.FullAddress, usePrivateKey); err != nil {
			return err
		}
	}
	return nil
}
