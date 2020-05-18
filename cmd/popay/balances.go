package main

import (
	"fmt"
	"os"

	"github.com/jhampac/pueblo/database"
	"github.com/spf13/cobra"
)

func balancesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balances",
		Short: "Interact with balances (list...)",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {},
	}

	cmd.AddCommand(balancesListCmd)

	return cmd
}

var balancesListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all balances",
	Run: func(cmd *cobra.Command, args []string) {
		state, err := database.NewStateFromDisk()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer state.Close()

		fmt.Printf("Accounts balances:")
		fmt.Println("------------------------")
		fmt.Println("")

		for account, balance := range state.Balances {
			fmt.Println(fmt.Sprintf("%s: %d", account, balance))
		}
	},
}