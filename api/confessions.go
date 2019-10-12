package api

// Confession is a row in the DB Table Confessions
type Confession struct {
	id       string
	title    string
	authors  []string
	date     string
	location string
	summary  string
}

// Chapter represents a chapter of the Westminster Confession of Faith
type Chapter struct {
	Title      string            `json:"title"`
	Number     int               `json:"number"`
	Paragraphs []string          `json:"paragraphs"`
	Proofs     map[string]string `json:"proofs"`
}
