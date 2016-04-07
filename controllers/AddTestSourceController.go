package controllers

import (
	"github.com/astaxie/beego"
	"t/models"
)

type AddTestSourceController struct {
	beego.Controller
}

func (this *AddTestSourceController) Get() {
	this.Data["Website"] = beego.AppConfig.String("Website")
	sess_username := this.GetSession("UserName")
	if sess_username != nil {
			this.Data["UserName"] = sess_username
	}
	models.RetuTestSource()
	this.TplName = "addtestsource.html"
}