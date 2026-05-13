package ikuai

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jakeslee/ikuai/action"
	"github.com/stretchr/testify/assert"
)

func TestGetMD5(t *testing.T) {
	tests := []struct {
		name     string
		password string
		expected string
	}{
		{
			name:     "simple password",
			password: "123456",
			expected: fmt.Sprintf("%x", md5.Sum([]byte("123456"))),
		},
		{
			name:     "empty password",
			password: "",
			expected: fmt.Sprintf("%x", md5.Sum([]byte(""))),
		},
		{
			name:     "complex password",
			password: "Test@123!abc",
			expected: fmt.Sprintf("%x", md5.Sum([]byte("Test@123!abc"))),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getMD5(tt.password)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewIKuai(t *testing.T) {
	tests := []struct {
		name               string
		url                string
		username           string
		password           string
		insecureSkipVerify bool
		autoLogin          bool
	}{
		{
			name:               "basic initialization",
			url:                "https://192.168.1.1",
			username:           "admin",
			password:           "password",
			insecureSkipVerify: false,
			autoLogin:          false,
		},
		{
			name:               "with insecure skip verify",
			url:                "https://192.168.1.1",
			username:           "admin",
			password:           "password",
			insecureSkipVerify: true,
			autoLogin:          false,
		},
		{
			name:               "with auto login",
			url:                "https://192.168.1.1",
			username:           "admin",
			password:           "password",
			insecureSkipVerify: false,
			autoLogin:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ikuai := NewIKuai(tt.url, tt.username, tt.password, tt.insecureSkipVerify, tt.autoLogin)

			assert.NotNil(t, ikuai)
			assert.Equal(t, tt.url, ikuai.Url)
			assert.Equal(t, tt.username, ikuai.Username)
			assert.Equal(t, tt.password, ikuai.Password)
			assert.NotNil(t, ikuai.client)
			assert.False(t, ikuai.debug)
		})
	}
}

func TestIKuai_Debug(t *testing.T) {
	ikuai := NewIKuai("https://192.168.1.1", "admin", "password", false, false)

	assert.False(t, ikuai.debug)

	ikuai.Debug()

	assert.True(t, ikuai.debug)
}

func TestIKuai_Login_Success(t *testing.T) {
	sessionKey := "test-session-key-12345"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/Action/login", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var loginReq LoginRequest
		err := json.NewDecoder(r.Body).Decode(&loginReq)
		assert.NoError(t, err)
		assert.Equal(t, "admin", loginReq.Username)

		expectedMD5 := getMD5("password")
		assert.Equal(t, expectedMD5, loginReq.Passwd)

		w.Header().Set("Set-Cookie", "sess_key="+sessionKey)
		w.WriteHeader(http.StatusOK)

		result := action.Result{
			Result: 0,
			ErrMsg: "Success",
		}
		json.NewEncoder(w).Encode(result)
	}))
	defer server.Close()

	ikuai := NewIKuai(server.URL, "admin", "password", false, false)
	session, err := ikuai.Login()

	assert.NoError(t, err)
	assert.Equal(t, sessionKey, session)
	assert.Equal(t, sessionKey, ikuai.session)
}

func TestIKuai_Login_Failure_NoCookie(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		result := action.Result{
			Result: 1,
			ErrMsg: "Invalid credentials",
		}
		json.NewEncoder(w).Encode(result)
	}))
	defer server.Close()

	ikuai := NewIKuai(server.URL, "admin", "wrong_password", false, false)
	session, err := ikuai.Login()

	assert.Error(t, err)
	assert.Empty(t, session)
	assert.Contains(t, err.Error(), "login error")
}

func TestIKuai_Login_NetworkError(t *testing.T) {
	ikuai := NewIKuai("http://invalid-host-that-does-not-exist:9999", "admin", "password", false, false)
	session, err := ikuai.Login()

	assert.Error(t, err)
	assert.Empty(t, session)
}

func TestIKuai_Run_Success(t *testing.T) {
	sessionKey := "test-session-key"
	expectedResponse := `{"Result":0,"ErrMsg":"Success","data":[]}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/Action/call", r.URL.Path)
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		cookie, err := r.Cookie("sess_key")
		assert.NoError(t, err)
		assert.Equal(t, sessionKey, cookie.Value)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expectedResponse))
	}))
	defer server.Close()

	ikuai := NewIKuai(server.URL, "admin", "password", false, false)
	a := &action.Action{
		Action:   "show",
		FuncName: "test_func",
		Param:    map[string]interface{}{"key": "value"},
	}

	var result action.Result
	response, err := ikuai.Run(sessionKey, a, &result)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}

func TestIKuai_Run_WithDebug(t *testing.T) {
	sessionKey := "test-session-key"
	expectedResponse := `{"Result":0,"ErrMsg":"Success"}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expectedResponse))
	}))
	defer server.Close()

	ikuai := NewIKuai(server.URL, "admin", "password", false, false)
	ikuai.Debug()

	a := &action.Action{
		Action:   "show",
		FuncName: "test_func",
	}

	var result action.Result
	response, err := ikuai.Run(sessionKey, a, &result)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
	assert.True(t, ikuai.debug)
}

func TestIKuai_Run_NetworkError(t *testing.T) {
	ikuai := NewIKuai("http://invalid-host-that-does-not-exist:9999", "admin", "password", false, false)

	a := &action.Action{
		Action:   "show",
		FuncName: "test_func",
	}

	var result action.Result
	response, err := ikuai.Run("session-key", a, &result)

	assert.Error(t, err)
	assert.Empty(t, response)
}

func TestIKuai_Login_AutoRetryOnSessionTimeout(t *testing.T) {
	sessionKey := "new-session-key"
	loginCount := 0
	callCount := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/Action/login" {
			loginCount++
			w.Header().Set("Set-Cookie", "sess_key="+sessionKey)
			w.WriteHeader(http.StatusOK)

			result := action.Result{
				Result: 0,
				ErrMsg: "Success",
			}
			json.NewEncoder(w).Encode(result)
		} else if r.URL.Path == "/Action/call" {
			callCount++
			w.WriteHeader(http.StatusOK)

			result := action.Result{
				Result: 10014,
				ErrMsg: "Session timeout",
			}
			json.NewEncoder(w).Encode(result)
		}
	}))
	defer server.Close()

	ikuai := NewIKuai(server.URL, "admin", "password", false, true)

	var result action.Result
	action := &action.Action{
		Action:   "show",
		FuncName: "test_func",
	}

	_, err := ikuai.Run(sessionKey, action, &result)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, callCount, 1)
	assert.GreaterOrEqual(t, loginCount, 1)
}

func TestNewIKuai_WithAutoLogin_RetryCondition(t *testing.T) {
	sessionKey := "auto-retry-session"
	requestCount := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/Action/login" {
			w.Header().Set("Set-Cookie", "sess_key="+sessionKey)
			w.WriteHeader(http.StatusOK)

			result := action.Result{
				Result: 0,
				ErrMsg: "Success",
			}
			json.NewEncoder(w).Encode(result)
		} else if r.URL.Path == "/Action/call" {
			requestCount++
			w.WriteHeader(http.StatusOK)

			if requestCount == 1 {
				result := action.Result{
					Result: 10014,
					ErrMsg: "Session timeout",
				}
				json.NewEncoder(w).Encode(result)
			} else {
				result := action.Result{
					Result: 0,
					ErrMsg: "Success",
				}
				json.NewEncoder(w).Encode(result)
			}
		}
	}))
	defer server.Close()

	ikuai := NewIKuai(server.URL, "admin", "password", false, true)

	var result action.Result
	action := &action.Action{
		Action:   "show",
		FuncName: "test_func",
	}

	_, err := ikuai.Run("old-session", action, &result)

	assert.NoError(t, err)
	assert.Equal(t, 2, requestCount)
}

func TestIKuai_Login_InvalidJSONResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json response"))
	}))
	defer server.Close()

	ikuai := NewIKuai(server.URL, "admin", "password", false, false)
	session, err := ikuai.Login()

	assert.Error(t, err)
	assert.Empty(t, session)
}

func TestIKuai_MultipleLogins(t *testing.T) {
	sessionKey1 := "session-key-1"
	sessionKey2 := "session-key-2"
	loginAttempts := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loginAttempts++
		sessionKey := sessionKey1
		if loginAttempts == 2 {
			sessionKey = sessionKey2
		}

		w.Header().Set("Set-Cookie", "sess_key="+sessionKey)
		w.WriteHeader(http.StatusOK)

		result := action.Result{
			Result: 0,
			ErrMsg: "Success",
		}
		json.NewEncoder(w).Encode(result)
	}))
	defer server.Close()

	ikuai := NewIKuai(server.URL, "admin", "password", false, false)

	session1, err1 := ikuai.Login()
	assert.NoError(t, err1)
	assert.Equal(t, sessionKey1, session1)

	session2, err2 := ikuai.Login()
	assert.NoError(t, err2)
	assert.Equal(t, sessionKey2, session2)
	assert.Equal(t, 2, loginAttempts)
}

func TestIKuai_ShowMonitorLan_Response_Invalid_JSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/Action/call" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"Result":30000,"ErrMsg":"Success","Data":{"data":timeout,"total":0}}`))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}))
	defer server.Close()

	ikuai := NewIKuai(server.URL, "admin", "password", false, true)
	lanResult, err := ikuai.ShowMonitorLan()

	assert.NoError(t, err)
	assert.True(t, lanResult.Ok())
	assert.Empty(t, lanResult.Data.Data)
}
