package umeng

import (
	"strconv"
	"time"
)

type UnicastIOS struct {
	NotificationIOS
}

func NewUnicastIOS(textMessage string, deviceToken string) {
	unicast := &UnicastIOS{
		NotificationIOS{},
	}
	unicast.setConfig()

	payload := &PayloadIOS{}
	aps := &PayloadIOSAps{
		Alert: textMessage,
	}
	payload.Aps = aps

	data := &UmengNotificationData{}
	data.Payload = payload
	data.AppKey = IOS_APP_KEY
	data.DeviceTokens = deviceToken
	data.ProductionMode = true
	data.Timestamp = strconv.Itoa(int(time.Now().Unix()))
	data.Type = UNICAST
	unicast.data = data
	unicast.send()
}
