package model

import "time"

type RequestAutomation struct {
	ID          int        `json:"id"`
	Period      string     `json:"period"`
	Intiplasma  string     `json:"intiplasma"`
	Psm         string     `json:"psm"`
	Region      string     `json:"region"`
	Estate      string     `json:"estate"`
	Division    string     `json:"division"`
	Complex     string     `json:"complex"`
	Blockarea   string     `json:"blockarea"`
	CreatedDate *time.Time `json:"created_date"`
	Viewmode    string     `json:"viewmode"`
	Timeview    string     `json:"timeview"`
	Areaview    string     `json:"areaview"`
}
