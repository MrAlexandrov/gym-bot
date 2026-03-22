package main

import "time"

func getWeekNumber(date time.Time) int {
	_, week := date.ISOWeek()
	return week
}

func getYear(date time.Time) int {
	return date.Year()
}

func getWeekAndYear(date time.Time) (int, int) {
	return getWeekNumber(date), getYear(date)
}

func getCurrentWeekAndYear() (int, int) {
	return getWeekAndYear(time.Now())
}
