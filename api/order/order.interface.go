package order

type JdOrderInfo struct {
	OrderType          int       `json:"orderType,omitempty" codec:"orderType,omitempty"`                   // 订单类型（22 SOP； 75 LOC） 可选字段，需要在输入参数optional_fields中写入才能返回
	OrderCreated       int64     `json:"orderCreated,omitempty" codec:"orderCreated,omitempty"`             // 订单创建时间
	Skus               []SkuInfo `json:"skus,omitempty" codec:"skus,omitempty"`                             // 商品信息
	OrderId            uint64    `json:"orderId,omitempty" codec:"orderId,omitempty"`                       // 订单id
	City               string    `json:"city,omitempty" codec:"city,omitempty"`                             // 市
	Town               string    `json:"town,omitempty" codec:"town,omitempty"`                             // 县
	PaymentConfirmTime int64     `json:"paymentConfirmTime,omitempty" codec:"paymentConfirmTime,omitempty"` // 支付时间
	Province           string    `json:"province,omitempty" codec:"province,omitempty"`                     // 省
	ParentId           uint64    `json:"parentId,omitempty" codec:"parentId,omitempty"`                     // 父订单
	Pin                string    `json:"pin,omitempty" codec:"pin,omitempty"`                               // pin
	County             string    `json:"county,omitempty" codec:"county,omitempty"`                         // 区
	OrderEndTime       int64     `json:"orderEndTime,omitempty" codec:"orderEndTime,omitempty"`             // 订单完成时间
	OrderPrice         float64   `json:"orderPrice,omitempty" codec:"orderPrice,omitempty"`                 // 订单价格
	ShopId             uint64    `json:"shopId,omitempty" codec:"shopId,omitempty"`                         // 商家ID
	Status             string    `json:"status,omitempty" codec:"status,omitempty"`                         // NOT_PAY 等待付款 ；SUSPEND 暂停； WAIT_DELIVERY 等待出库；WAIT_SHIPMENTS 等待发货；FINISH 完成 ；CANCEL 取消；LOCK 锁定;
	SendPay            string    `json:"sendPay,omitempty" codec:"sendPay,omitempty"`                       // SendPay
	DeliveryTime       int64     `json:"deliveryTime,omitempty" codec:"deliveryTime,omitempty"`             // 出库发货时间
	PaymentType        string    `json:"paymentType,omitempty" codec:"paymentType,omitempty"`               // 支付方式
	OrderCancelTime    int64     `json:"orderCancelTime,omitempty" codec:"orderCancelTime,omitempty"`       // 取消时间
	DParentId          uint64    `json:"dParentId,omitempty" codec:"dParentId,omitempty"`                   // 父订单号
	OpenIdBuyer        string    `json:"open_id_buyer,omitempty" codec:"open_id_buyer,omitempty"`           // pin
	XidBuyer           string    `json:"xid_buyer,omitempty" codec:"xid_buyer,omitempty"`                   // pin
}

type SkuInfo struct {
	SkuName  string  `json:"skuName"`          // 商品名称
	Number   int     `json:"number"`           // 数量
	SkuPrice float64 `json:"skuPrice"`         // 价格
	SkuId    string  `json:"skuId"`            // SKUID
	WareId   string  `json:"wareId,omitempty"` // WareID
}

type OrderInfo struct {
	OrderId             string           `json:"orderId,omitempty" codec:"orderId,omitempty"`                         // 订单id
	VenderId            string           `json:"venderId,omitempty" codec:"venderId,omitempty"`                       // 商家id
	OrderType           string           `json:"orderType,omitempty" codec:"orderType,omitempty"`                     // 	订单类型（22 SOP； 75 LOC） 可选字段，需要在输入参数optional_fields中写入才能返回
	PayType             string           `json:"payType,omitempty" codec:"payType,omitempty"`                         // 支付方式（1货到付款, 2邮局汇款, 3自提, 4在线支付, 5公司转账, 6银行卡转账）
	OrderTotalPrice     string           `json:"orderTotalPrice,omitempty" codec:"orderTotalPrice,omitempty"`         // 订单总金额。总金额=订单金额（不减优惠，不加运费服务费税费）
	OrderSellerPrice    string           `json:"orderSellerPrice,omitempty" codec:"orderSellerPrice,omitempty"`       // 订单货款金额（订单总金额-商家优惠金额）
	OrderPayment        string           `json:"orderPayment,omitempty" codec:"orderPayment,omitempty"`               // 用户应付金额。应付款=货款-用户优惠-余额+运费+税费+服务费
	FreightPrice        string           `json:"freightPrice,omitempty" codec:"freightPrice,omitempty"`               // 商品的运费
	SellerDiscount      string           `json:"sellerDiscount,omitempty" codec:"sellerDiscount,omitempty"`           // 商家优惠金额
	OrderState          string           `json:"orderState,omitempty" codec:"orderState,omitempty"`                   // 	1）WAIT_SELLER_STOCK_OUT 等待出库 2）WAIT_GOODS_RECEIVE_CONFIRM 等待确认收货 3）WAIT_SELLER_DELIVERY等待发货（只适用于海外购商家，含义为'等待境内发货'标签下的订单,非海外购商家无需使用） 4) POP_ORDER_PAUSE POP暂停 5）FINISHED_L 完成 6）TRADE_CANCELED 取消 7）LOCKED 已锁定 8）WAIT_SEND_CODE 等待发码（LOC订单特有状态）
	OrderStateRemark    string           `json:"orderStateRemark,omitempty" codec:"orderStateRemark,omitempty"`       // 订单状态说明（中文）
	DeliveryType        string           `json:"deliveryType,omitempty" codec:"deliveryType,omitempty"`               // 送货（日期）类型（1-只工作日送货(双休日、假日不用送);2-只双休日、假日送货(工作日不用送);3-工作日、双休日与假日均可送货;其他值-返回“任意时间”）
	InvoiceEasyInfo     *InvoiceEasyInfo `json:"invoiceEasyInfo,omitempty" codec:"invoiceEasyInfo,omitempty"`         // 	发票组装信息。电子发票从这里获取。可替代invoiceInfo
	InvoiceInfo         string           `json:"invoiceInfo,omitempty" codec:"invoiceInfo,omitempty"`                 // 	发票信息。电子发票请从invoiceEasyInfo里取 “invoice_info: 不需要开具发票”下无需开具发票；其它返回值请正常开具发票
	InvoiceCode         string           `json:"invoiceCode,omitempty" codec:"invoiceCode,omitempty"`                 // 普通发票纳税人识别码
	OrderRemark         string           `json:"orderRemark,omitempty" codec:"orderRemark,omitempty"`                 // 	买家下单时订单备注
	OrderStartTime      string           `json:"orderStartTime,omitempty" codec:"orderStartTime,omitempty"`           // 下单时间
	OrderEndTime        string           `json:"orderEndTime,omitempty" codec:"orderEndTime,omitempty"`               // 	结单时间 如返回信息为“0001-01-01 00:00:00”和“1970-01-01 00:00:00”，可认为此订单为未完成状态。
	ConsigneeInfo       *ConsigneeInfo   `json:"consigneeInfo,omitempty" codec:"consigneeInfo,omitempty"`             // 收货人基本信息
	ItemInfoList        []*ItemInfo      `json:"itemInfoList,omitempty" codec:"itemInfoList,omitempty"`               // 商品详细信息
	CouponDetailList    []*CouponDetail  `json:"couponDetailList,omitempty" codec:"couponDetailList,omitempty"`       // 优惠详细信息
	VenderRemark        string           `json:"venderRemark,omitempty" codec:"venderRemark,omitempty"`               // 商家订单备注（不大于500字符） 可选字段，需要在输入参数optional_fields中写入才能返回
	BalanceUsed         string           `json:"balanceUsed,omitempty" codec:"balanceUsed,omitempty"`                 // 余额支付金额 可选字段，需要在输入参数optional_fields中写入才能返回
	Pin                 string           `json:"pin,omitempty" codec:"pin,omitempty"`                                 // 买家的账号信息 可选字段，需要在输入参数optional_fields中写入才能返回
	ReturnOrder         string           `json:"returnOrder,omitempty" codec:"returnOrder,omitempty"`                 // 售后订单标记 0:不是换货订单 1返修发货,直接赔偿,客服补件 2售后调货 可选字段，需要在输入参数optional_fields中写入才能返回
	PaymentConfirmTime  string           `json:"paymentConfirmTime,omitempty" codec:"paymentConfirmTime,omitempty"`   // 付款确认时间 如果没有付款时间 默认返回0001-01-01 00:00:00 可选字段，需要在输入参数optional_fields中写入才能返回
	Waybill             string           `json:"waybill,omitempty" codec:"waybill,omitempty"`                         // 运单号(当厂家自送时运单号可为空，不同物流公司的运单号用|分隔，如果同一物流公司有多个运单号，则用英文逗号分隔) 可选字段，需要在输入参数optional_fields中写入才能返回
	LogisticsId         string           `json:"logisticsId,omitempty" codec:"logisticsId,omitempty"`                 // 物流公司ID 可选字段，需要在输入参数optional_fields中写入才能返回
	VatInfo             *VatInfo         `json:"vatInfo,omitempty" codec:"vatInfo,omitempty"`                         // 增值税发票 可选字段，需要在输入参数optional_fields中写入才能返回
	Modified            string           `json:"modified,omitempty" codec:"modified,omitempty"`                       // 订单更新时间
	DirectParentOrderId string           `json:"directParentOrderId,omitempty" codec:"directParentOrderId,omitempty"` // 	直接父订单号 可选字段，需要在输入参数optional_fields中写入才能返回
	ParentOrderId       string           `json:"parentOrderId,omitempty" codec:"parentOrderId,omitempty"`             // 	根父订单号 可选字段，需要在输入参数optional_fields中写入才能返回
	Customs             string           `json:"customs,omitempty" codec:"customs,omitempty"`                         // 保税区信息
	CustomsModel        string           `json:"customsModel,omitempty" codec:"customsModel,omitempty"`               // 保税模型：直邮，保税集货，保税备货
	OrderSource         string           `json:"orderSource,omitempty" codec:"orderSource,omitempty"`                 // 订单来源。如：移动端订单
	StoreOrder          string           `json:"storeOrder,omitempty" codec:"storeOrder,omitempty"`                   // 京仓订单/云仓订单/空值“”
	IdSopShipmenttype   int              `json:"idSopShipmenttype,omitempty" codec:"idSopShipmenttype,omitempty"`     // 是否京配。68=京配，69=京配自提。
	ScDT                string           `json:"scDT,omitempty" codec:"scDT,omitempty"`                               // 最早生产时间。预约出库时间
	ServiceFee          string           `json:"serviceFee,omitempty" codec:"serviceFee,omitempty"`                   // 服务费。目前没有此值
	//PauseBizInfo        string           `json:"pauseBizInfo,omitempty" codec:"pauseBizInfo,omitempty"` // 订单暂停业务数据
	TaxFee      string `json:"taxFee,omitempty" codec:"taxFee,omitempty"`           // 税价分离 税费
	TuiHuoWuYou string `json:"tuiHuoWuYou,omitempty" codec:"tuiHuoWuYou,omitempty"` // 退货无忧
	OrderSign   string `json:"orderSign,omitempty" codec:"orderSign,omitempty"`     // sendpay
	StoreId     string `json:"storeId,omitempty" codec:"storeId,omitempty"`         // 	仓库Id
}

type InvoiceEasyInfo struct {
	InvoiceType                  string `json:"invoiceType,omitempty" codec:"invoiceType,omitempty"`                                     // 发票类型，0=不开发票、1=普通发票、2=增值税发票、3=电子发票
	InvoiceTitle                 string `json:"invoiceTitle,omitempty" codec:"invoiceTitle,omitempty"`                                   // 发票抬头。个人/公司名称
	InvoiceContentId             string `json:"invoiceContentId,omitempty" codec:"invoiceContentId,omitempty"`                           // 	发票内容
	InvoiceConsigneeEmail        string `json:"invoiceConsigneeEmail,omitempty" codec:"invoiceConsigneeEmail,omitempty"`                 // 电子发票联系人邮箱
	InvoiceConsigneePhone        string `json:"invoiceConsigneePhone,omitempty" codec:"invoiceConsigneePhone,omitempty"`                 // 电子发票联系人手机号
	InvoiceCode                  string `json:"invoiceCode,omitempty" codec:"invoiceCode,omitempty"`                                     // 发票纳税人识别号
	EncryptInvoiceTitle          string `json:"encrypt_invoiceTitle,omitempty" codec:"encrypt_invoiceTitle,omitempty"`                   // 发票抬头。个人/公司名称
	EncryptInvoiceConsigneeEmail string `json:"encrypt_invoiceConsigneeEmail,omitempty" codec:"encrypt_invoiceConsigneeEmail,omitempty"` // 电子发票联系人邮箱
	EncryptInvoiceConsigneePhone string `json:"encrypt_invoiceConsigneePhone,omitempty" codec:"invoiceConsigneePhone,omitempty"`         // 电子发票联系人手机号
}

type ConsigneeInfo struct {
	Fullname           string `json:"fullname,omitempty" codec:"fullname,omitempty"`       // 姓名
	Telephone          string `json:"telephone,omitempty" codec:"telephone,omitempty"`     // 固定电话
	Mobile             string `json:"mobile,omitempty" codec:"mobile,omitempty"`           // 手机
	FullAddress        string `json:"fullAddress,omitempty" codec:"fullAddress,omitempty"` // 地址
	Province           string `json:"province,omitempty" codec:"province,omitempty"`
	City               string `json:"city,omitempty" codec:"city,omitempty"`
	County             string `json:"county,omitempty" codec:"county,omitempty"`
	Town               string `json:"town,omitempty" codec:"town,omitempty"`
	ProvinceId         string `json:"provinceId,omitempty" codec:"provinceId,omitempty"`
	CityId             string `json:"cityId,omitempty" codec:"cityId,omitempty"`
	CountyId           string `json:"countyId,omitempty" codec:"countyId,omitempty"`
	TownId             string `json:"townId,omitempty" codec:"townId,omitempty"`
	EncryptFullname    string `json:"encrypt_fullname,omitempty" codec:"encrypt_fullname,omitempty"`       // 姓名
	EncryptTelephone   string `json:"encrypt_telephone,omitempty" codec:"encrypt_telephone,omitempty"`     // 固定电话
	EncryptMobile      string `json:"encrypt_mobile,omitempty" codec:"encrypt_mobile,omitempty"`           // 手机
	EncryptFullAddress string `json:"encrypt_fullAddress,omitempty" codec:"encrypt_fullAddress,omitempty"` // 地址
}

type ItemInfo struct {
	SkuId       string `json:"skuId,omitempty" codec:"skuId,omitempty"`             // 京东内部SKU的ID
	OuterSkuId  string `json:"outerSkuId,omitempty" codec:"outerSkuId,omitempty"`   // SKU外部ID（极端情况下不保证返回，建议从商品接口获取
	SkuName     string `json:"skuName,omitempty" codec:"skuName,omitempty"`         // 商品的名称+SKU规格
	JdPrice     string `json:"jdPrice,omitempty" codec:"jdPrice,omitempty"`         // SKU的京东价
	GiftPoint   string `json:"giftPoint,omitempty" codec:"giftPoint,omitempty"`     // 赠送积分
	WareId      string `json:"wareId,omitempty" codec:"wareId,omitempty"`           // 京东内部商品ID（极端情况下不保证返回，建议从商品接口获取
	ItemTotal   string `json:"itemTotal,omitempty" codec:"itemTotal,omitempty"`     // 	数量
	ProductNo   string `json:"productNo,omitempty" codec:"productNo,omitempty"`     // 货号
	ServiceName string `json:"serviceName,omitempty" codec:"serviceName,omitempty"` // 服务项名称
	NewStoreId  string `json:"newStoreId,omitempty" codec:"newStoreId,omitempty"`   // 	item维度的仓库id
}

type OrderItemInfo struct {
	ItemInfo

	OrderId string `json:"orderId,omitempty" codec:"orderId,omitempty"`
}

type CouponDetail struct {
	OrderId     string `json:"orderId,omitempty" codec:"orderId,omitempty"`         // 订单编号
	SkuId       string `json:"skuId,omitempty" codec:"skuId,omitempty"`             // 	京东sku编号。(只有30-单品促销优惠 才有skuId，其他时值为””)
	CouponType  string `json:"couponType,omitempty" codec:"couponType,omitempty"`   // 优惠类型: 20-套装优惠, 28-闪团优惠, 29-团购优惠, 30-单品促销优惠, 34-手机红包, 35-满返满送(返现), 39-京豆优惠,41-京东券优惠, 52-礼品卡优惠,100-店铺优惠
	CouponPrice string `json:"couponPrice,omitempty" codec:"couponPrice,omitempty"` // 	优惠金额。详细的sku优惠/分摊金额 请联系jos找ofc提供对应接口。
}

type VatInfo struct {
	VatNo              string `json:"vatNo,omitempty" codec:"vatNo,omitempty"`                             // 纳税人识别号
	AddressRegIstered  string `json:"addressRegIstered,omitempty" codec:"addressRegIstered,omitempty"`     // 	注册地址
	PhoneRegIstered    string `json:"phoneRegIstered,omitempty" codec:"phoneRegIstered,omitempty"`         // 注册电话
	DepositBank        string `json:"depositBank,omitempty" codec:"depositBank,omitempty"`                 // 	开户银行
	BankAccount        string `json:"bankAccount,omitempty" codec:"bankAccount,omitempty"`                 // 	银行账户
	UserAddress        string `json:"userAddress,omitempty" codec:"userAddress,omitempty"`                 // 收票人地址
	UserName           string `json:"userName,omitempty" codec:"userName,omitempty"`                       // 	收票人姓名
	UserPhone          string `json:"userPhone,omitempty" codec:"userPhone,omitempty"`                     // 	收票人电话
	EncryptBankAccount string `json:"encrypt_bankAccount,omitempty" codec:"encrypt_bankAccount,omitempty"` //    银行账户
	EncryptUserAddress string `json:"encrypt_userAddress,omitempty" codec:"encrypt_userAddress,omitempty"` // 收票人地址
	EncryptUserName    string `json:"encrypt_userName,omitempty" codec:"encrypt_userName,omitempty"`       //    收票人姓名
	EncryptUserPhone   string `json:"encrypt_userPhone,omitempty" codec:"encrypt_userPhone,omitempty"`     //    收票人电话
}
