package controller

const (
	CodeSuccess = 1000 + iota
	CodeError
	CodeParam
	CodeSignUp
	CodeLogin
	CodeServerBusy
)

var codeMsgMap = map[int]string{
	CodeSuccess:    "success",
	CodeError:      "error",
	CodeParam:      "param error",
	CodeSignUp:     "sign up error",
	CodeLogin:      "login error",
	CodeServerBusy: "server busy",
}

func GetCodeMsg(code int) string {
	msg, ok := codeMsgMap[code]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}
