package fourpx_express

import (
	"github.com/bytedance/sonic"
	"github.com/zengweigg/fourpx-express/model"
)

type trackService service

// OrderTracking 物流轨迹查询
func (s trackService) OrderTracking(postData model.OrderTrackingPost) (model.TrackingResp, error) {
	respData := model.TrackingResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("tr.order.tracking.get")
	// fmt.Println(string(resp.Body()))
	if err != nil {
		return respData, err
	}
	// 解析数据
	err = sonic.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return respData, nil
	}
	return respData, nil
}
