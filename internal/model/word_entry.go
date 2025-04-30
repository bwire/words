package model

type WordEntry struct {
	Word      string `json:"word"`
	Meaning   string `json:"translation"`
	Progress  int    `json:"progress"`
	StartDate string `json:"start_date"`
	HitsCount int    `json:"hits_count"`
}
