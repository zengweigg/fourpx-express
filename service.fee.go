package fourpx_express

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/zengweigg/fourpx-express/model"
)

// 费用

type feeService service

// GetLabelList 查询订单费用信息
func (s feeService) GetFreightFee(postData model.GetFreightPost) (model.FreightResp, error) {
	respData := model.FreightResp{}
	// 请求数据
	resp, err := s.httpClient.R().
		SetBody(postData).
		Post("ds.xms.order.getFreight")
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
