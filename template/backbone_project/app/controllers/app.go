package controllers

import "github.com/revel/revel"

import "github.com/deepglint/backbone/app"

type App struct {
	*revel.Controller
}

func (c App) GetVersion() revel.Result {
	return c.RenderJson(struct {
		Code     int    `json:"code"`
		Msg      string `json:"msg"`
		Redirect string `json:"redirect"`
		Data     string `json:"data"`
	}{
		0,
		app.Version,
		"",
		"",
	})
}

func (c App) Index() revel.Result {
	return c.Render()
}
