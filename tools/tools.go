package tools

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

func CheckError(err error, msg string) {
	if err != nil {
		logs.Info(msg, err)
	}
}

func GetErrorMap(errs []*validation.Error) map[string]string {
	var errorMaps = make(map[string]string)
	for _, err := range errs {
		errorMaps[err.Key] = err.Message
	}
	return errorMaps
}

// md5加密
func MD5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func JsonReturn(message string, result interface{}, code int) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"result":  result,
		"code":    code,
	}
}
