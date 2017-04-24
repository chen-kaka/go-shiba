package wechatalert

import (
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego"
)


func Alert(wechatId string, subject string, content string, html string, color string)  {
	req := httplib.Post("http://203.195.193.160/alarm/send")
	req.Param("userid",wechatId)
	req.Param("subject",subject)
	if content != "" {
		req.Param("content",content)
	}
	if html != "" {
		req.Param("html",html)
	}
	if color != "" {
		req.Param("color",color)
	}
	str, err := req.String()
	if err != nil {
		beego.Error("send failed, ret error: ", err)
	}
	beego.Info("send succ, ret info: ", str)
}