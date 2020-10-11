package model

import "time"

type CarLoanRequest struct {
	Comment         string        `json:"comment"`
	CustomerParty   CustomerParty `json:"customer_party"`
	Datetime        time.Time     `json:"datetime"`
	InterestRate    float64       `json:"interest_rate"`
	RequestedAmount int           `json:"requested_amount"`
	RequestedTerm   int           `json:"requested_term"`
	TradeMark       string        `json:"trade_mark"`
	VehicleCost     int           `json:"vehicle_cost"`
}
type Person struct {
	BirthDateTime          string `json:"birth_date_time"`
	BirthPlace             string `json:"birth_place"`
	FamilyName             string `json:"family_name"`
	FirstName              string `json:"first_name"`
	Gender                 string `json:"gender"`
	MiddleName             string `json:"middle_name"`
	NationalityCountryCode string `json:"nationality_country_code"`
}
type CustomerParty struct {
	Email        string `json:"email"`
	IncomeAmount int    `json:"income_amount"`
	Person       Person `json:"person"`
	Phone        string `json:"phone"`
}


type CarLoanResponse struct {
	Application Application `json:"application"`
	Datetime    time.Time   `json:"datetime"`
}
type DecisionReport struct {
	ApplicationStatus string `json:"application_status"`
	Comment           string `json:"comment"`
	DecisionDate      string `json:"decision_date"`
	DecisionEndDate   string `json:"decision_end_date"`
	MonthlyPayment    int    `json:"monthly_payment"`
}
type Application struct {
	VTBClientID    int            `json:"VTB_client_ID"`
	DecisionReport DecisionReport `json:"decision_report"`
}
