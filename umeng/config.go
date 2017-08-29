package umeng

/*
 *全局配置，方便起见在代码中配置，也可以写在配置表
 */
//base config
const (
	HOST        = "http://msg.umeng.com"
	UPLOAD_PATH = "/upload"
	POST_PATH   = "/api/send"

	ANDROID_APP_KEY           = "xxxx"
	ANDROID_APP_MASTER_SECRET = "xxxx"

	IOS_APP_KEY           = "xxxx"
	IOS_APP_MASTER_SECRET = "xxxx"
)

//after open
const (
	GO_APP      = "go_app"
	GO_URL      = "go_url"
	GO_ACTIVITY = "go_activity"
	GO_CUSTOM   = "go_custom"
)

//display payload type
const (
	NOTIFICATION = "notification"
	MESSAGE      = "message"
)

//os type 手机类型 1.苹果手机 2.安卓手机
const (
	IOS     = 1
	ANDROID = 2
)

//message type 消息类型，如定向发送
const(
	UNICAST = "unicast"
)