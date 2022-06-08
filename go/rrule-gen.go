package main

import (
	"fmt"
	"math/rand"
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

func ProjectDailies() []time.Time {
	dtstart := time.Date(2021, 2, 1, 10, 30, 0, 0, time.UTC)
	sampleRule, _ := rrule.NewRRule(rrule.ROption{
		Freq:     rrule.DAILY,
		Dtstart:  dtstart,
		Interval: 1,
	})

	return sampleRule.Between(dtstart, dtstart.AddDate(0, 0, 10000), true)
}

func ProjectMultipleRulesDaily() []time.Time {
	dtstart := time.Date(2021, 2, 1, 10, 30, 0, 0, time.UTC)
	sampleRule, _ := rrule.NewRRule(rrule.ROption{
		Freq:     rrule.DAILY,
		Dtstart:  dtstart,
		Interval: 1,
	})

	iterations := 10000

	res := make([]time.Time, 0)

	for i := 0; i < iterations; i++ {
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

	iterations := 10000

	res := make([]time.Time, 0)

	for i := 0; i < iterations; i++ {
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

	return sampleRule.Between(dtstart, dtstart.AddDate(0, 0, 3333*7), true)
}
func ProjectMultipleRulesWeekly() []time.Time {
	dtstart := time.Date(2021, 2, 1, 10, 30, 0, 0, time.UTC)
	sampleRule, _ := rrule.NewRRule(rrule.ROption{
		Freq:      rrule.WEEKLY,
		Dtstart:   dtstart,
		Interval:  1,
		Byweekday: []rrule.Weekday{rrule.MO, rrule.WE, rrule.FR},
	})

	iterations := 10000 / 3

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

	iterations := 10000 / 3

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
