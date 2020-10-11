package model

import (
	"fmt"
	"strconv"
	"strings"
)

type ModelNameInternal struct {
	Absentee bool   `json:"absentee"`
	Alias    string `json:"alias"`
	ID       int    `json:"id"`
	Prefix   string `json:"prefix"`
	Title    string `json:"title"`
	TitleRus string `json:"titleRus"`
}

type ModelInternal struct {
	Alias                string     `json:"alias"`
	Bodies               []*Body    `json:"bodies"`
	Brand                *Brand     `json:"brand"`
	CarID                string     `json:"carId"`
	ColorsCount          int        `json:"colorsCount"`
	Count                int        `json:"count"`
	HasSpecialPrice      bool       `json:"hasSpecialPrice"`
	ID                   int        `json:"id"`
	MetallicPay          int        `json:"metallicPay"`
	Minprice             int        `json:"minprice"`
	Model                *ModelName `json:"model"`
	RenderPhotos         []Photo    `json:"renderPhotos"`
	Photo                string     `json:"photo"`
	Prefix               string     `json:"prefix"`
	PremiumPriceSpecials []float32  `json:"premiumPriceSpecials"`
	Title                string     `json:"title"`
	TitleRus             string     `json:"titleRus"`
}

type ReturnCar struct {
	Absentee bool          `json:"absentee"`
	Alias    string        `json:"alias"`
	Country  *Country      `json:"country"`
	Logo     string        `json:"logo"`
	Models   ModelInternal `json:"models"`
	Title    string        `json:"title"`
	TitleRus string        `json:"titleRus"`
	Prefix   string        `json:"prefix"`
	Info     string        `json:"info"`
}

type CarInternal struct {
	Absentee bool                     `json:"absentee"`
	Alias    string                   `json:"alias"`
	Country  *Country                 `json:"country"`
	Logo     string                   `json:"logo"` // ???????
	Models   map[string]ModelInternal `json:"models"`
	Title    string                   `json:"title"`
	TitleRus string                   `json:"titleRus"`
	Prefix   string
}

type CarListInternal struct {
	Cars map[string]CarInternal
}

type AllInfo struct {
	Country              string    `json:"country"`
	Title                string    `json:"Title"`
	Bodies               []*Body   `json:"Bodies"`
	Count                int       `json:"count"`
	HasSpecialPrice      bool      `json:"hasSpecialPrice"`
	Minprice             int       `json:"minprice"`
	RenderPhotos         []Photo   `json:"renderPhotos"`
	Photo                string    `json:"photo"`
	PremiumPriceSpecials []float32 `json:"premiumPriceSpecials"`
	PrettyPrice          string    `json:"prettyPrice"`
	Info                 string    `json:"info"`
}

type CariOS struct {
	Cars []AllInfo
}

func CarToInternal(list CarList) *CarListInternal {
	internalList := CarListInternal{Cars: map[string]CarInternal{}}

	for _, car := range list.List {
		modelList := map[string]ModelInternal{}

		for _, model := range car.Models {
			var renderPhoto []Photo

			for _, value := range model.RenderPhotos {
				renderPhoto = append(renderPhoto, Photo{
					Height: value.Height,
					Path:   value.Path,
					Width:  value.Width,
				})
			}

			internalModels := ModelInternal{
				Alias:                model.Alias,
				Bodies:               model.Bodies,
				Brand:                model.Brand,
				CarID:                model.CarID,
				ColorsCount:          model.ColorsCount,
				Count:                model.Count,
				HasSpecialPrice:      model.HasSpecialPrice,
				ID:                   model.ID,
				MetallicPay:          model.MetallicPay,
				Minprice:             model.Minprice,
				Model:                model.Model, //
				Photo:                model.Photo,
				Prefix:               model.Prefix,
				PremiumPriceSpecials: model.PremiumPriceSpecials,
				Title:                model.Title,
				TitleRus:             model.TitleRus,
				RenderPhotos:         renderPhoto,
			}

			modelTitle := strings.ToUpper(model.Title)
			if strings.Index(modelTitle, "СЕРИИ") != -1 || strings.Index(modelTitle, "СЕРИЯ") != -1 {
				modelTitle = strings.Split(modelTitle, " ")[0]
			}

			modelList[modelTitle] = internalModels
		}
		internalCar := CarInternal{
			Absentee: car.Absentee,
			Alias:    car.Alias,
			Country: &Country{
				Code:  car.Country.Code,
				ID:    car.Country.ID,
				Title: car.Country.Title,
			},
			Logo:     car.Logo,
			Models:   modelList,
			Title:    car.Title,
			TitleRus: car.TitleRus,
		}

		internalList.Cars[strings.ToUpper(car.Title)] = internalCar
	}

	return &internalList
}

func PredictToiOS(retCar *ReturnCar) *CariOS {
	var renderPhoto []Photo

	renderPhoto = append(renderPhoto, Photo{
		Height: 0,
		Path:   retCar.Models.Photo,
		Width:  0,
	})

	for _, body := range retCar.Models.Bodies {
		renderPhoto = append(renderPhoto, Photo{
			Height: 0,
			Path:   body.Photo,
			Width:  0,
		})
	}

	for _, photo := range retCar.Models.RenderPhotos {
		if photo.Path != "" {
			renderPhoto = append(renderPhoto, photo)
		}
	}

	retCar.Models.RenderPhotos = renderPhoto

	cars := AllInfo{
		Country:              retCar.Country.Title,
		Title:                retCar.Title,
		Bodies:               retCar.Models.Bodies,
		Count:                retCar.Models.Count,
		HasSpecialPrice:      retCar.Models.HasSpecialPrice,
		Minprice:             retCar.Models.Minprice,
		RenderPhotos:         retCar.Models.RenderPhotos,
		Photo:                retCar.Models.Photo,
		PremiumPriceSpecials: retCar.Models.PremiumPriceSpecials,
		PrettyPrice:          prettyPrice(retCar.Models.Minprice),
		Info:                 retCar.Info,
	}

	return &CariOS{Cars: []AllInfo{cars}}
}

func GetBestProb(selector *CarResponse) string {
	var bestKey string
	var bestProb float32 = 0

	for key, value := range selector.Probabilities {
		if value > bestProb {
			bestProb = value
			bestKey = key
		}
	}

	return bestKey
}

func prettyPrice(price int) string {
	priceString := []byte(strconv.Itoa(price))

	var prettyString []byte
	for i := len(priceString) - 1; i >= 0; i-- {
		if i != (len(priceString)-1) && i%3 == 0 {
			prettyString = append([]byte{' '}, prettyString...)
		}
		prettyString = append([]byte{priceString[i]}, prettyString...)
	}

	return fmt.Sprintf("%s р", prettyString)
}
