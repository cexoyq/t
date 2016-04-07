package controllers

import (
	"github.com/astaxie/beego"
)

type Status_dvrController struct {
	beego.Controller
}

func (this *Status_dvrController) Get() {
	this.Data["Website"] = beego.AppConfig.String("Website")
	sess_username := this.GetSession("UserName")
	if sess_username != nil {
			this.Data["UserName"] = sess_username
	}
	this.TplName = "status_dvr.html"
}