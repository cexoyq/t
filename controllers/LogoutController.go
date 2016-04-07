package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {
	this.DestroySession();
	//this.TplName = "login.html"
	this.Ctx.Redirect(302,"/login")
}

