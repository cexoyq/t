package routers

import (
	"t/controllers"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"log"
)

func init() {
	var FilterUser = func(ctx *context.Context){
		if ctx.Input.CruSession == nil {
			log.Println("session not start!")
        	ctx.Input.CruSession , _ = beego.GlobalSessions.SessionStart(ctx.ResponseWriter,ctx.Request)
		}
		_,ok := ctx.Input.Session("id").(int)
		if !ok && ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302,"/login")
		}
	}

	beego.InsertFilter("/",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/login",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/logout",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/index",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/player_4",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/status",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/status_dvr",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/addtestsource",beego.BeforeRouter,FilterUser)
	
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/player_4", &controllers.Player_4Controller{})
	beego.Router("/status", &controllers.StatusController{})
	beego.Router("/status_dvr", &controllers.Status_dvrController{})
	beego.Router("/addtestsource", &controllers.AddTestSourceController{})
	
	//beego.Router("/ws", &controllers.WebSocketController{})
}
