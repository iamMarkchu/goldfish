package requests

type LoginRequest struct {
	UserName string `form:"userName" valid:"Required;MaxSize(50)"`
	Password string `form:"password" valid:"Required;MaxSize(100)"`
}

type RegisterRequest struct {
	UserName string `form:"userName" valid:"Required;MaxSize(50)"`
	Password string `form:"password" valid:"Required;MaxSize(100)"`
	RePassword string `form:"rePassword" valid:"Required;MaxSize(100)"`
	Mobile     string `form:"mobile" valid:"Required;Mobile"`
}
