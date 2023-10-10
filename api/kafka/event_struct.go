package kafka

// view_home
type ViewHome struct {
	EventName string `json:"event_name,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	DeviceID  string `json:"device_id,omitempty"`
	Platform  string `json:"platform,omitempty"`
}

// view_searchResult
type YearRange struct {
	From int `json:"from,omitempty"`
	To   int `json:"to,omitempty"`
}

type PriceRange struct {
	From int `json:"from,omitempty"`
	To   int `json:"to,omitempty"`
}

type Filter struct {
	Segment   string      `json:"segment,omitempty"`
	Fuel      string      `json:"fuel,omitempty"`
	Region    string      `json:"region,omitempty"`
	Color     string      `json:"color,omitempty"`
	YearType  *YearRange  `json:"year_type,omitempty"`
	PriceType *PriceRange `json:"price_type,omitempty"`
}

type Parameters struct {
	Filter Filter `json:"filter,omitempty"`
}

type ViewSearchResult struct {
	EventName  string      `json:"event_name,omitempty"`
	UserID     string      `json:"user_id,omitempty"`
	DeviceID   string      `json:"device_id,omitempty"`
	Platform   string      `json:"platform,omitempty"`
	Parameters *Parameters `json:"parameters,omitempty"`
}
