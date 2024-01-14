package finance

import "github.com/samber/lo"

// CentToDollar converts cent to dollar
func CentToDollar(amount int64) float64 {
	return float64(amount) / 100
}

// DollarToCent converts dollar to cent
func DollarToCent(amount float64) int64 {
	return int64(amount * 100)
}

// ConvertAmountCurrency converts the smallest base to dollar
// for dollar currencies
func ConvertAmountCurrency(amount int64, currencyCode string) float64 {
	if lo.Contains([]string{
		"USD", "SGD", "CAD", "EUR", "AUD", "HKD", "DZD",
		"BSD",
	}, currencyCode) {
		return CentToDollar(amount)
	}

	return float64(amount)
}
