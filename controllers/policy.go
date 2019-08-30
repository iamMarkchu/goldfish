package controllers

import (
	"github.com/iamMarkchu/goldfish/tools"
	"github.com/iamMarkchu/goldfish/tools/permission"
	"net/http"
)

// PolicyController operations for Policy
type PolicyController struct {
	ApiController
}

// URLMapping ...
func (c *PolicyController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Policy
// @Param	sub formData string true
// @Param   obj formData string true
// @Param   act formData string true
// @Success 201 {object} models.Policy
// @Failure 403 body is empty
// @router / [post]
func (c *PolicyController) Post() {
	var sub = c.Input().Get("sub")
	var obj = c.Input().Get("obj")
	var act = c.Input().Get("act")
	permission.EnForcerInstance().AddPolicy(sub, obj, act)
	err := permission.EnForcerInstance().SavePolicy()
	go tools.CheckError(err, "保存策略")
	c.JsonReturn("保存策略接口", nil, http.StatusOK)
}

// GetOne ...
// @Title GetOne
// @Description get Policy by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Policy
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PolicyController) GetOne() {
	c.JsonReturn("获取策略接口", permission.EnForcerInstance().GetPolicy(), http.StatusOK)
}

// GetAll ...
// @Title GetAll
// @Description get Policy
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Policy
// @Failure 403
// @router / [get]
func (c *PolicyController) GetAll() {
	c.JsonReturn("获取策略接口", permission.EnForcerInstance().GetPolicy(), http.StatusOK)
}

// Put ...
// @Title Put
// @Description update the Policy
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Policy	true		"body for Policy content"
// @Success 200 {object} models.Policy
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PolicyController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Policy
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PolicyController) Delete() {

}
