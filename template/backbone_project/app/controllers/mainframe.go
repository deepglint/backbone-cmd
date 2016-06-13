package controllers

import (
	"github.com/deepglint/backbone/app"
	"github.com/revel/revel"
	//"regexp"
	//"time"
)

type MainFrame struct {
	*revel.Controller
}

func (c MainFrame) Index() revel.Result {
	//println(c.Session["hello"])
	//c.Session["hello"] = "world"
	//println(c.Session["username"])

	var username = c.Session["username"]
	var version = app.Version
	return c.Render(username, version)

	// if username != "" {
	// 	return c.Render(username, version)
	// } else {
	// 	return c.Redirect("login")
	// }
}
