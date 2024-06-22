package models

import "time"

type Event struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Location string `json:"location"`
	DateTime time.Time `json:"dateTime"`
	UserId int `json:"userId"`
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
