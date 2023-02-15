package testutil

import (
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	campaigntypes "github.com/spellshape/network/x/campaign/types"
	launchtypes "github.com/spellshape/network/x/launch/types"
	profiletypes "github.com/spellshape/network/x/profile/types"
	rewardtypes "github.com/spellshape/network/x/reward/types"
)

//go:generate mockery --name CampaignClient --case underscore --output ../mocks
type CampaignClient interface {
	campaigntypes.QueryClient
}

//go:generate mockery --name ProfileClient --case underscore --output ../mocks
type ProfileClient interface {
	profiletypes.QueryClient
}

//go:generate mockery --name LaunchClient --case underscore --output ../mocks
type LaunchClient interface {
	launchtypes.QueryClient
}

//go:generate mockery --name RewardClient --case underscore --output ../mocks
type RewardClient interface {
	rewardtypes.QueryClient
}

//go:generate mockery --name BankClient --case underscore --output ../mocks
type BankClient interface {
	banktypes.QueryClient
}

//go:generate mockery --name StakingClient --case underscore --output ../mocks
type StakingClient interface {
	stakingtypes.QueryClient
}

//go:generate mockery --name AccountInfo --case underscore --output ../mocks
type AccountInfo interface {
	keyring.Record
}
