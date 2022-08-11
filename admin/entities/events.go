package entities

import "github.com/JenswBE/go-commerce/entities"

type EventsListData struct {
	BaseData
	Events []entities.Event
}

type EventsFormData struct {
	BaseData
	IsNew bool
	Event entities.Event
}
