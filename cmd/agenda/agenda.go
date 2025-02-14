package agenda

import (
	"fmt"
	"log"

	"github.com/cleberbonifacio/go-cli-calendar/internal/calendar"
	"github.com/spf13/cobra"
)

var AgendaCmd = &cobra.Command{
	Use:   "agenda",
	Short: "list all you agenda",
	Long:  "check all your agenda",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewClient()
		err := c.InsertAgenda(args[0])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("sucesso")
	},
}
