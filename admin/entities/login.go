package entities

type LoginTemplate struct {
	BaseData
}

func (t LoginTemplate) GetTemplateName() string {
	return "login"
}
