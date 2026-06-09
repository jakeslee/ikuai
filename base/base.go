package base

import (
	"crypto/tls"
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/jakeslee/ikuai/action"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type IKuaiBase struct {
	Client   *resty.Client
	DebugOn  bool
	Url      string
	Username string
	Password string
	session  string
}

func (i *IKuaiBase) Run(action *action.Action, result interface{}) (string, error) {
	url := i.Url + "/Action/call"

	response, err := i.Client.R().
		SetDebug(i.DebugOn).
		EnableGenerateCurlOnDebug().
		SetHeader("Content-Type", "application/json").
		SetCookie(&http.Cookie{Name: "sess_key", Value: i.session}).
		SetBody(action).
		SetResult(result).
		Post(url)

	if err != nil {
		return "", err
	}

	return response.String(), nil
}

func (i *IKuaiBase) Debug() {
	i.DebugOn = true
}

type Retrier interface {
	RetryHandle(response *resty.Response, err error) bool
}

func CreateHttpClient(insecureSkipVerify, autoLogin bool, retrier Retrier) *resty.Client {
	client := resty.New()
	client.JSONUnmarshal = func(b []byte, v interface{}) error {
		body := string(b)
		// Handle invalid JSON structure when ikuai returns "data: timeout"
		results := gjson.GetMany(body, "results.data", "results.total")
		if results[0].Raw == "timeout" {
			if results[1].Exists() {
				set, _ := sjson.Set(body, "results.data", []string{})
				body = set
			}

			logrus.WithFields(logrus.Fields{
				"result":     string(b),
				"normalized": body,
			}).Warn("ikuai returns invalid JSON: \"data: timeout\"")
		}

		return json.Unmarshal([]byte(body), v)
	}

	if insecureSkipVerify {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	if autoLogin {
		client.SetRetryCount(2)
		client.AddRetryCondition(func(response *resty.Response, err error) bool {
			body := response.Body()

			if err != nil {
				logrus.WithFields(logrus.Fields{
					"result": string(body),
				}).WithError(err).Error("request error")
			}

			return retrier.RetryHandle(response, err)
		})
	}

	return client
}
