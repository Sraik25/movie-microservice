package model

// RecordID defines a record id. Together with RecordType
// identifies unique records across all types.
type RecordID string

// RecordType defines a record id. Together with RecordID
// identifies unique records across all types.
type RecordType string

const RecordTypeMovie = RecordType("movie")

type UserID string

type RatingValue int

type Rating struct {
	RecordID   string      `json:"record_id"`
	RecordType string      `json:"record_type"`
	UserID     UserID      `json:"user_id"`
	Value      RatingValue `json:"value"`
}
