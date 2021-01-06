package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// MultiHARDHooks combine multiple HARD hooks, all hook functions are run in array sequence
type MultiHARDHooks []HARDHooks

// NewMultiHARDHooks returns a new MultiHARDHooks
func NewMultiHARDHooks(hooks ...HARDHooks) MultiHARDHooks {
	return hooks
}

// BeforeDepositCreated runs before a deposit is created
func (h MultiHARDHooks) BeforeDepositCreated(ctx sdk.Context, deposit Deposit) {
	for i := range h {
		h[i].BeforeDepositCreated(ctx, deposit)
	}
}

// BeforeDepositModified runs before a deposit is modified
func (h MultiHARDHooks) BeforeDepositModified(ctx sdk.Context, deposit Deposit) {
	for i := range h {
		h[i].BeforeDepositModified(ctx, deposit)
	}
}

// BeforeBorrowCreated runs before a borrow is created
func (h MultiHARDHooks) BeforeBorrowCreated(ctx sdk.Context, borrow Borrow) {
	for i := range h {
		h[i].BeforeBorrowCreated(ctx, borrow)
	}
}

// BeforeBorrowModified runs before a borrow is modified
func (h MultiHARDHooks) BeforeBorrowModified(ctx sdk.Context, borrow Borrow) {
	for i := range h {
		h[i].BeforeBorrowModified(ctx, borrow)
	}
}