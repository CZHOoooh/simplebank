package utils

// Contains for all supported currency
const (
	USD = "USD"
	EUR = "EUR"
	RMB = "RMB"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, RMB:
		return true
	default:
		return false
	}
}
