package model

type ArchiveEntry struct {
	Word        string `json:"word"`
	Meaning     string `json:"translation"`
	StartDate   string `json:"start_date"`
	HitsCount   int    `json:"hits_count"`
	ArchiveDate string `json:"archive_date"`
}
