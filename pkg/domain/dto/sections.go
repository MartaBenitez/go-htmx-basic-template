package dto

type SectionsToEdit struct {
	Query  string         `json:"query"`
	Module string         `json:"module"`
	Result ResultSections `json:"result"`
}

type ResultSections struct {
	BusinessLine               string `json:"business_line"`
	CommercialNetworkAttribute string `json:"commercial_network_attribute"`
	ProductPaymentMethod       string `json:"product_payment_method"`
	ProductRenewalCycle        string `json:"product_renewal_cycle"`
	RenewalParameter           string `json:"renewal_parameter"`
}
