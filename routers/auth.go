package routers

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/iamMarkchu/goldfish/tools"
	"github.com/iamMarkchu/goldfish/tools/permission"
	"github.com/iamMarkchu/goldfish/tools/token"
	"net/http"
)

var auth = func(c *context.Context) {
	var (
		tokenStr string
		userId   string
		role     string
		isValid  bool
		err      error
	)
	if tokenStr, err = request.AuthorizationHeaderExtractor.ExtractToken(c.Request); err != nil {
		logs.Info("token err:", err.Error())
		c.Output.Status = http.StatusUnauthorized
		if err = c.Output.JSON(tools.JsonReturn("token解析错误:"+err.Error(), nil, http.StatusUnauthorized), true, true); err != nil {
			logs.Info("auth输出json错误:" + err.Error())
		}
		return
	} else if userId, role, isValid = token.CheckToken(tokenStr); !isValid {
		logs.Info("token err:", "token 不合法, "+tokenStr)
		c.Output.Status = http.StatusUnauthorized
		if err = c.Output.JSON(tools.JsonReturn("token不合法", nil, http.StatusUnauthorized), true, true); err != nil {
			logs.Info("auth 输出json错误:" + err.Error())
		}
		return
	}
	c.Input.SetParam("userId", userId)
	c.Input.SetParam("role", role)

	logs.Info("userId:", userId, "role:", role, "resource:", c.Request.URL.Path, "action:", c.Request.Method)
	if permit := permission.EnForcerInstance().Enforce(role, c.Request.URL.Path, c.Request.Method); !permit {
		logs.Info("权限验证失败:没权限")
		c.Output.Status = http.StatusUnauthorized
		if err = c.Output.JSON(tools.JsonReturn("权限验证失败:没权限", nil, http.StatusUnauthorized), true, true); err != nil {
			logs.Info("权限验证输出json错误:" + err.Error())
		}
		return
	}
	logs.Info("tokenStr:", tokenStr, "userId:", userId, "roles:", role)
}
