package model

type Photo struct {
	Height int    `json:"height"`
	Path   string `json:"path"`
	Width  int    `json:"width"`
}

type Country struct {
	Code  string `json:"code"`
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Brand struct {
	Absentee   bool     `json:"absentee"`
	Alias      string   `json:"alias"`
	Country    *Country `json:"country"`
	ID         int      `json:"id"`
	IsOutbound bool     `json:"isOutbound"`
	Logo       string   `json:"logo"`
	Title      string   `json:"title"`
	TitleRus   string   `json:"titleRus"`
}

type Body struct {
	Alias     string `json:"alias"`
	Doors     int    `json:"doors"`
	Photo     string `json:"photo"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	TypeTitle string `json:"typeTitle"`
}

type ModelName struct {
	Absentee bool   `json:"absentee"`
	Alias    string `json:"alias"`
	ID       int    `json:"id"`
	Prefix   string `json:"prefix"`
	Title    string `json:"title"`
	TitleRus string `json:"titleRus"`
}

type Model struct {
	Absentee             bool             `json:"absentee"`
	Alias                string           `json:"alias"`
	Bodies               []*Body          `json:"bodies"`
	Brand                *Brand           `json:"brand"`
	CarID                string           `json:"carId"`
	ColorsCount          int              `json:"colorsCount"`
	Count                int              `json:"count"`
	HasSpecialPrice      bool             `json:"hasSpecialPrice"`
	ID                   int              `json:"id"`
	MetallicPay          int              `json:"metallicPay"`
	Minprice             int              `json:"minprice"`
	Model                *ModelName       `json:"model"`
	OwnTitle             string           `json:"ownTitle"`
	PearlPay             int              `json:"pearlPay"`
	Photo                string           `json:"photo"`
	Prefix               string           `json:"prefix"`
	PremiumPriceSpecials []float32        `json:"premiumPriceSpecials"`
	RenderPhotos         map[string]Photo `json:"renderPhotos"`
	SizesPhotos          struct {
		Width250 string `json:"width250"`
	} `json:"sizesPhotos"`
	SpecmetallicPay int    `json:"specmetallicPay"`
	Title           string `json:"title"`
	TitleRus        string `json:"titleRus"`
	TransportType   struct {
		Alias string `json:"alias"`
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"transportType"`
}

type Car struct {
	Absentee           bool          `json:"absentee"`
	Alias              string        `json:"alias"`
	Country            *Country      `json:"country"`
	CurrentCarCount    int           `json:"currentCarCount"`
	CurrentModelsTotal int           `json:"currentModelsTotal"`
	Generations        []interface{} `json:"generations"`
	ID                 int           `json:"id"`
	IsOutbound         bool          `json:"isOutbound"`
	Logo               string        `json:"logo"`
	Models             []*Model      `json:"models"`
	Title              string        `json:"title"`
	TitleRus           string        `json:"titleRus"`
}

type CarList struct {
	List []*Car `json:"list"`
}
