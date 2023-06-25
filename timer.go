package main

import (
	"fmt"
	"time"
)

var Timers = make(map[string]*time.Timer)

func StartTimer(id string, unix int64) {

	fmt.Println(time.Now().Unix())

	timer := time.NewTimer(time.Until(time.Unix(unix, 0)))

	s := <-time.After(time.Millisecond)
	Timers[id] = timer
	fmt.Printf("Timer with id %s got launched at %v\n\n", id, s)

	f := <-timer.C
	fmt.Printf("Timer done with id %s at %v\n\n", id, f)

	// Send the email
	err := ReportScheduleAppointment()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CancelTimer(id string) string {
	timer, exist := Timers[id]

	if exist {
		if !timer.Stop() {
			<-timer.C
		}
		delete(Timers, id)
		fmt.Printf("Timer with id %s has terminated\n\n", id)
		return fmt.Sprintf("Timer with id %s has terminated\n\n", id)
	} else {
		fmt.Printf("Timer with id %s doesn't exist\n\n", id)
		return fmt.Sprintf("Timer with id %s doesn't exist\n\n", id)
	}
}
