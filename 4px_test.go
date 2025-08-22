package fourpx_express

import (
	"fmt"
	"github.com/zengweigg/fourpx-express/config"
	"github.com/zengweigg/fourpx-express/model"
	"testing"
)

func Test001(m *testing.T) {
	// 初始化
	cfg := config.LoadConfig()
	fourClient := NewFourService(*cfg)
	// 构造测试请求数据
	// postData := model.OrderLabelPost{
	// 	RequestNo: "4PX3001957565250CN",
	// }
	// // 获取标签
	// resp, err := fourClient.Services.Order.GetLabelList(postData)
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return
	// }

	postData := model.CancelOrderPost{
		RequestNo:    "3000210210",
		CancelReason: "TEST",
	}
	// 取消单号
	resp, err := fourClient.Services.Order.CancelOrder(postData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(resp.Data)
}
