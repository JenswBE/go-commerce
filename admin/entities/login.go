package entities

type LoginFailedTemplate struct {
	BaseData
	Reason string
}

func (t LoginFailedTemplate) GetTemplateName() string {
	return "loginFailed"
}
