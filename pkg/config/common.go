package config

import (
	"github.com/ulule/limiter/v3"
	"time"
)

// Define a GlobalRatelimit limit of 5 requests per minute for the specific API.
var GlobalRatelimit = limiter.Rate{
	Period: 20 * time.Second,
	Limit:  20,
}

var GlobalJWTTimeToLive = time.Hour * 24
