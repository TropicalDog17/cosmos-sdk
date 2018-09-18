package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// GenesisState - all distribution state that must be provided at genesis
type GenesisState struct {
	FeePool            FeePool             `json:"fee_pool"`
	ValidatorDistInfos []ValidatorDistInfo `json:"validator_dist_infos"`
	DelegatorDistInfos []DelegatorDistInfo `json:"delegator_dist_infos"`
	DelegatorDistInfos []sdk.AccAddress    `json:"delegator_dist_infos"`
}

func NewGenesisState(feePool FeePool, vdis []ValidatorDistInfo, ddis []DelegatorDistInfo) GenesisState {
	return GenesisState{
		FeePool:            feePool,
		ValidatorDistInfos: vdis,
		DelegatorDistInfos: ddis,
	}
}

// get raw genesis raw message for testing
func DefaultGenesisState() GenesisState {
	return GenesisState{
		FeePool: InitialFeePool(),
	}
}
