package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lovehome/models"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}


func (c *AreaController) GetArea() {
	beego.Info("connect success")
	resp := make(map[string]interface{})
	resp["errno"] = 0
	resp["errmsg"] = "OK"
	defer c.RetData(resp)
	// 从session中拿数据

	//从数据库中拿到area数据
	var areas []models.Area

	o := orm.NewOrm()
	num, err := o.QueryTable("area").All(&areas)
	if err != nil{
		beego.Error(err)
		resp["errno"] = 4001
		resp["errmsg"] = "查询数据错误"
		return
	}else {
		beego.Info(num)
	}



	// 打包成json返回给前端
}

