package cmd

import (
	"github.com/spellshape/cli/spellshape/pkg/cliui"
	"github.com/spellshape/cli/spellshape/pkg/cosmosutil"
	"github.com/spf13/cobra"

	"github.com/spellshape/cli-plugin-network/network"
	"github.com/spellshape/cli-plugin-network/network/networktypes"
)

// NewNetworkRequestRemoveAccount creates a new command to send remove account request.
func NewNetworkRequestRemoveAccount() *cobra.Command {
	c := &cobra.Command{
		Use:   "remove-account [launch-id] [address]",
		Short: "Send request to remove a genesis account",
		RunE:  networkRequestRemoveAccountHandler,
		Args:  cobra.ExactArgs(2),
	}

	flagSetClearCache(c)
	c.Flags().AddFlagSet(flagNetworkFrom())
	c.Flags().AddFlagSet(flagSetHome())
	c.Flags().AddFlagSet(flagSetKeyringBackend())
	c.Flags().AddFlagSet(flagSetKeyringDir())
	return c
}

func networkRequestRemoveAccountHandler(cmd *cobra.Command, args []string) error {
	session := cliui.New(cliui.StartSpinner())
	defer session.End()

	nb, err := newNetworkBuilder(cmd, CollectEvents(session.EventBus()))
	if err != nil {
		return err
	}

	// parse launch ID
	launchID, err := network.ParseID(args[0])
	if err != nil {
		return err
	}

	// get the address for the account and change the prefix for Spellshape Chain
	address, err := cosmosutil.ChangeAddressPrefix(args[1], networktypes.SPN)
	if err != nil {
		return err
	}

	n, err := nb.Network()
	if err != nil {
		return err
	}

	return n.SendAccountRemoveRequest(
		cmd.Context(),
		launchID,
		address,
	)
}
