package main

import (
    "anooc_go/models"
	"fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
	
    _"github.com/go-sql-driver/mysql"
	_ "anooc_go/routers"
)




func init() {
    // 注册驱动
    orm.RegisterDriver("mysql", orm.DRMySQL)
    // 注册默认数据库
    // 我的mysql的root用户密码为tom，打算把数据表建立在名为test数据库里
    // 备注：此处第一个参数必须设置为“default”（因为我现在只有一个数据库），否则编译报错说：必须有一个注册DB的别名为 default
    orm.RegisterDataBase("default", "mysql", "root:@/godb?charset=utf8")
}


func main() {
    // 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
    orm.Debug = true
    // 自动建表
    orm.RunSyncdb("default", false, true)

    // 创建一个 ormer 对象
    o := orm.NewOrm()
    o.Using("default")
    perfile := new(models.Profile)
    perfile.Age = 30

    user := new(models.User)
    user.Name = "tom"
    user.Profile = perfile

    // insert
    o.Insert(perfile)
    o.Insert(user)
    o.Insert(perfile)
    o.Insert(user)
    o.Insert(perfile)
    o.Insert(user)

    // update
    user.Name = "hezhixiong"
    num, err := o.Update(user)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)

    // delete
    o.Delete(&models.User{Id: 2})
	
    beego.Run()
}