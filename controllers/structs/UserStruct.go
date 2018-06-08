package structs

type UserLoginParams struct {
	UserName string  `valid:"Required;MaxSize(20)" form:"user_name"`
	Password string  `valid:"Required;MinSize(6);MaxSize(16)" form:"password"`
} //用户登录参数