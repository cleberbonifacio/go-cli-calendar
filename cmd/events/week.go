package cmd

import (
	"log"

	"github.com/cleberbonifacio/go-cli-calendar/internal/calendar"
	"github.com/spf13/cobra"
)

var EventsWeekCmd = &cobra.Command{
	Use:   "week",
	Short: "list all events week",
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		err := c.GetAgendaID()
		if err != nil {
			log.Fatal(err.Error())
		}
		c.ListWeekEvents()
	},
}
