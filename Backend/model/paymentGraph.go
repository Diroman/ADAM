package model

type PaymentRequest struct {
	ContractRate float64 `json:"contractRate"`
	LastPayment  float64 `json:"lastPayment"`
	LoanAmount   int     `json:"loanAmount"`
	Payment      int     `json:"payment"`
	Term         int     `json:"term"`
}

type PaymentResponse struct {
	Payments []Payments `json:"payments"`
}

type Payments struct {
	Order      int `json:"order"`
	Percent    int `json:"percent"`
	Debt       int `json:"debt"`
	Payment    int `json:"payment"`
	BalanceOut int `json:"balanceOut"`
}

type PayiOS struct {
	Order      int `json:"order"`
	BalanceOut int `json:"balanceOut"`
	Payment    int `json:"payment"`
}

type PaymentiOS struct {
	Payments []PayiOS
}

func PaymentRespToPaymentiOS(response PaymentResponse) PaymentiOS {
	result := response.Payments

	var payments []PayiOS
	for _, pay := range result {
		payments = append(payments, PayiOS{
			Order:      pay.Order,
			BalanceOut: pay.BalanceOut,
			Payment:    pay.Payment,
		})
	}
	return PaymentiOS{Payments: payments}
}
