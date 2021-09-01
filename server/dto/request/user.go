package request

// User register structure
type Register struct {
	Username string `json:"userName"`
	Password string `json:"passWord"`
	NickName string `json:"nickName" gorm:"default:''"`
	Avatar   string `json:"avatar" gorm:"default:''"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// Modify password structure
type ChangePasswordStruct struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserRole struct {
	UserID int   `json:"userId"` // 用户ID
	RoleID []int `json:"roleId"` // 角色ID
}
