package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/teambition/rrule-go"
)

type interval int

const (
	MINUTE interval = iota
	HOUR
	DAY
	WEEKLY
)

func GetUnixTimeValueForInterval(reqInterval interval) int64 {
	switch reqInterval {
	case MINUTE:
		return 60
	case HOUR:
		return 60 * GetUnixTimeValueForInterval(MINUTE)
	case DAY:
		return 24 * GetUnixTimeValueForInterval(HOUR)
	case WEEKLY:
		return 7 * GetUnixTimeValueForInterval(DAY)
	default:
		return 0
	}
}

func GenerateRandomDatesWithInterval(intervalType interval, amount int64) (time.Time, time.Time) {
	initialTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
	min := time.Unix(initialTime, 0)
	max := time.Unix(initialTime+GetUnixTimeValueForInterval(intervalType)*amount, 0)
	return min, max
}

func ProjectMaxExpectedEvents() []time.Time {

	sampleRule, _ := rrule.NewRRule(rrule.ROption{
		Freq:      rrule.WEEKLY,
		Dtstart:   time.Date(2021, 2, 1, 10, 30, 0, 0, time.UTC),
		Interval:  1,
		Byweekday: []rrule.Weekday{rrule.MO, rrule.WE, rrule.FR},
	})

	iterations := 1000000

	res := make([]time.Time, 0)

	for i := 0; i < iterations; i++ {
		min, max := GenerateRandomDatesWithInterval(DAY, 1)
		res = append(res, sampleRule.Between(min, max, true)...)
	}

	return res
}

const numberOfEvents = 100000

func ProjectDailies() []time.Time {
	dtstart := time.Date(2021, 2, 1, 10, 30, 0, 0, time.UTC)
	sampleRule, _ := rrule.NewRRule(rrule.ROption{
		Freq:     rrule.DAILY,
		Dtstart:  dtstart,
		Interval: 1,
	})

	return sampleRule.Between(dtstart, dtstart.AddDate(0, 0, numberOfEvents), true)
}

func ProjectMultipleRulesDaily() []time.Time {
	dtstart := time.Date(2021, 2, 1, 10, 30, 0, 0, time.UTC)
	sampleRule, _ := rrule.NewRRule(rrule.ROption{
		Freq:     rrule.DAILY,
		Dtstart:  dtstart,
		Interval: 1,
	})

	res := make([]time.Time, 0)

	for i := 0; i < numberOfEvents; i++ {
		res = append(res, sampleRule.Between(dtstart, dtstart, true)...)
	}

	return res
}

func ProjectMultipleRulesDailyOld() []time.Time {
	dtstart := time.Date(2021, 2, 1, 10, 30, 0, 0, time.UTC)
	sampleRule, _ := rrule.NewRRule(rrule.ROption{
		Freq:     rrule.DAILY,
		Dtstart:  dtstart.AddDate(-2, 0, 0),
		Interval: 1,
	})

	res := make([]time.Time, 0)

	for i := 0; i < numberOfEvents; i++ {
		res = append(res, sampleRule.Between(dtstart, dtstart, true)...)
	}

	return res
}

func ProjectWeekly() []time.Time {
	dtstart := time.Date(2021, 2, 1, 10, 30, 0, 0, time.UTC)
	sampleRule, _ := rrule.NewRRule(rrule.ROption{
		Freq:      rrule.WEEKLY,
		Dtstart:   dtstart,
		Interval:  1,
		Byweekday: []rrule.Weekday{rrule.MO, rrule.WE, rrule.FR},
	})

	return sampleRule.Between(dtstart, dtstart.AddDate(0, 0, numberOfEvents/3*7), true)
}
func ProjectMultipleRulesWeekly() []time.Time {
	dtstart := time.Date(2021, 2, 1, 10, 30, 0, 0, time.UTC)
	sampleRule, _ := rrule.NewRRule(rrule.ROption{
		Freq:      rrule.WEEKLY,
		Dtstart:   dtstart,
		Interval:  1,
		Byweekday: []rrule.Weekday{rrule.MO, rrule.WE, rrule.FR},
	})

	iterations := numberOfEvents / 3

	res := make([]time.Time, 0)

	for i := 0; i < iterations; i++ {
		res = append(res, sampleRule.Between(dtstart, dtstart.AddDate(0, 0, 6), true)...)
	}

	return res
}
func ProjectMultipleRulesWeeklyOld() []time.Time {
	dtstart := time.Date(2021, 2, 1, 10, 30, 0, 0, time.UTC)
	sampleRule, _ := rrule.NewRRule(rrule.ROption{
		Freq:      rrule.WEEKLY,
		Dtstart:   dtstart.AddDate(-2, 0, 0),
		Interval:  1,
		Byweekday: []rrule.Weekday{rrule.MO, rrule.WE, rrule.FR},
	})

	iterations := numberOfEvents / 3

	res := make([]time.Time, 0)

	for i := 0; i < iterations; i++ {
		res = append(res, sampleRule.Between(dtstart, dtstart.AddDate(0, 0, 6), true)...)
	}

	return res
}

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
		time.Date(2021, 8, 7, 23, 59, 59, 0, time.UTC), true)
	res = append(res, r2.Between(
		time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2021, 8, 7, 23, 59, 59, 0, time.UTC), true)...)
	res = append(res, set.Between(
		time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2021, 8, 7, 23, 59, 59, 0, time.UTC), true)...)

	return res
}

// func main() {
// 	fmt.Println(len(ProjectDailies()))
// 	fmt.Println(len(ProjectMultipleRulesDaily()))
// 	fmt.Println(len(ProjectMultipleRulesDailyOld()))
// 	fmt.Println(len(ProjectWeekly()))
// 	fmt.Println(len(ProjectMultipleRulesWeekly()))
// 	fmt.Println(len(ProjectMultipleRulesWeeklyOld()))
// }

func generateWeek(w http.ResponseWriter, r *http.Request) {
	dates := GenerateEvents()
	// fmt.Printf("%v", dates)
	json.NewEncoder(w).Encode(dates)
}

func handleRequests() {
	http.HandleFunc("/", generateWeek)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequests()
}
