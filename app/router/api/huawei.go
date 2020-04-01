package api

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"huage.tech/mini/app/dao"
	"time"
)

type errorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Title   string `json:"title"`
}
type tokenData struct {
	ExpiresAt time.Time `json:"expires_at"`
	IssuedAt  time.Time `json:"issued_at"`
}
type tokenSuccess struct {
	Token tokenData `json:"token"`
}

type tokenError struct {
	Error errorData `json:"error"`
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
		Post("https://iam.cn-north-4.myhuaweicloud.com/v3/auth/tokens")
	if err != nil {
		return
	}
	if resErr.Error.Code > 0 {
		err = errors.New(resErr.Error.Title + resErr.Error.Message)
		return
	}

	token = resp.Header().Get("X-Subject-Token")
	dao.SetConfig("HwToken", token)
	dao.SetConfig("HwExpireAt", resOK.Token.ExpiresAt.Format("2006-01-02 15:04:05"))

	return
}
