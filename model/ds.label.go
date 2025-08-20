package model

// 3. 获取面单 请求参数
type OrderLabelPost struct {
	RequestNo              string `json:"request_no"`                          // 请求单号（支持4PX单号、客户单号和面单号）
	ResponseLabelFormat    string `json:"response_label_format,omitempty"`     // 返回面单的格式（PDF：返回PDF下载链接；IMG：返回IMG图片下载链接）默认为PDF
	IsPrintTime            string `json:"is_print_time,omitempty"`             // 是否打印当前时间（Y：打印；N：不打印）默认为N
	IsPrintBuyerID         string `json:"is_print_buyer_id,omitempty"`         // 是否打印买家ID（Y：打印；N：不打印）默认为N
	IsPrintPickInfo        string `json:"is_print_pick_info,omitempty"`        // 是否在标签上打印配货信息（Y：打印；N：不打印）默认为N
	IsPrintDeclarationList string `json:"is_print_declaration_list,omitempty"` // 是否打印报关单（Y：打印；N：不打印）默认为N
	IsPrintCustomerWeight  string `json:"is_print_customer_weight,omitempty"`  // 报关单上是否打印客户预报重（Y：打印；N：不打印）默认为N
	CreatePackageLabel     string `json:"create_package_label,omitempty"`      // 是否单独打印配货单（Y：打印；N：不打印）默认为N
	IsPrintPickBarcode     string `json:"is_print_pick_barcode,omitempty"`     // 配货单上是否打印配货条形码（Y：打印；N：不打印）默认为N
	IsPrintMerge           string `json:"is_print_merge,omitempty"`            // 是否合并打印(Y：合并；N：不合并)默认为N
	BarCodeOrderType       string `json:"bar_code_order_type,omitempty"`       // 指定条码单号类型（0：默认类型；1：4PX单号），默认为0
}

// 3.获取面单 请响应参数
type OrderLabelResp struct {
	Response
	Data OrderLabelData `json:"data,omitempty"` // 成功的数据
}

type LabelURLInfo struct {
	LogisticsLabel string `json:"logistics_label,omitempty"` // 面单链接(①普通客户返回4PX标准物流链接；②特定客户返回物流服务商标签链接)
	CustomLabel    string `json:"custom_label,omitempty"`    // 报关标签链接(特定客户且特定产品专用)
	PackageLabel   string `json:"package_label,omitempty"`   // 配货标签链接(特定客户且特定产品专用)
	InvoiceLabel   string `json:"invoice_label,omitempty"`   // DHL发票链接(特定客户且特定产品专用)
}

type OrderLabelData struct {
	LabelBarcode      string       `json:"label_barcode"`       // 面单条码(①普通客户返回面单号；②特定客户且特定产品返回物流服务商单号)
	ChildLabelBarcode []string     `json:"child_label_barcode"` // 子面单号
	LabelURLInfo      LabelURLInfo `json:"label_url_info,"`     // 标签链接信息
}
