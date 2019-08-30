package permission

import (
	"fmt"
	"github.com/casbin/casbin"
)

var e *casbin.Enforcer

func init()  {
	e = casbin.NewEnforcer("tools/permission/model.conf", "tools/permission/policy.csv")
	fmt.Println("e:", e)
}

func EnForcerInstance() *casbin.Enforcer  {
	return e
}
