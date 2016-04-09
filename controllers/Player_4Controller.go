package controllers

import (
	"github.com/astaxie/beego"
	"t/models"
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
	p,_ := models.RetuTestSource()
	this.Data["Parent"] = p;
	this.TplName = "player_4.html"
}

