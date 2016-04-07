package controllers

import (
	"github.com/astaxie/beego"
)

type Player_4Controller struct {
	beego.Controller
}

func (this *Player_4Controller) Get() {
	this.Data["Website"] = beego.AppConfig.String("Website")
	sess_username := this.GetSession("UserName")
	if sess_username != nil {
			this.Data["UserName"] = sess_username
	}
	this.TplName = "player_4.html"
}

