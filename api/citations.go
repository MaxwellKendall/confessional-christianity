package api

// Citation is a row in the DB Table Citations
type Citation struct {
	ID           string
	ConfessionID string
	HeadingID    string
	PassageID    string
	ReferenceID  string
	Scripture    []string
	Tags         []string
}
