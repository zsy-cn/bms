package internal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/zsy-cn/bms/conf"
)

func (cs *DefaultCameraService) requestAccessToken() {
	accessToken, err := cs.RequestAccessToken()
	if err != nil {
		cs.l.Errorf("request access token failed in requestAccessToken(): %s", err.Error())
		return
	}
	err = cs.writeConfigFile("access-token", accessToken)
	if err != nil {
		cs.l.Errorf("write access token failed in requestAccessToken(): %s", err.Error())
		return
	}
	cs.accessToken = accessToken
}

func (cs *DefaultCameraService) RequestAccessToken() (accessToken string, err error) {
	cs.l.Debug("request access token in requestToken()")
	contentType := "application/x-www-form-urlencoded"

	dataURLVal := url.Values{}
	dataURLVal.Add("appKey", cs.appKey)
	dataURLVal.Add("appSecret", cs.appSecret)
	data := strings.NewReader(dataURLVal.Encode())
	resp, err := http.Post(conf.CameraServiceYS7Addr, contentType, data)
	if err != nil {
		cs.l.Errorf("request ys to get access token failed: %s", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cs.l.Errorf("request ys to get access token failed: %s", err.Error())
		return
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		cs.l.Errorf("unmarshal access token response failed: %s", err.Error())
		return
	}
	if code, ok := result["code"]; !ok || code.(string) != "200" {
		cs.l.Errorf("we didn't get a valid access token: %s", result["msg"])
		return
	}

	resultDataField := result["data"].(map[string]interface{})
	accessToken = resultDataField["accessToken"].(string)

	return
}
