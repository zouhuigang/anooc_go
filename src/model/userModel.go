package model

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	. "db"
	"github.com/polaris1119/goutils"
	"github.com/polaris1119/logger"
	"golang.org/x/net/context"
	"net/url"
)

// 用户登录信息
type UserLogin struct {
	Uid       int       `json:"uid" xorm:"pk"`
	Username  string    `json:"username"`
	Passwd    string    `json:"passwd"`
	Email     string    `json:"email"`
	LoginTime time.Time `json:"login_time" xorm:"<-"`
	Passcode  string    `json:"passcode"` // 加密随机串
}

func (this *UserLogin) TableName() string {
	return "user_login"
}

// 生成加密密码
func (this *UserLogin) GenMd5Passwd() error {
	if this.Passwd == "" {
		return errors.New("password is empty!")
	}
	this.Passcode = fmt.Sprintf("%x", rand.Int31())
	// 密码经过md5(passwd+passcode)加密保存
	this.Passwd = goutils.Md5(this.Passwd + this.Passcode)
	return nil
}

const (
	UserStatusNoAudit = iota
	UserStatusAudit   // 已激活
	UserStatusRefuse
	UserStatusFreeze // 冻结
	UserStatusOutage // 停用
)

// 用户基本信息
type User struct {
	Uid         int       `json:"uid" xorm:"pk autoincr"`
	Username    string    `json:"username" validate:"min=4,max=20,regexp=^[a-zA-Z0-9_]*$"`
	Email       string    `json:"email"`
	Open        int       `json:"open"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	City        string    `json:"city"`
	Company     string    `json:"company"`
	Github      string    `json:"github"`
	Weibo       string    `json:"weibo"`
	Website     string    `json:"website"`
	Monlog      string    `json:"monlog"`
	Introduce   string    `json:"introduce"`
	Unsubscribe int       `json:"unsubscribe"`
	Status      int       `json:"status"`
	IsRoot      bool      `json:"is_root"`
	Ctime       OftenTime `json:"ctime" xorm:"created"`
	Mtime       time.Time `json:"mtime" xorm:"<-"`

	// 非用户表中的信息，为了方便放在这里
	Roleids   []int    `xorm:"-"`
	Rolenames []string `xorm:"-"`
}

func (this *User) TableName() string {
	return "user_info"
}

func (this *User) String() string {
	buffer := goutils.NewBuffer()
	buffer.Append(this.Username).Append(this.Email).Append(this.Uid).Append(this.Mtime)

	return buffer.String()
}

// Me 代表当前用户
type Me struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
	MsgNum   int    `json:"msgnum"`
	IsAdmin  bool   `json:"isadmin"`
	IsRoot   bool   `json:"is_root"`
}

// 活跃用户信息
// 活跃度规则：
//	1、注册成功后 +2
//	2、登录一次 +1
//	3、修改资料 +1
//	4、发帖子 + 10
//	5、评论 +5
//	6、创建Wiki页 +10
type UserActive struct {
	Uid      int       `json:"uid" xorm:"pk"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Avatar   string    `json:"avatar"`
	Weight   int       `json:"weight"`
	Mtime    time.Time `json:"mtime" xorm:"<-"`
}

// 用户角色信息
type UserRole struct {
	Uid    int    `json:"uid"`
	Roleid int    `json:"roleid"`
	ctime  string `xorm:"-"`
}

//申明model
type UsersModel struct{}

var Users = UsersModel{}

//用户注册
func (self UsersModel) CreateUser(ctx context.Context, form url.Values) (errMsg string, err error) {
	objLog := GetLogger(ctx)

	user := &User{}
	err = schemaDecoder.Decode(user, form)
	if err != nil {
		objLog.Errorln("user schema Decode error:", err)
		errMsg = err.Error()
		return
	}

	session := MasterDB.NewSession()
	defer session.Close()

	session.Begin()

	_, err = session.Insert(user)
	if err != nil {
		session.Rollback()
		errMsg = "内部服务器错误"
		objLog.Errorln(errMsg, ":", err)
		return
	}

	// 存用户登录信息
	userLogin := &UserLogin{}
	err = schemaDecoder.Decode(userLogin, form)
	if err != nil {
		session.Rollback()
		errMsg = err.Error()
		objLog.Errorln("CreateUser error:", err)
		return
	}
	userLogin.Uid = user.Uid
	err = userLogin.GenMd5Passwd()
	if err != nil {
		session.Rollback()
		errMsg = err.Error()
		return
	}
	if _, err = session.Insert(userLogin); err != nil {
		session.Rollback()
		errMsg = "内部服务器错误"
		logger.Errorln(errMsg, ":", err)
		return
	}

	session.Commit()

	return

}

func (self UsersModel) Userexit(ctx context.Context, field, val string) bool {
	objLog := GetLogger(ctx)

	userLogin := &UserLogin{}
	_, err := MasterDB.Where(field+"=?", val).Get(userLogin)
	if err != nil || userLogin.Uid == 0 {
		if err != nil {
			objLog.Errorln("user logic UserExists error:", err)
		}
		return false
	}
	return true

}
