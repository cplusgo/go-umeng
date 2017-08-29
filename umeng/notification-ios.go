package umeng

type PayloadIOS struct {
	Aps *PayloadIOSAps `json:"aps"`
}

type PayloadIOSAps struct {
	Alert string `json:"alert"`
}

type NotificationIOS struct {
	UmengNotification
}

func (this *NotificationIOS) setConfig() {
	this.appMasterSecret = IOS_APP_MASTER_SECRET
	this.uploadPath = UPLOAD_PATH
	this.postPath = POST_PATH
	this.host = HOST
}
