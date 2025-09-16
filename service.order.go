package fourpx_express

import (
	"github.com/bytedance/sonic"
	"github.com/zengweigg/fourpx-express/model"
)

type orderService service

// CreateOrder 下单
func (s orderService) CreateOrder(postData model.CreateOrderPost) (model.OrderResp, error) {
	respData := model.OrderResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ds.xms.order.create")
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

// CancelOrder 取消下单
func (s orderService) CancelOrder(postData model.CancelOrderPost) (model.CancelOrderResp, error) {
	respData := model.CancelOrderResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ds.xms.order.cancel")
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

// GetLabelList 获取标签(面单)
func (s orderService) GetLabelList(postData model.OrderLabelPost) (model.OrderLabelResp, error) {
	respData := model.OrderLabelResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ds.xms.label.get")
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
