package fourpx_express

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/gox/bytex"
	"github.com/zengweigg/fourpx-express/config"
	"github.com/zengweigg/fourpx-express/model"
	"net/http"
	"strconv"
	"time"
)

const (
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36"
)

type FourClient struct {
	config     *config.Config
	httpClient *resty.Client
	logger     Logger   // Logger
	Services   services // API Services
}

func NewFourService(cfg config.Config) *FourClient {
	FourPX := &FourClient{
		config: &cfg,
		logger: createLogger(),
	}
	// OnBeforeRequest：设置请求发送前的钩子函数，允许在请求发送之前对请求进行修改或添加逻辑。
	// OnAfterResponse：设置响应接收后的钩子函数，允许在接收到响应后处理响应数据或执行其他逻辑。
	// SetRetryCount：设置请求失败时的最大重试次数。
	// SetRetryWaitTime：设置每次重试之间的等待时间（最小等待时间）。
	// SetRetryMaxWaitTime：设置每次重试之间的最大等待时间，实际等待时间会在最小和最大等待时间之间随机选取。
	// AddRetryCondition：添加自定义的重试条件，当满足该条件时触发重试机制。
	httpClient := resty.
		New().
		SetDebug(FourPX.config.Debug).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"User-Agent":   userAgent,
		})
	if cfg.Sandbox {
		httpClient.SetBaseURL("https://open-test.4px.com/router/api/service")
	} else {
		httpClient.SetBaseURL("https://open.4px.com/router/api/service")
	}
	httpClient.
		SetTimeout(time.Duration(cfg.Timeout) * time.Second).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			param := "{}"
			if request.Body != nil {
				bd, e := sonic.Marshal(request.Body)
				if e != nil {
					return e
				}
				param = string(bd)
			}
			timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
			// fmt.Println("时间戳：", timestamp)
			v := cfg.Version
			if request.URL == "ds.xms.order.cancel" {
				v = "1.0.0"
			}
			sign, err := GetSign(param, request.URL, cfg.APIKey, cfg.APISecret, timestamp, v)
			if err != nil {
				return err
			}
			request.SetQueryParams(map[string]string{
				"app_key":   cfg.APIKey,
				"format":    "json",
				"language":  "en",
				"method":    request.URL,
				"sign":      sign,
				"timestamp": timestamp,
				"v":         v,
			})
			request.URL = ""
			return nil
		}).
		OnAfterResponse(func(client *resty.Client, response *resty.Response) (err error) {
			if response.IsError() {
				return fmt.Errorf("%s: %s", response.Status(), bytex.ToString(response.Body()))
			}
			contentType := response.RawResponse.Header.Get("Content-Type")
			if contentType == "application/octet-stream" {
				return
			}
			r := model.Response{}
			if err = sonic.Unmarshal(response.Body(), &r); err == nil {
				if r.Result != "1" {
					return fmt.Errorf("%s: %s", r.Result, r.Msg)
				}
				//  自定义响应数据
				// if r.ApiResultData != "" {
				// 	ok, err := DecodeMsg(r.ApiResultData, tempToken, cfg.EncodingAesKey, cfg.APIKey)
				// 	if err == nil {
				// 		response.SetBody([]byte(ok))
				// 	}
				// }
			} else {
				FourPX.logger.Errorf("JSON Unmarshal error: %s", err.Error())
			}
			if err != nil {
				FourPX.logger.Errorf("OnAfterResponse error: %s", err.Error())
			}
			return
		}).
		SetRetryCount(1).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(10 * time.Second).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			text := r.Request.URL
			if r == nil {
				return false
			}
			if err != nil {
				text += fmt.Sprintf(", error: %s", err.Error())
				FourPX.logger.Debugf("Retry request: %s", text)
				return true // 如果有错误则重试
			}
			// 检查响应状态码是否不是200
			if r.StatusCode() != http.StatusOK {
				text += fmt.Sprintf(", error: %d", r.StatusCode())
				FourPX.logger.Debugf("Retry request: %s", text)
				return true
			}
			type ResponseBody struct {
				Result string `json:"result"`
			}
			// 解析响应体JSON
			var responseBody ResponseBody
			if err := sonic.Unmarshal(r.Body(), &responseBody); err != nil {
				text += fmt.Sprintf(", error: %s", string(r.Body()))
				FourPX.logger.Debugf("Retry request: %s", text)
				return true // 如果解析错误则重试
			}
			// 检查apiResultCode字段是否不是1
			if responseBody.Result != "1" {
				text += fmt.Sprintf(", error: %s", responseBody.Result)
				FourPX.logger.Debugf("Retry request: %s", text)
				return true
			}
			return false
		})
	FourPX.httpClient = httpClient
	xService := service{
		config:     &cfg,
		logger:     FourPX.logger,
		httpClient: FourPX.httpClient,
	}
	FourPX.Services = services{
		Order:       (orderService)(xService),       // 订单
		Track:       (trackService)(xService),       // 追踪
		Fee:         (feeService)(xService),         // 费用
		Appointment: (appointmentService)(xService), // 预约
	}
	return FourPX
}
