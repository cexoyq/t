package controllers

import (
	"github.com/astaxie/beego"
)

type StatusController struct {
	beego.Controller
}

func (this *StatusController) Get() {
	this.Data["Website"] = beego.AppConfig.String("Website")
	sess_username := this.GetSession("UserName")
	if sess_username != nil {
			this.Data["UserName"] = sess_username
	}
	this.TplName = "status.html"
}