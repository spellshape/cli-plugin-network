package network

import (
	"errors"
	"testing"
	"time"

	"github.com/spellshape/cli/spellshape/pkg/cosmosaccount"
	"github.com/spellshape/cli/spellshape/pkg/xtime"
	"github.com/stretchr/testify/require"

	"github.com/spellshape/cli-plugin-network/network/testutil"
)

var sampleTime = time.Unix(1000, 1000)

func newSuite(account cosmosaccount.Account) (testutil.Suite, Network) {
	suite := testutil.NewSuite()
	return suite, New(
		suite.CosmosClientMock,
		account,
		WithCampaignQueryClient(suite.CampaignQueryMock),
		WithLaunchQueryClient(suite.LaunchQueryMock),
		WithProfileQueryClient(suite.ProfileQueryMock),
		WithRewardQueryClient(suite.RewardClient),
		WithStakingQueryClient(suite.StakingClient),
		WithMonitoringConsumerQueryClient(suite.MonitoringConsumerClient),
		WithBankQueryClient(suite.BankClient),
		WithCustomClock(xtime.NewClockMock(sampleTime)),
	)
}

func TestParseID(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want uint64
		err  error
	}{
		{
			name: "valid number",
			id:   "10",
			want: 10,
		},
		{
			name: "invalid uint",
			id:   "-10",
			err:  errors.New("error parsing ID: strconv.ParseUint: parsing \"-10\": invalid syntax"),
		},
		{
			name: "invalid string",
			id:   "test",
			err:  errors.New("error parsing ID: strconv.ParseUint: parsing \"test\": invalid syntax"),
		},
		{
			name: "invalid launch id",
			id:   "0",
			err:  errors.New("ID must be greater than 0"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseID(tt.id)
			if tt.err != nil {
				require.Error(t, err)
				require.Equal(t, tt.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func SampleSharePercent(t *testing.T, denom string, nominator, denominator uint64) SharePercent {
	sp, err := NewSharePercent(denom, nominator, denominator)
	require.NoError(t, err)
	return sp
}
