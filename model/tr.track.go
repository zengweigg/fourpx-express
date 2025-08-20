package model

// 物流轨迹 请求参数
type OrderTrackingPost struct {
	DeliveryOrderNo string `json:"deliveryOrderNo" maxLength:"50"` // 物流单号
}

type TrackingData struct {
	DeliveryOrderNo    string         `json:"deliveryOrderNo"`    // 返回传入的物流单号
	DestinationCountry string         `json:"destinationCountry"` // 目的地国家
	ServerName         string         `json:"serverName"`         // 服务商名称
	ServerNum          string         `json:"serverNum"`          // 服务商单号
	TrackingList       []TrackingInfo `json:"trackingList"`       // 物流轨迹信息集合
}

type TrackingInfo struct {
	BusinessLinkCode string `json:"businessLinkCode"` // 轨迹代码
	OccurDatetime    string `json:"occurDatetime"`    // 轨迹发生时间
	TrackingContent  string `json:"trackingContent"`  // 轨迹描述
	TimeZone         string `json:"timeZone"`         // 轨迹发生的时区
	OccurLocation    string `json:"occurLocation"`    // 轨迹发生地
	Country          string `json:"country"`          // 轨迹国家
	City             string `json:"city"`             // 轨迹城市
}

type TrackingResp struct {
	Response
	Data TrackingData `json:"data,omitempty"` // 成功的数据
}
