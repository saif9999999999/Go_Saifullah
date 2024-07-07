package main

func defineGrade(score float64) string {
	if score < 0 || score > 100 {
		return "Invalid score"
	} else if score >= 85 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 55 {
		return "C"
	} else if score >= 40 {
		return "D"
	} else {
		return "E"
	}
}
