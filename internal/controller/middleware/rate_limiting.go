package middleware

// global rate limiting policy
type GlobalRateLimitingPolicy struct {
	RefillRate int64
	Capacity   int64
	Fallback   bool // if true, the rate limiting policy is a fallback policy
	Bypassed   bool // if true, the rate limiting policy is bypassed
}

// user rate limiting policy
type UserRateLimitingPolicy struct {
	TenantID   string
	UserID     string
	RefillRate int64
	Capacity   int64
	Fallback   bool // if true, the rate limiting policy is a fallback policy
	Bypassed   bool // if true, the rate limiting policy is bypassed
}
