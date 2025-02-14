package cmd

import (
	"fmt"
	"os"

	agenda "github.com/cleberbonifacio/go-cli-calendar/cmd/agenda"
	events "github.com/cleberbonifacio/go-cli-calendar/cmd/events"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "calendar",
		Short: "Your calendar CLI",
	}

	rootCmd.AddCommand(events.EventsCmd)
	rootCmd.AddCommand(agenda.AgendaCmd)
	return rootCmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
