package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/go-faker/faker/v4"
)

// send_data
func sendMessage(eventName string, data []byte) {
	url := "http://localhost:3000/?event_name=" + eventName
	_, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func maybeSetField(probability float64) bool {
	return rand.Float64() < probability
}

// view_home
type ViewHome struct {
	EventName string `json:"event_name,omitempty"`
	UserID    string `faker:"username" json:"user_id,omitempty"`
	DeviceID  string `faker:"uuid_hyphenated" json:"device_id,omitempty"`
	Platform  string `faker:"oneof: web, ios, android" json:"platform,omitempty"`
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
	From int `faker:"boundary_start=1980, boundary_end=2000" json:"from,omitempty"`
	To   int `faker:"boundary_start=2001, boundary_end=2022" json:"to,omitempty"`
}

type PriceRange struct {
	From int `faker:"boundary_start=1000000, boundary_end=5000000" json:"from,omitempty"`
	To   int `faker:"boundary_start=5000000, boundary_end=50000000" json:"to,omitempty"`
}

type Filter struct {
	Segment   string      `faker:"oneof: SUV, RV, 버스, 경차, 소형차, 준중형차, 대형차, 스포츠카, 승합차" json:"segment,omitempty"`
	Fuel      string      `faker:"oneof: 가솔린, 디젤, LPG, 하이브리드, 전기" json:"fuel,omitempty"`
	Region    string      `faker:"oneof: 서울, 경기, 인천, 강원, 충북, 충남, 대전, 경북, 경남, 대구, 부산, 울산, 전북, 전남, 광주, 제주" json:"region,omitempty"`
	Color     string      `faker:"oneof: 검정, 흰색, 은색, 회색, 빨강, 주황, 노랑, 초록, 파랑, 남색, 보라, 갈색, 기타" json:"color,omitempty"`
	YearType  *YearRange  `json:"year_type,omitempty"`
	PriceType *PriceRange `json:"price_type,omitempty"`
}

type Parameters struct {
	Filter Filter `json:"filter,omitempty"`
}

type ViewSearchResult struct {
	EventName  string      `json:"event_name,omitempty"`
	UserID     string      `faker:"username" json:"user_id,omitempty"`
	DeviceID   string      `faker:"uuid_hyphenated" json:"device_id,omitempty"`
	Platform   string      `faker:"oneof: web, ios, android" json:"platform,omitempty"`
	Parameters *Parameters `json:"parameters,omitempty"`
}

func GenerateFilter(filter *Filter) Filter {
	if maybeSetField(0.9) {
		err := faker.FakeData(&filter)
		if err != nil {
			panic(err)
		}
		if maybeSetField(0.8) {
			yearRange := YearRange{}
			err := faker.FakeData(&yearRange)
			if err != nil {
				panic(err)
			}
			filter.YearType = &yearRange
		}
		if maybeSetField(0.8) {
			priceRange := PriceRange{}
			err := faker.FakeData(&priceRange)
			if err != nil {
				panic(err)
			}
			filter.PriceType = &priceRange
		}
	} else {
		return *filter
	}
	return *filter
}

func GenerateViewSearchResult() (data []byte) {
	filter := Filter{}
	generatedFilter := GenerateFilter(&filter)
	ViewSearchResult := &ViewSearchResult{}
	err := faker.FakeData(&ViewSearchResult)
	if err != nil {
		panic(err)
	}
	ViewSearchResult.EventName = "view_searchResult"
	ViewSearchResult.Parameters.Filter = generatedFilter

	data, err = json.Marshal(ViewSearchResult)
	if err != nil {
		panic(err)
	}
	return data
}

func main() {
	viewHomeData := GenerateViewHome()
	sendMessage("view_home", viewHomeData)
	// viewSearchResultData := GenerateViewSearchResult()
	// sendMessage(viewSearchResultData)
}
