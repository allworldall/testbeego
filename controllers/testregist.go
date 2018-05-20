package controllers

import "github.com/astaxie/beego"

type ObjectController struct {
	beego.Controller
}


func (c *ObjectController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
}

//@router /staticblock/:key [get]
func (this *ObjectController) StaticBlock() {

}