package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["JivoxCensusGo/controllers:ScrapeController"] = append(beego.GlobalControllerRouter["JivoxCensusGo/controllers:ScrapeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

}
