package airvisual

const errPrefix = "airvisual: "

var (
	StateNotSupportedError = err(errPrefix + "state not supported")
	CityNotSupportedError  = err(errPrefix + "city not supported")
	RateLimitError         = err(errPrefix + "rate limit exceeded")
)
