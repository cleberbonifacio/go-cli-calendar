package cmd

import (
	"log"

	"github.com/cleberbonifacio/go-cli-calendar/internal/calendar"
	"github.com/spf13/cobra"
)

var EventsTodayCmd = &cobra.Command{
	Use:   "today",
	Short: "list all events for today",
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		err := c.GetAgendaID()
		if err != nil {
			log.Fatal(err.Error())
		}
		c.ListTodayEvents()
	},
}
