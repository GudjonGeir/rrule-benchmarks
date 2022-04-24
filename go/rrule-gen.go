package main

import (
	"fmt"
	"time"

	"github.com/teambition/rrule-go"
)

func GenerateEvents() []time.Time {
	// 3 days a week
	r1, _ := rrule.NewRRule(rrule.ROption{
		Freq:      rrule.WEEKLY,
		Dtstart:   time.Date(2021, 2, 1, 10, 30, 0, 0, time.UTC),
		Interval:  1,
		Byweekday: []rrule.Weekday{rrule.MO, rrule.WE, rrule.FR},
	})

	// Daily
	r2, _ := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Dtstart: time.Date(2021, 2, 1, 10, 30, 0, 0, time.UTC),
	})

	// With exceptions
	set := rrule.Set{}
	r3, _ := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.WEEKLY,
		Dtstart: time.Date(2021, 2, 1, 17, 30, 0, 0, time.UTC),
	})
	set.RRule(r3)
	set.ExDate(time.Date(2021, 8, 16, 17, 30, 0, 0, time.UTC))
	set.ExDate(time.Date(2021, 9, 27, 17, 30, 0, 0, time.UTC))
	set.RDate(time.Date(2021, 8, 5, 10, 30, 0, 0, time.UTC))
	set.RDate(time.Date(2021, 8, 16, 10, 30, 0, 0, time.UTC))
	set.RDate(time.Date(2021, 9, 27, 10, 30, 0, 0, time.UTC))

	res := r1.Between(
		time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC), true)
	res = append(res, r2.Between(
		time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC), true)...)
	res = append(res, set.Between(
		time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC), true)...)

	return res
}

func main() {
	fmt.Println(len(GenerateEvents()))
}
