package controllers

import (
	"reflect"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"go-blog/controllers/structs"
)


//控制器声明
type BaseController struct {
	beego.Controller
}

/**
 * 接收参数方法
 * @param param interface{} 对应 structs 地址
 */
func (this *BaseController) RequestData(param interface{}) {
	this.ParseForm(param)
	valida := validation.Validation{}
	valida.Valid(param)
	message := ""
	if valida.HasErrors() {
		var field string
		for _ , e := range valida.Errors {
			if e.Tmpl == "格式不正确" {
				message = "格式不正确"
				field = e.Field
				break
			} else {
				message = e.Message
				field = e.Field
				break
			}

		}
		cnField := structs.GetCnField(field)
		if cnField != "" {
			message = cnField + message
		} else {
			message = field + " " + message
		}
		var data map[string]interface{}
		this.ReturnData( 3999 , message , data)
	}
}


/**
 * 返回数据，输出json中断当前请求
 * @param code int 返回状态码
 * @param message string 返回提示信息
 * @param data map[string]interface{} 返回的数据体
 */
func (this *BaseController) ReturnData(code int, message string, data interface{}) {
	this.Data["json"] = map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	}
	this.ServeJSON()
	this.StopRun()
}


/**
 * 将struct转化成map
 * @param : obj struct 要转换的struct
 * @return : data map[string]interface{} 转换后的map
 */
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[SnakeString(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}

func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}