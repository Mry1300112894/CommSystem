package message

const (
	LoginMesType  = "loginMes"
	ResMesType    = "loginResMes"
	EnrollMesType = "enrollMes"
	UserStatusType = "userStatusMes"
	SmsMesType = "smsMesType"
)

// 定义用户的状态
const (
	Online = iota	// 在线
	Busy			// 忙碌
	Offline			// 离线
)

type Message struct {
	Type string	`json:"type"`	// 发送信息的类型
	Date string	`json:"date"`	// 发送的信息
}

// 返回信息
type ResMes struct {
	Code int					`json:"code"`
	Error string				`json:"error"`
	OnlineUserList []string		`json:"online_user_list"`
}

// 用户信息
type User struct {
	AcoNum string	`json:"aco_num"`	// 账户
	AcoPwd string	`json:"aco_pwd"`	// 密码
	AcoName string	`json:"aco_name"`	// 用户名
	SafePwd string	`json:"safe_pwd"`	// 安全密码
	Status int		`json:"status"`		// 用户状态
}

// 用户状态
type UserStatus struct {
	AcoNum string	`json:"aco_num"`
	Status int		`json:"status"`
}

// 聊天信息
type SmsMes struct {
	Content string	`json:"content"`	// 用户发送的内容
	User			// 保存用户的账号，为了后续的开发，继承 User
}

/*// 注册
type Enroll struct {
	AcoNum string	`json:"aco_num"`	// 账户
	AcoPwd string	`json:"aco_pwd"`	// 密码
	AcoName string	`json:"aco_name"`	// 用户名
	SafePwd string	`json:"safe_pwd"`	// 安全密码
}

// 登录
type Login struct {
	AcoNum string	`json:"aco_num"`	// 账户
	AcoPwd string	`json:"aco_pwd"`	// 密码
}*/