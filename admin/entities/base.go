package entities

type WithBaseData interface {
	SetMessages(messages []Message)
}

type BaseData struct {
	Title      string
	ParentPath string
	Messages   []Message
}

func (b *BaseData) SetMessages(messages []Message) {
	b.Messages = messages
}
