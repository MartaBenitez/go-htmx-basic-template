package dto

type SectionsToEdit struct {
	Query  string         `json:"query"`
	Module string         `json:"module"`
	Result ResultSections `json:"result"`
}

type ResultSections struct {
	BusinessLine               string `json:"BusinessLine"`
	BusinessLineData           string `json:"BusinessLineData"`
	CommercialNetworkAttribute string `json:"CommercialNetworkAttribute"`
	ProductPaymentMethod       string `json:"ProductPaymentMethod"`
	ProductRenewalCycle        string `json:"ProductRenewalCycle"`
	RenewalParameter           string `json:"RenewalParameter"`
}
