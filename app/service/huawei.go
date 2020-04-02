package service

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"huage.tech/mini/app/dao"
	"strconv"
	"time"
)

const HwServer = "https://iam.cn-north-4.myhuaweicloud.com"
const iotServer = "https://iotdm.cn-north-4.myhuaweicloud.com"

type project struct {
	Id     string      `json:"id"`
	Name   string      `json:"name"`
	Domain interface{} `json:"domain"`
}
type errorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Title   string `json:"title"`
}
type Project struct {
	Domain map[string]string `json:"domain"`
	Id     string            `json:"id"`
	Name   string            `json:"name"`
}
type tokenData struct {
	ExpiresAt time.Time `json:"expires_at"`
	IssuedAt  time.Time `json:"issued_at"`
	Project   Project   `json:"project"`
}
type tokenSuccess struct {
	Token tokenData `json:"token"`
}

type tokenError struct {
	Error errorData `json:"error"`
}
type DevicesResult struct {
	Devices []SimpleDevice `json:"devices"`
}
type SimpleDevice struct {
	DeviceName string `json:"device_name"`
	DeviceId   string `json:"device_id"`
	ProjectId  string `json:"product_id"`
	Status     string `json:"status"`
}
type HwCommand struct {
	ServiceId   string            `json:"service_id"`
	CommandName string            `json:"command_name"`
	Paras       map[string]string `json:"paras"`
}
type HwCommandResponse struct {
	ResultCode   int            `json:"result_code"`
	ResponseName string            `json:"response_name"`
	Paras        map[string]string `json:"paras"`
}
type HwCommandResult struct {
	CommandId string            `json:"command_id"`
	Response  HwCommandResponse `json:"response"`
}

func HwToken() (token string, err error) {
	expire := dao.GetConfig("HwExpireAt")
	ex, err := time.Parse("2006-01-02 15:04:05", expire)
	sh, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(sh)
	if ex.Sub(now) > 10*time.Minute {
		token = dao.GetConfig("HwToken")
		return
	}

	client := resty.New()
	resOK := tokenSuccess{}
	resErr := tokenError{}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{
    "auth": {
        "identity": {
            "methods": [
                "password"
            ],
            "password": {
                "user": {
                    "name": "aczhanghua",
                    "password": "sdfihua1",
                    "domain": {
                        "name": "aczhanghua"
                    }
                }
            }
        },
        "scope": {
            "project": {
                "name": "cn-north-4"
            }
        }
    }
}`).SetResult(&resOK).SetError(&resErr).
		Post(HwServer + "/v3/auth/tokens")
	if err != nil {
		return
	}
	if resErr.Error.Code > 0 {
		err = errors.New(resErr.Error.Title + resErr.Error.Message)
		return
	}

	token = resp.Header().Get("X-Subject-Token")
	dao.SetConfig("HwToken", token)
	dao.SetConfig("HwProjectId", resOK.Token.Project.Id)
	dao.SetConfig("HwExpireAt", resOK.Token.ExpiresAt.Format("2006-01-02 15:04:05"))

	return
}
func GetProducts(offset, limit int64) (devices []SimpleDevice, err error) {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	project := dao.GetConfig("HwProjectId")
	token, err := HwToken()
	if err != nil {
		return
	}

	dr := DevicesResult{}
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"limit":  strconv.FormatInt(limit, 10),
			"offset": strconv.FormatInt(offset, 10),
		}).
		SetHeader("Accept", "application/json").
		SetHeader("X-Auth-Token", token).
		SetResult(&dr).
		Get(iotServer + "/v5/iot/" + project + "/products")
	fmt.Println(string(resp.Body()))
	devices = dr.Devices
	return
}
func GetDevices(offset, limit int64) (devices []SimpleDevice, err error) {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	project := dao.GetConfig("HwProjectId")
	token, err := HwToken()
	if err != nil {
		return
	}

	dr := DevicesResult{}
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"limit":  strconv.FormatInt(limit, 10),
			"offset": strconv.FormatInt(offset, 10),
		}).
		SetHeader("Accept", "application/json").
		SetHeader("X-Auth-Token", token).
		SetResult(&dr).
		Get(iotServer + "/v5/iot/" + project + "/devices")
	fmt.Println(string(resp.Body()))
	devices = dr.Devices
	return
}

func GetProductInfo(device_id string) (result string, err error) {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	project := dao.GetConfig("HwProjectId")
	token, err := HwToken()
	if err != nil {
		return
	}

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Auth-Token", token).
		//SetResult(&dr).
		Get(iotServer + "/v5/iot/" + project + "/products/" + device_id + "/properties")

	result = string(resp.Body())
	return
}

func SendDeviceCommand(device_id string, cmd HwCommand) (result HwCommandResult, err error) {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	project := dao.GetConfig("HwProjectId")
	token, err := HwToken()
	if err != nil {
		return
	}

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-Auth-Token", token).
		SetBody(cmd).
		SetResult(&result).
		Post(iotServer + "/v5/iot/" + project + "/devices/" + device_id + "/commands")
	fmt.Println(string(resp.Body()))
	return
}
