package model

import "gopkg.in/guregu/null.v4"

// CreateOrderPost 1.直发下单 请求参数
type CreateOrderPost struct {
	RefNo                    string           `json:"ref_no"`                               // 参考号（客户自有系统的单号，如客户单号）
	BusinessType             string           `json:"business_type"`                        // 业务类型(4PX内部调度所需，如需对接传值将说明，默认值：BDS。)
	DutyType                 string           `json:"duty_type"`                            // 税费费用承担方式(可选值：U、P); U：DDU由收件人支付关税; P：DDP 由寄件方支付关税 （如果物流产品只提供其中一种，则以4PX提供的为准）
	IsInsure                 string           `json:"is_insure"`                            // 是否投保(Y、N)
	LogisticsServiceInfo     LogisticsService `json:"logistics_service_info"`               // 物流服务信息
	ReturnInfo               ReturnInfo       `json:"return_info"`                          // 退件信息
	ParcelList               []Parcel         `json:"parcel_list"`                          // 包裹列表
	InsuranceInfo            Insurance        `json:"insurance_info"`                       // 保险信息
	Sender                   Sender           `json:"sender"`                               // 发件人信息
	RecipientInfo            Recipient        `json:"recipient_info"`                       // 收件人信息
	DeliverTypeInfo          DeliverType      `json:"deliver_type_info"`                    // 到仓方式
	DeliverToRecipientInfo   DeliveryInfo     `json:"deliver_to_recipient_info"`            // 投递信息
	TrackingNo               string           `json:"4px_tracking_no,omitempty"`            // 4PX跟踪号（预分配号段的客户可传此值）
	LabelBarcode             string           `json:"label_barcode,omitempty"`              // 面单条码（预分配号段的客户可传此值）
	CargoType                string           `json:"cargo_type,omitempty"`                 // 货物类型（1：礼品;2：文件;3：商品货样;5：其它；默认值：5）
	VatNo                    string           `json:"vat_no,omitempty"`                     // VAT税号(数字或字母)；欧盟国家(含英国)使用的增值税号；
	Mid                      string           `json:"mid,omitempty"`                        // MID码是美国进口清关所需的信息之一。用于制造商、托运人或出口商全称及地址的替代信息，并且对于美国正式海关申报是必不可少的。在进行美国进口清关时，商业发票上必须显示MID码
	Csp                      string           `json:"csp,omitempty"`                        // CSP是瑞士进口税结算账户
	EoriNo                   string           `json:"eori_no,omitempty"`                    // EORI号码(数字或字母)；欧盟入关时需要EORI号码，用于商品货物的清关
	IossNo                   string           `json:"ioss_no,omitempty"`                    // IOSS号码
	BuyerID                  string           `json:"buyer_id,omitempty"`                   // 买家ID(数字或字母)
	SalesPlatform            string           `json:"sales_platform,omitempty"`             // 销售平台（点击查看详情）
	TradeID                  string           `json:"trade_id,omitempty"`                   // 交易号ID(数字或字母)
	SellerID                 string           `json:"seller_id,omitempty"`                  // 卖家ID(数字或字母)
	IsCommercialInvoice      string           `json:"is_commercial_invoice,omitempty"`      // 能否提供商业发票（Y/N） Y：能提供商业发票(则系统不会生成形式发票)；N：不能提供商业发票(则系统会生成形式发票)； 默认为N；DHL产品必填，如产品代码A1/A5；
	ParcelQty                int              `json:"parcel_qty,omitempty"`                 // 包裹件数（一个订单有多少件包裹，就填写多少件数，请如实填写包裹件数，否则DHL无法返回准确的子单号数和子单号标签；DHL产品必填，如产品代码A1/A5；）
	FreightCharges           float64          `json:"freight_charges,omitempty"`            // 运费(客户填写自己估算的运输费用；支持的币种，根据物流产品+收件人国家配置)
	CurrencyFreight          string           `json:"currency_freight,omitempty"`           // 运费币种(按照ISO标准三字码；支持的币种，根据物流产品+收件人国家配置)
	DeclareInsurance         float64          `json:"declare_insurance,omitempty"`          // 申报保险费（是否必填，根据物流产品+目的国配置；根据欧盟IOSS政策，货值/运费/保险费可单独申报）支持小数点后2位
	CurrencyDeclareInsurance string           `json:"currency_declare_insurance,omitempty"` // 申报保险费币种（按照ISO标准，币种需和进出口国申报币种一致）
	Ext                      string           `json:"ext,omitempty"`                        // 扩展字段
	SortCode                 string           `json:"sort_code,omitempty"`                  // 分拣分区
	OrderAttachmentInfo      []Attachment     `json:"order_attachment_info,omitempty"`      // 订单附件列表
	PaymentInfo              PaymentInfo      `json:"payment_info,omitempty"`               // 付款信息
}

// 物流服务信息
type LogisticsService struct {
	LogisticsProductCode string `json:"logistics_product_code"`         // 物流产品代码(点击查看详情)
	CustomsService       string `json:"customs_service,omitempty"`      // 单独报关（Y：单独报关；N：不单独报关） 默认值：N
	SignatureService     string `json:"signature_service,omitempty"`    // 签名服务（Y/N)；默认值：N
	ValueAddedServices   string `json:"value_added_services,omitempty"` // 其他服务（待完善)
}

// 退件信息
type ReturnInfo struct {
	IsReturnOnDomestic string   `json:"is_return_on_domestic"`          // 境内/国内异常处理策略(Y：退件--实际是否支持退件，以及退件策略、费用，参考报价表；N：销毁；U：其他--等待客户指令) 默认值：N；
	IsReturnOnOversea  string   `json:"is_return_on_oversea"`           // 境外/国外异常处理策略(Y：退件--实际是否支持退件，以及退件策略、费用，参考报价表；N：销毁；U：其他--等待客户指令) 默认值：N；
	DomesticReturnAddr *Address `json:"domestic_return_addr,omitempty"` // 国内退件地址
	OverseaReturnAddr  *Address `json:"oversea_return_addr,omitempty"`  // 海外退件地址
}

// 包裹列表
type Parcel struct {
	Weight             float64          `json:"weight"`                 // 预报重量（g）
	Length             float64          `json:"length,omitempty"`       // 包裹长（cm）
	Width              float64          `json:"width,omitempty"`        // 包裹宽（cm）
	Height             float64          `json:"height,omitempty"`       // 包裹高（cm）
	ParcelValue        float64          `json:"parcel_value"`           // 包裹申报价值（最多2位小数）
	Currency           string           `json:"currency"`               // 包裹申报价值币别（按照ISO标准三字码；支持的币种，根据物流产品+收件人国家配置；币种需和进出口国申报币种一致）
	IncludeBattery     string           `json:"include_battery"`        // 是否含电池（Y/N）
	BatteryType        string           `json:"battery_type,omitempty"` // 带电类型(具体产品可支持的带电类型请咨询销售）
	ProductList        []Product        `json:"product_list,omitempty"` // 投保物品信息
	DeclareProductInfo []DeclareProduct `json:"declare_product_info"`   // 海关申报信息
}

// 投保物品信息
type Product struct {
	SkuCode                string  `json:"sku_code,omitempty"`                 // 投保SKU（客户自定义SKUcode）（数字或字母或空格）
	StandardProductBarcode string  `json:"standard_product_barcode,omitempty"` // 投保商品标准条码（UPC、EAN、JAN…）
	ProductName            string  `json:"product_name,omitempty"`             // 投保商品名称
	ProductDescription     string  `json:"product_description,omitempty"`      // 投保商品描述
	ProductUnitPrice       float64 `json:"product_unit_price,omitempty"`       // 投保商品单价（按对应币别的法定单位，最多2位小数点）
	Currency               string  `json:"currency,omitempty"`                 // 投保商品单价币别（按照ISO标准三字码；支持的币种，根据物流产品+收件人国家配置；币种需和进出口国申报币种一致）
	Qty                    int     `json:"qty,omitempty"`                      // 投保商品数量（单位为pcs）
}

// 海关申报信息
type DeclareProduct struct {
	DeclareProductCode     string  `json:"declare_product_code,omitempty"`    // 申报产品代码（在4PX已备案申报产品的代码）
	DeclareProductNameCn   string  `json:"declare_product_name_cn,omitempty"` // 申报品名(当地语言)
	DeclareProductNameEn   string  `json:"declare_product_name_en,omitempty"` // 申报品名（英语）
	Uses                   string  `json:"uses,omitempty"`                    // 用途
	Specification          string  `json:"specification,omitempty"`           // 规格
	Component              string  `json:"component,omitempty"`               // 成分
	UnitNetWeight          int     `json:"unit_net_weight,omitempty"`         // 单件商品净重（默认以g为单位）
	UnitGrossWeight        int     `json:"unit_gross_weight,omitempty"`       // 单件商品毛重（默认以g为单位）
	Material               string  `json:"material,omitempty"`                // 材质
	DeclareProductCodeQty  int     `json:"declare_product_code_qty"`          // 申报数量
	UnitDeclareProduct     string  `json:"unit_declare_product,omitempty"`    // 单位（点击查看详情；默认值：PCS）
	OriginCountry          string  `json:"origin_country,omitempty"`          // 原产地（ISO标准2字码）点击查看详情
	CountryExport          string  `json:"country_export,omitempty"`          // 出口国/起始国/发件人国家（ISO标准2字码）
	CountryImport          string  `json:"country_import,omitempty"`          // 进口国/目的国/收件人国家（ISO标准2字码）
	HscodeExport           string  `json:"hscode_export,omitempty"`           // 出口国/起始国/发件人国家_海关编码(只支持数字)
	HscodeImport           string  `json:"hscode_import,omitempty"`           // 进口国/目的国/收件人国家_海关编码(只支持数字)
	DeclareUnitPriceExport float64 `json:"declare_unit_price_export"`         // 出口国/起始国/发件人国家_申报单价（按对应币别的法定单位，最多2位小数点） 必填
	CurrencyExport         string  `json:"currency_export"`                   // 出口国/起始国/发件人国家_申报单价币种（按照ISO标准；支持的币种，根据物流产品+收件人国家配置；币种需和进口国申报币种一致）
	DeclareUnitPriceImport float64 `json:"declare_unit_price_import"`         // 进口国/目的国/收件人国家_申报单价（按对应币别的法定单位，最多2位小数点）
	CurrencyImport         string  `json:"currency_import"`                   // 进口国/目的国/收件人国家_申报单价币种（按照ISO标准；支持的币种，根据物流产品+收件人国家配置；币种需和出口国申报币种一致）
	BrandExport            string  `json:"brand_export"`                      // 出口国/起始国/发件人国家_品牌(必填；若无，填none即可)
	BrandImport            string  `json:"brand_import"`                      // 进口国/目的国/收件人国家_品牌(必填；若无，填none即可)
	SalesUrl               string  `json:"sales_url,omitempty"`               // 商品销售URL
	PackageRemarks         string  `json:"package_remarks,omitempty"`         // 配货字段（打印标签选择显示配货信息是将会显示：package_remarks*qty）
	ProductImageUrl        string  `json:"product_image_url,omitempty"`       // 该SKU的图片对应的url
}

// 保险信息
type Insurance struct {
	InsureType        string  `json:"insure_type,omitempty"`         // 保险类型（XY:4PX保价；XP:第三方保险） 5Y, 5元每票 8Y, 8元每票 6P, 0.6%保费
	InsureValue       float64 `json:"insure_value,omitempty"`        // 保险价值
	Currency          string  `json:"currency,omitempty"`            // 保险币别（按照ISO标准，目前只支持USD）
	InsurePerson      string  `json:"insure_person,omitempty"`       // 投保人/公司
	CertificateType   string  `json:"certificate_type,omitempty"`    // 投保人证件类型（暂时只支持身份证，类型为：ID）
	CertificateNo     string  `json:"certificate_no,omitempty"`      // 投保人证件号码
	CategoryCode      string  `json:"category_code,omitempty"`       // 保险类目ID（保险的类目，暂时不填，默认取第一个类目）
	InsureProductName string  `json:"insure_product_name,omitempty"` // 投保货物名称
	PackageQty        string  `json:"package_qty,omitempty"`         // 投保包装及数量
}

// 发件人信息
type Sender struct {
	FirstName       string       `json:"first_name"`                 // 名/姓名
	LastName        string       `json:"last_name,omitempty"`        // 姓
	Company         string       `json:"company,omitempty"`          // 公司名
	Phone           string       `json:"phone,omitempty"`            // 电话（必填）
	Phone2          string       `json:"phone2,omitempty"`           // 电话2
	Email           string       `json:"email,omitempty"`            // 邮箱
	PostCode        string       `json:"post_code,omitempty"`        // 邮编
	Country         string       `json:"country"`                    // 国家（国际二字码 标准ISO 3166-2 ）
	State           string       `json:"state,omitempty"`            // 州/省
	City            string       `json:"city"`                       // 城市
	District        string       `json:"district,omitempty"`         // 区、县
	Street          string       `json:"street,omitempty"`           // 街道/详细地址
	HouseNumber     string       `json:"house_number,omitempty"`     // 门牌号
	CertificateInfo *Certificate `json:"certificate_info,omitempty"` // 证件信息
}

// 收件人信息
type Recipient struct {
	FirstName       string       `json:"first_name"`                 // 名/姓名
	LastName        string       `json:"last_name,omitempty"`        // 姓
	Company         string       `json:"company,omitempty"`          // 公司名
	Phone           string       `json:"phone"`                      // 电话（必填）
	Phone2          string       `json:"phone2,omitempty"`           // 电话2
	Email           string       `json:"email,omitempty"`            // 邮箱
	PostCode        string       `json:"post_code,omitempty"`        // 邮编,非必填（部分产品需要填，具体以返回提示为准）
	Country         string       `json:"country"`                    // 国家（国际二字码 标准ISO 3166-2 ）
	State           string       `json:"state,omitempty"`            // 州/省
	City            string       `json:"city"`                       // 城市
	District        string       `json:"district,omitempty"`         // 区、县（可对应为adress 2）
	Street          string       `json:"street"`                     // 街道/详细地址（可对应为adress 1）
	HouseNumber     string       `json:"house_number,omitempty"`     // 门牌号
	SecondName      string       `json:"second_name,omitempty"`      // 非必填，备用名字，一般用于有两个名字的国家，比如日本清关要求必需片假名
	CertificateInfo *Certificate `json:"certificate_info,omitempty"` // 证件信息
}

// 证件信息
type Certificate struct {
	IDType     string `json:"id_type,omitempty"`      // 证件类型（点击查看详情）
	IDNo       string `json:"id_no,omitempty"`        // 证件号
	IDFrontURL string `json:"id_front_url,omitempty"` // 证件正面照URL
	IDBackURL  string `json:"id_back_url,omitempty"`  // 证件背面照URL
}

// 到仓方式
type DeliverType struct {
	DeliverType       string        `json:"deliver_type"`                    // 到仓方式（1:上门揽收；2:快递到仓；3:自送到仓；5:自送门店）
	WarehouseCode     string        `json:"warehouse_code,omitempty"`        // 收货仓库/门店代码（仓库代码）
	PickUpInfo        *PickUpInfo   `json:"pick_up_info,omitempty"`          // 揽收信息
	ExpressTo4pxInfo  *ExpressInfo  `json:"express_to_4px_info,omitempty"`   // 快递到仓信息
	SelfSendTo4pxInfo *SelfSendInfo `json:"self_send_to_4px_info,omitempty"` // 自己送仓信息
}

// 揽收信息
type PickUpInfo struct {
	ExpectPickUpEarliestTime int64    `json:"expect_pick_up_earliest_time,omitempty"` // 期望提货最早时间（*注：时间格式的传入值需要转换为long类型格式。）
	ExpectPickUpLatestTime   int64    `json:"expect_pick_up_latest_time,omitempty"`   // 期望提货最迟时间（*注：时间格式的传入值需要转换为long类型格式。）
	PickUpAddressInfo        *Address `json:"pick_up_address_info,omitempty"`         // 收货地址
}

type ExpressInfo struct {
	ExpressCompany string `json:"express_company,omitempty"` // 快递公司
	TrackingNo     string `json:"tracking_no,omitempty"`     // 追踪号
}

type SelfSendInfo struct {
	BookingEarliestTime int64 `json:"booking_earliest_time,omitempty"` // 预约送仓最早时间（*注：时间格式的传入值需要转换为long类型格式。）
	BookingLatestTime   int64 `json:"booking_latest_time,omitempty"`   // 预约送仓最晚时间（*注：时间格式的传入值需要转换为long类型格式。）
}

// 投递信息
type DeliveryInfo struct {
	DeliverType string `json:"deliver_type,omitempty"` // 投递类型：HOME_DELIVERY-投递到门；SELF_PICKUP_STATION-投递门店（自提点）；SELF_SERVICE_STATION-投递自提柜(自助点）；默认：HOME_DELIVERY；注：目前暂时不支持投递门店、投递自提柜
	StationCode string `json:"station_code,omitempty"` // 自提门店/自提点的信息(选择自提时必传，点击获取详情)
}

type Address struct {
	FirstName   string `json:"first_name"`             // 名/姓名
	Phone       string `json:"phone"`                  // 电话（必填）
	PostCode    string `json:"post_code"`              // 邮编
	Country     string `json:"country"`                // 国家（国际二字码 标准ISO 3166-2 ）
	City        string `json:"city"`                   // 城市
	Street      string `json:"street"`                 // 街道/详细地址
	LastName    string `json:"last_name,omitempty"`    // 姓
	Company     string `json:"company,omitempty"`      // 公司名
	Phone2      string `json:"phone2,omitempty"`       // 电话2
	Email       string `json:"email,omitempty"`        // 邮箱
	State       string `json:"state,omitempty"`        // 州/省
	District    string `json:"district,omitempty"`     // 区、县
	HouseNumber string `json:"house_number,omitempty"` // 门牌号
}

// 订单附件列表
type Attachment struct {
	FileData         string `json:"file_data,omitempty"`         // base64文件数据
	FileURL          string `json:"file_url,omitempty"`          // 文件url
	FileType         string `json:"file_type,omitempty"`         // 文件类型
	AttachmentFormat string `json:"attachment_format,omitempty"` // 附件格式
	AttachmentType   string `json:"attachment_type,omitempty"`   // 附件类型
}

// 付款信息
type PaymentInfo struct {
	PaymentMethod               string `json:"payment_method,omitempty"`                  // 包裹的付款方式（信用卡”“借记卡”等）
	IssuingEntity               string `json:"issuing_entity,omitempty"`                  // 支付渠道（如Visa、万事达）
	LastFourDigitsPaymentMethod string `json:"last_four_digits_payment_method,omitempty"` // 付款银行卡后四位
	BirthDate                   string `json:"birth_date,omitempty"`                      // 出生年月日，示例：DDMMYYYY
}

// OrderData 1.下单响应参数
type OrderData struct {
	DsConsignmentNo    string `json:"ds_consignment_no"`    // 直发委托单号
	TrackingNo         string `json:"4px_tracking_no"`      // 4PX跟踪号
	LabelBarcode       string `json:"label_barcode"`        // 标签条码号。*注：参数为deprecated状态
	RefNo              string `json:"ref_no"`               // 客户单号/客户参考号
	LogisticsChannelNo string `json:"logistics_channel_no"` // 物流渠道号码。如果结果返回为空字符，表示暂时没有物流渠道号码，请稍后主动调用查询直发委托单接口查询
	OdaResultSign      string `json:"oda_result_sign"`      // ODA标识(偏远地址：Y 非偏远地址：N)
}

type OrderResp struct {
	Response
	Data OrderData `json:"data,omitempty"` // 成功的数据
}

// 2 取消下单
type CancelOrderPost struct {
	Currency           string `json:"currency" maxLength:"10"`                      // 包裹申报价值币别（按照ISO标准三字码；支持的币种，根据物流产品+收件人国家配置；币种需和进出口国申报币种一致）
	CurrencyDeprecated string `json:"currency_deprecated,omitempty" maxLength:"10"` // [@即将废弃]投保商品单价币别（按照ISO标准三字码；支持的币种，根据物流产品+收件人国家配置；币种需和进出口国申报币种一致）
	DeliverType        string `json:"deliver_type" maxLength:"32"`                  // 到仓方式（1:上门揽收；2:快递到仓；3:自送到仓；5:自送门店）
	RequestNo          string `json:"request_no" maxLength:"64"`                    // 请求单号
	CancelReason       string `json:"cancel_reason" maxLength:"128"`                // 取消原因
}

// 2.取消下单响应
type CancelOrderResp struct {
	Response
	Data null.String `json:"data,omitempty"`
}
