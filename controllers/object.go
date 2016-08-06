package controllers

import (
	"JivoxCensusGo/api"
	"github.com/astaxie/beego"
)

// Operations about Scrape
type ScrapeController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all state data
// @Success 200 {object} models.Census
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ScrapeController) GetAll() {
	o.Data["json"] = api.Scrape()
	o.ServeJSON()
}
