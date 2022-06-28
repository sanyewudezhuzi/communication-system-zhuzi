package message

// 消息常量
const (
	LoginMesType            = "LoginMes"       // 登录
	LoginResMesType         = "LoginResMes"    // 登录响应
	RegisterMesType         = "RegisterMes"    // 注册
	RegisterResMesType      = "RegisterResMes" // 注册响应
	NotifyUserStatusMesType = "NotifyUserStatusMes"
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息内容
}

const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

// 登录消息
type LoginMes struct {
	UserId   int    `json:"userId"`   // 用户ID
	UserPwd  string `json:"userPwd"`  // 用户密码
	UserName string `json:"userName"` // 用户名
}

// 登录回送消息
type LoginResMes struct {
	Code    int    `json:"code"`    // 状态码 500-未注册 200-成功
	UserIds []int  `json:"userIds"` // 保存用户ID的切片
	Error   string `json:"error"`   // 返回错误信息
}

// 注册消息
type RegisterMes struct {
	User User `json:"user"` // 类型就是User结构体
}

// 注册回送消息
type RegisterResMes struct {
	Code  int    `json:"code"` // 状态码 400-已占用 200-成功
	Error string `json:"error"`
}

// 配合服务器端推送上线通知用户状态变化消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}
