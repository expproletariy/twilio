package enums

type QueryStatus string

func (q QueryStatus) String() string {
	return string(q)
}

const (
	QueryStatusPendingReview QueryStatus = "pending-review"
	QueryStatusReviewed      QueryStatus = "reviewed"
	QueryStatusDiscarded     QueryStatus = "discarded"
)
