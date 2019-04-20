package api

// WCFChapter is a WCF chapter
type WCFChapter struct {
	ID         string            `json:"id"`
	Title      string            `json:"title"`
	Number     int               `json:"number"`
	Paragraphs []string          `json:"paragraphs"`
	Proofs     map[string]string `json:"proofs"`
}

// WCFService methods represent an endpoint that transports JSON over HTTP to the www
type WCFService interface {
	GetChapter(int) (WCFChapter, error)
}
