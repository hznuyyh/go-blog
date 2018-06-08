package auth

import (
	"go-blog/controllers"
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
	"go-blog/controllers/structs"
	"go-blog/models"
)

type UserController struct {
	controllers.BaseController
}

func (this *UserController) UserLogin()  {
	//返回结果参数定义
	var code    int
	var message string
	var data    map[string]interface{}
	//获取参数
	params    := structs.UserLoginParams{}
	this.RequestData(&params)
	//匹配模型
	user      := models.User{}
	o   := orm.NewOrm()
	err := o.Read(&user,"user_name")
	//数据分析
	if err == orm.ErrNoRows || err == orm.ErrMissPK{
		code    = controllers.ERROR_CODE
		message = "不存在的账号"
		this.ReturnData(code,message,data)
	}
	//密码比较
	md5_s         := md5.New()
	md5_s.Write([]byte(params.Password))
	true_password := user.Password
	if hex.EncodeToString(md5_s.Sum(nil)) != true_password  {
		code    = controllers.ERROR_CODE
		message = "密码错误"
	} else {
		code              = controllers.SUCCESS_CODE
		message           = "登录成功"
		data              = make(map[string]interface{})
		data["user_id"]   = user.Id
		data["user_name"] = user.UserName
		data["nick_name"] = user.NickName
		this.SetSession("uid",user.Id)
	}
	//数据输出
	this.ReturnData(code,message,data)
}
