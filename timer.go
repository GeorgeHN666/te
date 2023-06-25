package main

import (
	"fmt"
	"time"
)

var Timers = make(map[string]*time.Timer)

func StartTimer(id string, unix int64) {

	timer := time.NewTimer(time.Until(time.Unix(unix, 0)))

	select {
	case <-timer.C:
		fmt.Printf("Starting timer with id %s at %v\n\n", id, time.Now())

		// Send the email
		ReportScheduleAppointment()
	case <-time.After(time.Millisecond):
		Timers[id] = timer
		fmt.Printf("Timer with id %s got ready at %v\n\n", id, time.Now())
	}

}

func CancelTimer(id string) string {
	timer, exist := Timers[id]

	if exist {
		if !timer.Stop() {
			<-timer.C
		}
		delete(Timers, id)
		return fmt.Sprintf("Timer with id %s has terminated\n\n", id)
	} else {
		return fmt.Sprintf("Timer with id %s doesn't exist\n\n", id)
	}
}
