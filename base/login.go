package base

import (
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

type LoginRequest struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
}

func (i *IKuaiBase) Login() (string, error) {
	response, err := i.Client.R().
		SetBody(&LoginRequest{
			Username: i.Username,
			Passwd:   GetMD5(i.Password),
		}).
		Post(i.Url + "/Action/login")

	if err != nil {
		return "", err
	}

	for _, cookie := range response.Cookies() {
		if cookie.Name == "sess_key" {
			i.session = cookie.Value
			return cookie.Value, nil
		}
	}

	body := response.Body()

	logrus.WithFields(logrus.Fields{
		"URL":      i.Url + "/Action/login",
		"username": i.Username,
		"result":   string(body),
	}).Error("failed to log in")

	return "", errors.New("login error")
}

func GetMD5(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x", sum)
}
