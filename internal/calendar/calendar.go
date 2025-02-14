package calendar

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	gCalendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const AGENDA = "Roles"

var (
	ErrAgendaNotFound = errors.New("agenda not found")
	ErrAddAgenda      = errors.New("error add agenda")
	ErrEventsWeek     = errors.New("error list week")
)

type Calendar struct {
	Service    *gCalendar.Service
	CalendarID string
}

func NewClient() *Calendar {
	ctx := context.Background()
	credentials, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatal("unable to read credentials json")
	}
	service, err := gCalendar.NewService(ctx, option.WithCredentialsJSON(credentials))
	if err != nil {
		log.Fatalf("error to create google calendar service: %s", err.Error())
	}
	return &Calendar{Service: service}
}

func (c *Calendar) GetAgendaID() error {
	list, err := c.Service.CalendarList.List().Do()
	if err != nil {
		return ErrAgendaNotFound
	}
	for _, v := range list.Items {
		if v.Summary == AGENDA {
			c.CalendarID = v.Id
		}
	}

	return nil
}

func (c *Calendar) InsertAgenda(id string) error {
	entry := &gCalendar.CalendarListEntry{
		Id: c.CalendarID,
	}
	_, err := c.Service.CalendarList.Insert(entry).Do()
	if err != nil {
		return ErrAddAgenda
	}

	return nil
}

func (c *Calendar) ListWeekEvents() error {
	now := time.Now()
	weekday := now.Weekday()
	startDate := now.AddDate(0, 0, -int(weekday))
	endDate := startDate.AddDate(0, 0, 7)
	events, err := c.Service.Events.List(c.CalendarID).TimeMin(startDate.Format(time.RFC3339)).TimeMax(endDate.Format(time.RFC3339)).Do()
	if err != nil {
		return ErrEventsWeek
	}

	for _, v := range events.Items {
		fmt.Printf("%s | %s at %s\n", v.Summary, v.Start.DateTime, v.Start.DateTime)
	}

	return nil
}

func (c *Calendar) ListTodayEvents() error {
	year, month, day := time.Now().Date()
	startDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 0, 1)
	events, err := c.Service.Events.List(c.CalendarID).TimeMin(startDate.Format(time.RFC3339)).TimeMax(endDate.Format(time.RFC3339)).Do()
	if err != nil {
		return ErrEventsWeek
	}

	for _, v := range events.Items {
		fmt.Printf("%s | %s at %s\n", v.Summary, v.Start.DateTime, v.Start.DateTime)
	}

	return nil
}
