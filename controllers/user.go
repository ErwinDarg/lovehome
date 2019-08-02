package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *UserController) Reg() {
	resp := make(map[string]interface{})

	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	beego.Info(resp)
	resp["errno"] = 4001
	resp["errmsg"] = "查询失败"
	this.RetData(resp)
}