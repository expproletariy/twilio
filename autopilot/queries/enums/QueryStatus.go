package enums

type QueryStatus string

const (
	QueryStatusPendingReview QueryStatus = "pending-review"
	QueryStatusReviewed      QueryStatus = "reviewed"
	QueryStatusDiscarded     QueryStatus = "discarded"
)
