package umeng

import (
	"encoding/json"
	"github.com/cplusgo/go-umeng/helper"
	"net/http"
	"bytes"
	"io/ioutil"
	"log"
)

type UmengNotificationData struct {
	AppKey         string `json:"appkey"`
	Timestamp      string `json:"timestamp"`
	Type           string `json:"type"`
	DeviceTokens   string `json:"device_tokens"`
	ProductionMode bool `json:"production_mode"`
	Payload        interface{} `json:"payload"`
}

type UmengNotification struct {
	host            string
	uploadPath      string
	postPath        string
	appMasterSecret string
	data            *UmengNotificationData
}

func (this *UmengNotification) send() error {
	url := this.host + this.postPath
	postBody, err := json.Marshal(this.data)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	sign := helper.MD5("POST" + url + string(postBody) + this.appMasterSecret)
	url = url + "?sign=" + sign
	bufReader := bytes.NewReader(postBody)
	resp, err := http.Post(url, "application/json", bufReader)
	defer func() {
		resp.Body.Close()
	}()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println(string(content))
	return nil
}
