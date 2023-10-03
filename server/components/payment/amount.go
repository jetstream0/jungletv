package payment

import (
	"math/big"

	"github.com/palantir/stacktrace"
	"github.com/shopspring/decimal"
)

type Amount struct {
	*big.Int
}

func NewAmount(numbersToAdd ...*big.Int) Amount {
	n := big.NewInt(0)
	for _, toAdd := range numbersToAdd {
		n.Add(n, toAdd)
	}
	return Amount{n}
}

func NewAmountFromDecimal(d decimal.Decimal) Amount {
	return Amount{d.BigInt()}
}

func NewAmountFromAPIString(s string) (Amount, error) {
	bigint, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return Amount{}, stacktrace.NewError("invalid Amount value %s", s)
	}
	return Amount{bigint}, nil
}

func (a Amount) Decimal() decimal.Decimal {
	return decimal.NewFromBigInt(a.Int, 0)
}

func (a Amount) SerializeForAPI() string {
	if a.Int == nil {
		return "0"
	}
	return a.String()
}
