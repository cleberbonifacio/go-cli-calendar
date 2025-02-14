package cmd

import (
	"log"

	"github.com/cleberbonifacio/go-cli-calendar/internal/calendar"
	"github.com/spf13/cobra"
)

func init() {
	EventsCmd.AddCommand(EventsWeekCmd)
	EventsCmd.AddCommand(EventsTodayCmd)
}

var EventsCmd = &cobra.Command{
	Use:   "events",
	Short: "list all you calendar events",
	Long:  "check all your events",
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		err := c.GetAgendaID()
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}
