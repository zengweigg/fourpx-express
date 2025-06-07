package fourpx_express

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/zengweigg/fourpx-express/model"
)

type orderService service

// 获取标签
func (s orderService) GetLabelList(postData model.OrderLabelPost) (model.OrderLabelResp, error) {
	respData := model.OrderLabelResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ds.xms.label.get")
	fmt.Println(string(resp.Body()))
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
