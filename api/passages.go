package api

// Passage is a row in the DB Table Passages
type Passage struct {
	id            string
	confession    string
	heading       string
	passageNumber string
	headingNumber string
	passage       string
}
