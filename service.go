package fourpx_express

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/zengweigg/fourpx-express/config"
	"io"
)

type service struct {
	config     *config.Config // Config
	logger     Logger         // Logger
	httpClient *resty.Client  // HTTP client
}

type services struct {
	Order       orderService
	Track       trackService
	Fee         feeService
	Appointment appointmentService
}

// MD5加密函数，返回小写结果
func MD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// GetSign 消息加密获取签名
func GetSign(param, method, apiKey, apiSecret, timestamp, v string) (sign string, err error) {
	if param == "" || method == "" || apiKey == "" || apiSecret == "" || timestamp == "" {
		return "", errors.New("参数不能为空")
	}
	var str string
	str = fmt.Sprintf("app_key%sformatjsonmethod%stimestamp%sv%s%s%s", apiKey, method, timestamp, v, param, apiSecret)
	//  MD5加密
	sign = MD5(str)
	return sign, nil
}
