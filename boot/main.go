package main

import "github.com/cplusgo/go-umeng/umeng"

func main() {
	os := 1
	if os == umeng.IOS {
		umeng.NewUnicastIOS("恭喜你，你的作品被添加到首页！", "f39a628311ffe034ebe9be661315e76382672d770b9a15667fdb1fbac9668cce")
	} else if os == umeng.ANDROID {
		umeng.NewAndroidUnicast("借我说得出口的旦旦誓言", "AjfquSK5oypIYnr3HnYYD2aEjxrIQapcJrAmMF95CRZe")
	}
}
