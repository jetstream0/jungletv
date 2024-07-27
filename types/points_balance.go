package types

import (
	"errors"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/palantir/stacktrace"
	"github.com/tnyim/jungletv/utils/transaction"
)

// PointsBalance is the points balance of an address
type PointsBalance struct {
	RewardsAddress string `dbKey:"true"`
	Balance        int
}

// GetPointsBalanceForAddress returns the points balance of the given address
func GetPointsBalanceForAddress(ctx transaction.WrappingContext, address string) (*PointsBalance, error) {
	s := sdb.Select().
		Where(sq.Eq{"points_balance.rewards_address": address})
	items, err := GetWithSelect[*PointsBalance](ctx, s)
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}

	if len(items) == 0 {
		return &PointsBalance{
			RewardsAddress: address,
		}, nil
	}
	return items[0], nil
}

// ErrInsufficientPointsBalance is returned when there is an insufficient points balance for the requested operation
var ErrInsufficientPointsBalance = errors.New("insufficient points balance")

// AdjustPointsBalanceOfAddress adjusts the points balance of the specified address by the specified amount
func AdjustPointsBalanceOfAddress(ctx transaction.WrappingContext, address string, amount int) error {
	ctx, err := transaction.Begin(ctx)
	if err != nil {
		return stacktrace.Propagate(err, "")
	}
	defer ctx.Rollback()

	// a CHECK (balance >= 0) exists in the table to prevent overdraw, even in concurrent transactions
	// the CHECK runs on the INSERT even if there is a conflict and fails the whole statement,
	// hence we first do an insert with zero if the balance row doesn't exist yet,
	// then we do an update to adjust the balance
	_, err = sdb.Insert("points_balance").
		Columns("rewards_address", "balance").
		Values(address, 0).
		Suffix("ON CONFLICT DO NOTHING").RunWith(ctx).ExecContext(ctx)
	if err != nil {
		return stacktrace.Propagate(err, "")
	}

	_, err = sdb.Update("points_balance").
		Where(sq.Eq{"points_balance.rewards_address": address}).
		Set("balance", sq.Expr("balance + ?", amount)).
		RunWith(ctx).ExecContext(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "points_balance_balance_check") {
			return stacktrace.Propagate(ErrInsufficientPointsBalance, "")
		}
		return stacktrace.Propagate(err, "")
	}

	return stacktrace.Propagate(ctx.Commit(), "")
}
