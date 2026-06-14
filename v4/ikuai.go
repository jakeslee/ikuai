package v4

import (
	"errors"
	"net"

	"github.com/go-resty/resty/v2"
	"github.com/jakeslee/ikuai/base"
	"github.com/jakeslee/ikuai/v4/action"
	"github.com/sirupsen/logrus"
)

type IKuaiV4 struct {
	*base.IKuaiBase
}

func (i *IKuaiV4) RetryHandle(response *resty.Response, err error) bool {
	body := string(response.Body())

	var result action.Status
	rErr := i.Client.JSONUnmarshal([]byte(body), &result)
	if rErr != nil {
		var nErr net.Error
		if errors.Is(rErr, base.ErrIKuaiTimeout) || errors.As(err, &nErr) && nErr.Timeout() {
			logrus.WithFields(logrus.Fields{
				"method": response.Request.Method,
				"URL":    response.Request.URL,
			}).Warn("timeout")
			return true
		}

		logrus.WithFields(logrus.Fields{
			"result": body,
		}).WithError(rErr).Error("unmarshal body error")

		return false
	}

	if !result.Ok() {
		logger := logrus.WithFields(logrus.Fields{
			"URL":     response.Request.URL,
			"result":  result.Result,
			"message": result.Message,
		})

		if !result.Is(action.SessionTimeout) {
			logger = logger.WithField("response", body)
		}

		logger.WithError(err).Warn("failed to invoke ikuai")
	}

	if result.Is(action.SessionTimeout) {
		logrus.WithFields(logrus.Fields{
			"URL":    response.Request.URL,
			"result": result,
		}).Info("session timeout")
		_, err := i.Login()
		if err != nil {
			return false
		}

		logrus.WithFields(logrus.Fields{
			"username": i.Username,
		}).Info("login successful")

		return true
	}

	return false
}
