package entities

import (
	"fmt"
	"time"
)

const TimeFormatDate = "2006-01-02"

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

func parseDateString(dateString string) (time.Time, error) {
	date, err := time.Parse(TimeFormatDate, dateString)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date %s into format %s: %w", dateString, TimeFormatDate, err)
	}
	return date, nil
}
