package controllers

import (
	"JivoxCensusGo/api"
	"github.com/astaxie/beego"
)

// Operations about object
type ScrapeController struct {
	beego.Controller
}

// // @Title Get
// // @Description find object by objectid
// // @Param	objectId		path 	string	true		"the objectid you want to get"
// // @Success 200 {object} models.Object
// // @Failure 403 :objectId is empty
// // @router /:objectId [get]
// func (o *ObjectController) Get() {
// 	objectId := o.Ctx.Input.Param(":objectId")
// 	if objectId != "" {
// 		ob, err := models.GetOne(objectId)
// 		if err != nil {
// 			o.Data["json"] = err.Error()
// 		} else {
// 			o.Data["json"] = ob
// 		}
// 	}
// 	o.ServeJSON()
// }

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ScrapeController) GetAll() {
	o.Data["json"] = api.Scrape()
	o.ServeJSON()
}
