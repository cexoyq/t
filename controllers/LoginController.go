
package controllers

import (
	"t/models"
	"log"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	log.Println("LoginController:Get")
	var(
		//s loginUser
		//err error
	) 
	//s = loginUser{}
	id := this.GetSession("id")
    if id == nil {			//用户还未登陆
		this.Data["Err"] = "请登陆！"
		this.TplName = "login.html"
    } else {
        this.Redirect("/",200)
    }

}

func (this *LoginController) Post() {
	var (
		s models.LoginUser
		err error
	)
	s = models.LoginUser{}
	log.Println("LoginController:Post")
	if err = this.ParseForm(&s); err != nil {
		this.Data["Err"] = "登陆失败，不能获取您的帐号及密码,请重试！"
		this.TplName = "login.html"
		return
	}
	s,err = models.ValidateUser(s.UserName,s.Password)
	if err != nil {
		this.Data["Err"] = "登陆失败，帐号及密码不符,请重试！"
		this.TplName = "login.html"
		return
	}
		this.SetSession("id",s.Id)
		this.SetSession("UserName",s.UserName)
		this.SetSession("level",s.Level)
		
		log.Println("用户成功登录！")
		//this.Ctx.Redirect(200,"/")
		this.Redirect("/",200)
		//this.TplName = "index.html"

}
