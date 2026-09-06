package entity

type RateLimitingPolicy struct {
	TenantID string
	UserID   string

	RefillRate int64
	Capacity   int64

	fallback   bool // if true, the rate limiting policy is a fallback policy
	bypassed   bool // if true, the rate limiting policy is bypassed
}
