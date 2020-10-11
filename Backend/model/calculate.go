package model

import "strconv"

type CalculateRequest struct {
	ClientTypes       []string `json:"clientTypes"`
	Cost              string   `json:"cost"`
	InitialFee        int      `json:"initialFee"`
	KaskoValue        int      `json:"kaskoValue"`
	Language          string   `json:"language"`
	ResidualPayment   float64  `json:"residualPayment"`
	SettingsName      string   `json:"settingsName"`
	SpecialConditions []string `json:"specialConditions"`
	Term              int      `json:"term"`
}

type CalculateRequestiOS struct {
	ClientTypes       []string `json:"clientTypes"`
	Cost              int      `json:"cost"`
	InitialFee        int      `json:"initialFee"`
	KaskoValue        int      `json:"kaskoValue"`
	Language          string   `json:"language"`
	ResidualPayment   float64  `json:"residualPayment"`
	SettingsName      string   `json:"settingsName"`
	SpecialConditions []string `json:"specialConditions"`
	Term              int      `json:"term"`
}

type CalculateResponse struct {
	Program Program `json:"program"`
	Ranges  Ranges  `json:"ranges"`
	Result  Result  `json:"result"`
}

type Cost struct {
	Filled bool `json:"filled"`
	Max    int  `json:"max"`
	Min    int  `json:"min"`
}

type Program struct {
	Cost        Cost   `json:"cost"`
	ID          string `json:"id"`
	ProgramName string `json:"programName"`
	ProgramURL  string `json:"programUrl"`
	RequestURL  string `json:"requestUrl"`
}

type InitialFee struct {
	Filled bool `json:"filled"`
	Max    int  `json:"max"`
	Min    int  `json:"min"`
}

type ResidualPayment struct {
	Filled bool `json:"filled"`
	Max    int  `json:"max"`
	Min    int  `json:"min"`
}

type Term struct {
	Filled bool `json:"filled"`
	Max    int  `json:"max"`
	Min    int  `json:"min"`
}

type Ranges struct {
	Cost            Cost            `json:"cost"`
	InitialFee      InitialFee      `json:"initialFee"`
	ResidualPayment ResidualPayment `json:"residualPayment"`
	Term            Term            `json:"term"`
}

type Result struct {
	ContractRate float64 `json:"contractRate"`
	KaskoCost    int     `json:"kaskoCost"`
	LastPayment  float64 `json:"lastPayment"`
	LoanAmount   int     `json:"loanAmount"`
	Payment      int     `json:"payment"`
	Subsidy      float64 `json:"subsidy"`
	Term         int     `json:"term"`
}

type CalculateiOS struct {
	ContractRate float64 `json:"contractRate"`
	KaskoCost    int     `json:"kaskoCost"`
	Payment      int     `json:"payment"`
	Term         int     `json:"term"`
}

func CalcResponseToCalciOS(response CalculateResponse) CalculateiOS {
	result := response.Result
	return CalculateiOS{
		ContractRate: result.ContractRate,
		KaskoCost:    result.KaskoCost,
		Payment:      result.Payment,
		Term:         result.Term,
	}
}

func CalcReqToiOS(request CalculateRequest) CalculateRequestiOS {
	cost, _ := strconv.Atoi(request.Cost)
	return CalculateRequestiOS{
		ClientTypes:       request.ClientTypes,
		Cost:              cost,
		InitialFee:        request.InitialFee,
		KaskoValue:        request.KaskoValue,
		Language:          request.Language,
		ResidualPayment:   request.ResidualPayment,
		SettingsName:      request.SettingsName,
		SpecialConditions: request.SpecialConditions,
		Term:              request.Term,
	}
}