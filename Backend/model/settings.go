package model

type SettingsResponse struct {
	Anchor            string              `json:"anchor"`
	ClientTypes       []string            `json:"clientTypes"`
	Cost              int                 `json:"cost"`
	InitialFee        int                 `json:"initialFee"`
	KaskoDefaultValue int                 `json:"kaskoDefaultValue"`
	Language          string              `json:"language"`
	Name              string              `json:"name"`
	OpenInNewTab      bool                `json:"openInNewTab"`
	Programs          []string            `json:"programs"`
	SpecialConditions []SpecialConditions `json:"specialConditions"`
}
type SpecialConditions struct {
	ExcludingConditions []string `json:"excludingConditions"`
	ID                  string   `json:"id"`
	IsChecked           bool     `json:"isChecked"`
	Language            string   `json:"language"`
	Name                string   `json:"name"`
}
