package main

// base
type Base struct {
	EventName  string   `json:"event_name,omitempty"`
	UserID     string   `faker:"username" json:"user_id,omitempty"`
	DeviceID   string   `faker:"uuid_hyphenated" json:"device_id,omitempty"`
	Platform   string   `faker:"oneof: web, ios, android" json:"platform,omitempty"`
	Ip         string   `faker:"ipv4" json:"ip,omitempty"`
	Parameters struct{} `json:"parameters,omitempty"`
}

// home_view
type HomeView struct {
	Base
	EventName  string             `json:"event_name,omitempty"`
	Parameters HomeViewParameters `json:"parameters,omitempty"`
}

type HomeViewParameters struct{}

// login_done
type LoginDone struct {
	Base
	Parameters LoginDoneParameters `json:"parameters,omitempty"`
}

type LoginDoneParameters struct {
	Method string `faker:"oneof: kakao, naver, google" json:"method,omitempty"`
}
