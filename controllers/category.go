package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/iamMarkchu/goldfish/controllers/requests"
	"github.com/iamMarkchu/goldfish/models"
	"github.com/iamMarkchu/goldfish/services"
	"github.com/iamMarkchu/goldfish/tools"
	"net/http"
	"strconv"
)

//  CategoryController operations for Category
type CategoryController struct {
	ApiController
	*services.CategoryService
}

func (c *CategoryController) Prepare() {
	c.CategoryService = services.NewCategoryService(models.NewCategory())
}

// URLMapping ...
func (c *CategoryController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Category
// @Param	name				formData 	string	true		"name"
// @Param	alias_name			formData 	string	false		"alias_name"
// @Param	display_order		formData 	int		true		"display_order"
// @Param	parent_id			formData 	int		true		"parent_id"
// @Success 201 {int} models.Category
// @Failure 403 body is empty
// @router / [post]
func (c *CategoryController) Post() {
	var (
		scr      requests.StoreCategoryRequest
		err      error
		valid    = validation.Validation{}
		isValid  bool
		category *models.Category
	)
	if err = c.ParseForm(&scr); err != nil {
		c.JsonReturn("创建类别接口解析参数失败"+err.Error(), nil, http.StatusExpectationFailed)
		return
	}
	if isValid, _ = valid.Valid(&scr); !isValid {
		c.JsonReturn("创建类别接口验证参数失败", tools.GetErrorMap(valid.Errors), http.StatusExpectationFailed)
		return
	}
	userIdStr := c.Ctx.Input.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)
	if category, err = c.CategoryService.Create(scr, userId); err != nil {
		c.JsonReturn("创建类别接口保存失败:"+err.Error(), nil, http.StatusExpectationFailed)
		return
	}
	c.JsonReturn("创建类别接口", category, http.StatusOK)
}

// GetOne ...
// @Title Get One
// @Description get Category by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Category
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CategoryController) GetOne() {
	var (
		idStr    string
		id       int
		category *models.Category
		err      error
	)
	idStr = c.Ctx.Input.Param(":id")
	id, _ = strconv.Atoi(idStr)
	if category, err = c.CategoryService.GetCategoryById(id); err != nil {
		c.JsonReturn("获取类别接口报错:"+err.Error(), category, http.StatusOK)
		return
	}
	c.JsonReturn("获取类别接口", category, http.StatusOK)
}

// GetAll ...
// @Title Get All
// @Description get Category
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Category
// @Failure 403
// @router / [get]
func (c *CategoryController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Category
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Category	true		"body for Category content"
// @Success 200 {object} models.Category
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CategoryController) Put() {
	var (
		scr      requests.StoreCategoryRequest
		err      error
		valid    = validation.Validation{}
		isValid  bool
		category *models.Category
	)
	if err = c.ParseForm(&scr); err != nil {
		c.JsonReturn("更新类别接口解析参数失败"+err.Error(), nil, http.StatusExpectationFailed)
		return
	}
	if isValid, _ = valid.Valid(&scr); !isValid {
		c.JsonReturn("更新类别接口验证参数失败", tools.GetErrorMap(valid.Errors), http.StatusExpectationFailed)
		return
	}
	userIdStr := c.Ctx.Input.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)
	idStr := c.Ctx.Input.Param(":id")
	id, _ :=  strconv.Atoi(idStr)
	if category, err = c.CategoryService.Update(scr, userId, id); err != nil {
		c.JsonReturn("更新类别接口保存失败:"+err.Error(), nil, http.StatusExpectationFailed)
		return
	}
	c.JsonReturn("更新类别接口", category, http.StatusOK)
}

// Delete ...
// @Title Delete
// @Description delete the Category
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CategoryController) Delete() {

}
