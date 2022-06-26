package message

// 消息常量
const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息内容
}

// 登录消息
type LoginMes struct {
	UserId   int    `json:"userId"`   // 用户ID
	UserPwd  string `json:"userPwd"`  // 用户密码
	UserName string `json:"userName"` // 用户名
}

// 登录回送消息
type LoginResMes struct {
	Code  int    `json:"code"`  // 状态码 500-未注册 200-成功
	Error string `json:"error"` // 返回错误信息
}
