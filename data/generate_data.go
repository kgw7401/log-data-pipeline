package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-faker/faker/v4"
)

// send_data
func sendMessage(data []byte) {
	_, err := http.Post("http://localhost:3000", "application/json", bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
}

// view_home
type ViewHome struct {
	EventName string `json:"event_name"`
	UserID    string `faker:"username" json:"user_id"`
	DeviceID  string `faker:"uuid_hyphenated" json:"device_id"`
	Platform  string `faker:"oneof: web, ios, android" json:"platform"`
}

func GenerateViewHome() (data []byte) {
	ViewHome := ViewHome{}
	err := faker.FakeData(&ViewHome)
	if err != nil {
		panic(err)
	}
	ViewHome.EventName = "view_home"

	data, err = json.Marshal(ViewHome)
	if err != nil {
		panic(err)
	}
	return data
}

// view_searchResult
type YearRange struct {
	From int `faker:"boundary_start=1980, boundary_end=2000" json:"from"`
	To   int `faker:"boundary_start=2001, boundary_end=2022" json:"to"`
}

type PriceRange struct {
	From int `faker:"boundary_start=1000000, boundary_end=5000000" json:"from"`
	To   int `faker:"boundary_start=5000000, boundary_end=50000000" json:"to"`
}

type Filter struct {
	Segment   string     `faker:"oneof: SUV, RV, 버스, 경차, 소형차, 준중형차, 대형차, 스포츠카, 승합차" json:"segment"`
	Fuel      string     `faker:"oneof: 가솔린, 디젤, LPG, 하이브리드, 전기" json:"fuel"`
	Region    string     `faker:"oneof: 서울, 경기, 인천, 강원, 충북, 충남, 대전, 경북, 경남, 대구, 부산, 울산, 전북, 전남, 광주, 제주" json:"region"`
	Color     string     `faker:"oneof: 검정, 흰색, 은색, 회색, 빨강, 주황, 노랑, 초록, 파랑, 남색, 보라, 갈색, 기타" json:"color"`
	YearType  YearRange  `json:"year_type"`
	PriceType PriceRange `json:"price_type"`
}

type Parameters struct {
	Filter Filter `json:"filter"`
}

type ViewSearchResult struct {
	EventName  string     `json:"event_name"`
	UserID     string     `faker:"username" json:"user_id"`
	DeviceID   string     `faker:"uuid_hyphenated" json:"device_id"`
	Platform   string     `faker:"oneof: web, ios, android" json:"platform"`
	Parameters Parameters `json:"parameters"`
}

func GenerateViewSearchResult() (data []byte) {
	ViewSearchResult := ViewSearchResult{}
	err := faker.FakeData(&ViewSearchResult)
	if err != nil {
		panic(err)
	}
	ViewSearchResult.EventName = "view_searchResult"

	data, err = json.Marshal(ViewSearchResult)
	if err != nil {
		panic(err)
	}
	return data
}

func main() {
	viewHomeData := GenerateViewHome()
	viewSerachResult := GenerateViewSearchResult()

	sendMessage(viewHomeData)
	sendMessage(viewSerachResult)
}
