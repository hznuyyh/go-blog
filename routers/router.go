package routers

import (
	"github.com/astaxie/beego"
	"go-blog/controllers/auth"
)

func init() {
    beego.Router("/user/login",&auth.UserController{},"*:UserLogin" )
}
