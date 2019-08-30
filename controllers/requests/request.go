package requests

type LoginRequest struct {
	UserName string `form:"userName" valid:"Required;MaxSize(50)"`
	Password string `form:"password" valid:"Required;MaxSize(100)"`
}

type RegisterRequest struct {
	UserName   string `form:"userName" valid:"Required;MaxSize(50)"`
	Password   string `form:"password" valid:"Required;MaxSize(100)"`
	RePassword string `form:"rePassword" valid:"Required;MaxSize(100)"`
	Mobile     string `form:"mobile" valid:"Required;Mobile"`
}

type StoreCategoryRequest struct {
	Name         string `form:"name" valid:"Required;MaxSize(50)"`
	AliasName    string `form:"alias_name" valid:"MaxSize(50)"`
	ParentId     int    `form:"parent_id"`
	DisplayOrder uint32 `form:"display_order" valid:"Required"`
	Image        string `form:"image"`
}
