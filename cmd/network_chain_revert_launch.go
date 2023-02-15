package cmd

import (
	"github.com/spellshape/cli/spellshape/pkg/cliui"
	"github.com/spf13/cobra"

	"github.com/spellshape/cli-plugin-network/network"
	"github.com/spellshape/cli-plugin-network/network/networkchain"
)

// NewNetworkChainRevertLaunch creates a new chain revert launch command
// to revert a launched chain.
func NewNetworkChainRevertLaunch() *cobra.Command {
	c := &cobra.Command{
		Use:   "revert-launch [launch-id]",
		Short: "Revert launch of a network",
		Args:  cobra.ExactArgs(1),
		RunE:  networkChainRevertLaunchHandler,
	}

	c.Flags().AddFlagSet(flagNetworkFrom())
	c.Flags().AddFlagSet(flagSetKeyringBackend())
	c.Flags().AddFlagSet(flagSetKeyringDir())

	return c
}

func networkChainRevertLaunchHandler(cmd *cobra.Command, args []string) error {
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

	n, err := nb.Network()
	if err != nil {
		return err
	}

	chainLaunch, err := n.ChainLaunch(cmd.Context(), launchID)
	if err != nil {
		return err
	}

	c, err := nb.Chain(networkchain.SourceLaunch(chainLaunch))
	if err != nil {
		return err
	}

	return n.RevertLaunch(cmd.Context(), launchID, c)
}
