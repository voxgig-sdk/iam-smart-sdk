package core

type IamSmartError struct {
	IsIamSmartError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewIamSmartError(code string, msg string, ctx *Context) *IamSmartError {
	return &IamSmartError{
		IsIamSmartError: true,
		Sdk:              "IamSmart",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *IamSmartError) Error() string {
	return e.Msg
}
