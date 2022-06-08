package main

import "testing"

// func BenchmarkGenerateEvents(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		GenerateEvents()
// 	}
// }

// func BenchmarkMaxProjectedEvents(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		ProjectMaxExpectedEvents()
// 	}
// }

// -------------------------------------------------------------
// func TestProjectDailies(b *testing.T) {
// 	res := ProjectDailies()
// 	print("%d", len(res))
// }

func BenchmarkProjectDailies(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectDailies()
	}
}

// -------------------------------------------------------------
// func TestProjectMultipleRulesDaily(b *testing.T) {
// 	res := ProjectMultipleRulesDaily()
// 	print("%d", len(res))
// }
func BenchmarkProjectMultipleRulesDaily(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectMultipleRulesDaily()
	}
}

// -------------------------------------------------------------
// func TestProjectMultipleRulesDailyOld(b *testing.T) {
// 	res := ProjectMultipleRulesDailyOld()
// 	print("%d", len(res))
// }
func BenchmarkProjectMultipleRulesDailyOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectMultipleRulesDailyOld()
	}
}

// -------------------------------------------------------------

// func TestProjectWeekly(b *testing.T) {
// 	res := ProjectWeekly()
// 	print("%d", len(res))
// }

func BenchmarkProjectWeekly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectWeekly()
	}
}

// -------------------------------------------------------------

// func TestProjectMultipleRulesWeekly(b *testing.T) {
// 	res := ProjectMultipleRulesWeekly()
// 	print("%d", len(res))
// }
func BenchmarkProjectMultipleRulesWeekly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectMultipleRulesWeekly()
	}
}

// -------------------------------------------------------------
func TestProjectMultipleRulesWeeklyOld(b *testing.T) {
	res := ProjectMultipleRulesWeeklyOld()
	print("%d", len(res))
}
func BenchmarkProjectMultipleRulesWeeklyOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProjectMultipleRulesWeeklyOld()
	}
}

// -------------------------------------------------------------
