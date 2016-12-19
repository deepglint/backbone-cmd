package controllers

import (
	// "encoding/json"
	"bytes"
	"github.com/deepglint/backbone/app"
	"github.com/revel/revel"
	"io/ioutil"
	"net/http"
	"strings"
)

type Login struct {
	*revel.Controller
}

type ReturnMsg struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Redirect string `json:"redirect"`
	Data     string `json:"data"`
}

var adu string

func init() {
	//println(revel.Config.StringDefault("app.aduservice", "127.0.0.1:8086"))

}

func (c Login) Index() revel.Result {
	var version = app.Version
	return c.Render(version)
}

func (c Login) Logout() revel.Result {
	c.Session["username"] = ""
	return c.Redirect("login")
}

func (c Login) Login() revel.Result {
	// deside what result when user click the login button
	// 1. redirect to login place and tell him what's wrong
	// 2. into the dashboard page,new a session,store the username in the dashboardpage
	adu = revel.Config.StringDefault("aduservice", "http://127.0.0.1:8008/api")
	println(adu, "adu")
	var url = adu + "/login"
	var username = c.Request.Form.Get("username")
	var password = c.Request.Form.Get("password")
	var body = []byte(username + ":" + password)
	println(username, password)
	resp, err := http.Post(url, "application/text", bytes.NewReader(body))
	if err != nil {
		// return c.RenderError(err)
		return c.RenderJson(ReturnMsg{2, "连接服务失败", "login", ""})
	}
	s, _ := ioutil.ReadAll(resp.Body)
	if string(s) == "SUCCESS" || string(s) == "true" {
		c.Session["username"] = username
		c.Session["password"] = password
		return c.RenderJson(ReturnMsg{
			0,
			"登陆成功",
			"./",
			"",
		})
	} else {
		return c.RenderJson(ReturnMsg{1, "用户名或密码错误", "login", ""})
	}
}

func (c Login) ChangeAuth() revel.Result {
	adu = revel.Config.StringDefault("aduservice", "http://127.0.0.1:8086/api")
	var url = adu + "/changepwd"
	var oldname = c.Request.Form.Get("oldname")
	var oldpwd = c.Request.Form.Get("oldpwd")
	var newpwd = c.Request.Form.Get("newpwd")
	var body = []byte(oldname + ":" + oldpwd + ":" + newpwd)
	// old:{"oldcode":"admin:admin", "newcode":"admin:admin"}
	// var bodyjson = "{\"oldcode\":\"" + oldname + ":" + oldpwd + "\", \"newcode\":\"" + oldname + ":" + newpwd + "\"}"
	// println(bodyjson)
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		// return c.RenderError(err)
		return c.RenderJson(ReturnMsg{2, "连接服务失败", "login", ""})
	}
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.RenderJson(ReturnMsg{5, "连接服务失败", "/", ""})
	}
	println(string(s))
	if strings.EqualFold(string(s), "SUCCESS") {
		c.Session["username"] = ""
		return c.RenderJson(ReturnMsg{
			0,
			string(s),
			"./",
			"",
		})
	} else {
		return c.RenderJson(ReturnMsg{3, string(s), "login", ""})
	}
}
