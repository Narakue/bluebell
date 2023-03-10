package controller

type ResCode int

const (
	CodeSuccess ResCode = 1000 + iota
	CodeError
	CodeParam
	CodeSignUp
	CodeLogin
	CodeServerBusy
	CodeNotLogin
	CodeAuthError
	CodeReLogin
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:    "success",
	CodeError:      "error",
	CodeParam:      "param error",
	CodeSignUp:     "sign up error",
	CodeLogin:      "login error",
	CodeServerBusy: "server busy",
	CodeNotLogin:   "not login",
	CodeAuthError:  "auth error",
	CodeReLogin:    "please re login",
}

func GetCodeMsg(code ResCode) string {
	msg, ok := codeMsgMap[code]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}
