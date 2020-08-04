package message

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
)
type Message struct {
	Type string	`json:"type"`
	Data string	`json:"data"`
}

type LoginMes struct {
	UserId int	`json:"user_id"`
	UserPwd string	`json:"user_pwd"`
	UserName string	`json:"user_name"`
}

type LoginResMes struct {
	Code int	`json:"code"`
	Error string	`json:"error"`
	
}
