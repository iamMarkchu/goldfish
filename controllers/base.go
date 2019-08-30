package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

type HttpCode int

type JsonResponse struct {
	Code    HttpCode    `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func (c *BaseController) JsonReturn(message string, result interface{}, code int) {
	c.Data["json"] = JsonResponse{
		Message: message,
		Result:  result,
		Code:    HttpCode(code),
	}
	c.ServeJSON(true)
}
