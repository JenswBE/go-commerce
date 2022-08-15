package entities

type LogoutSuccessfulTemplate struct {
	BaseData
}

func (t LogoutSuccessfulTemplate) GetTemplateName() string {
	return "logoutSuccessful"
}
