package dto

type IdMongo struct {
	Result ResultMongo `json:"result"`
}

type ResultMongo struct {
	Lob_id string `json:"lob_id"`
}
