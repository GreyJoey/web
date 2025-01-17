package DBstruct

import (
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/jinzhu/gorm"
)

// password不会被json传输，这样客户端在登录之后，不会有密码的缓存，相对安全（人走开后被偷看密码和修改密码）
type User struct {
	gorm.Model
	Phone string `json:"phone"` //这里有一个“电话”信息，但是系统利用的电话信息主要是从Address记录中获取的，与这里的无关，只是为了保证扩展性而添加的冗余字段
	//User_id  int    `json:"user_id" gorm:"primary_key"`
	UserName string `json:"username" gorm:"unique"`
	Password string `json:"-"`
	Avatar   string `json:"avatar"` //头像的url
}

// AvatarURL 头像地址
func (user *User) AvatarURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(user.Avatar, oss.HTTPGet, 24*60*60)
	//if strings.Contains(signedGetURL, "http://ailiaili-img-av.oss-cn-hangzhou.aliyuncs.com/?Exp") {
	//signedGetURL := "https://ailiaili-img-av.oss-cn-hangzhou.aliyuncs.com/img/noface.png"
	//}
	return signedGetURL
}
