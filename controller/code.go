package controller

type ResCode int64
const (
	CodeSuccess ResCode=1000+iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServeBusy

	CodeNeedLogin
	CodeInvalidToken
)
var codeMsgMap=map[ResCode]string{
	CodeSuccess: "success",
	CodeInvalidParam:"请求参数错误",
	CodeUserExist: "用户已存在",
	CodeUserNotExist:"用户不存在",
	CodeInvalidPassword:"用户名或者密码错误",
	CodeServeBusy:"服务繁忙",

	CodeNeedLogin:"需要登录",
	CodeInvalidToken:"无效Token",
}

func (c ResCode) Msg() string{
	msg,ok:=codeMsgMap[c]
	if !ok{
		msg=codeMsgMap[CodeServeBusy]
	}
	return msg
}