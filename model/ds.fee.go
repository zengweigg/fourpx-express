package model

// 查询订单费用信息  请求参数
type GetFreightPost struct {
	RequestNo string `json:"request_no" maxLength:"32"` // 请求单号
}

type FreightData struct {
	Data               string           `json:"data"`                 // 成功的数据
	FourPxTrackingNo   string           `json:"4px_tracking_no"`      // 4PX单号
	LogisticsChannelNo string           `json:"logistics_channel_no"` // 服务商单号
	RefNo              string           `json:"ref_no"`               // 客户单号
	TotalFee           string           `json:"total_fee"`            // 总费用
	ChargeWeight       string           `json:"charge_weight"`        // 计费重
	Currency           string           `json:"currency"`             // 币种
	CreateTime         string           `json:"create_time"`          // 计费时间
	Subs               []FreightSubItem `json:"subs"`                 // 费用项列表
}

type FreightSubItem struct {
	Currency  string `json:"currency"`   // 币种
	FeeAmount string `json:"fee_amount"` // 费用金额
	FeeName   string `json:"fee_name"`   // 费用名称
}

type FreightResp struct {
	Response
	Data FreightData `json:"data,omitempty"`
}
