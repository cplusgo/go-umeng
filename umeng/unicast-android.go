package umeng

import (
	"strconv"
	"time"
)

type UnicastAndroid struct {
	NotificationAndroid
}

func NewAndroidUnicast(textMessage string, deviceToken string) *UnicastAndroid {
	unicast := &UnicastAndroid{
		NotificationAndroid{},
	}
	unicast.setConfig()
	payload := &PayloadAndroid{}
	payload.DisplayType = NOTIFICATION
	payload.Extra = make(map[string]interface{})
	payload.Body = &PayloadBodyAndroid{}
	payload.Body.Title = "消息通知"
	payload.Body.Ticker = "消息通知"
	payload.Body.PlayLights = true
	payload.Body.PlaySound = true
	payload.Body.PlayVibrate = true
	payload.Body.Text = textMessage
	payload.Body.AfterOpen = GO_APP

	data := &UmengNotificationData{}
	data.Payload = payload
	data.AppKey = ANDROID_APP_KEY
	data.DeviceTokens = deviceToken
	data.ProductionMode = true
	data.Timestamp = strconv.Itoa(int(time.Now().Unix()))
	data.Type = UNICAST
	unicast.data = data

	unicast.send()
	return unicast
}
