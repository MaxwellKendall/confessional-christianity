package api

// Chapter is a WCF chapter
type Chapter struct {
	Title      string            `json:"title"`
	Number     int               `json:"number"`
	Paragraphs []string          `json:"paragraphs"`
	Proofs     map[string]string `json:"proofs"`
}
