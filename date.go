package godes

import (
	"fmt"
	"math"
)

func GetDay() float64 {
	day := math.Max(1, math.Ceil((stime / (24 * 60))))
	return day
}

func GetHour() float64 {
	elapsed_hours := math.Floor(stime / 60)
	day := math.Floor(stime / (24*60))
	hours := math.Floor(elapsed_hours - (day * 24))
	return hours
}

func GetMinute() float64 {
	// stime is total number of elapsed minutes
	// need to subtract out number of elapsed HOURS (in minutes)
	elapsed_hours := math.Floor(stime / 60)
	minute := stime - elapsed_hours * 60
	
	return minute
}

type SimulationDate struct {
	day float64
	hour float64
	minute float64
}

type SimulationTime struct {
	Hour float64
	Minute float64
}

func GetSimulationTime() SimulationTime {
	return SimulationTime {
		Hour: GetHour(),
		Minute: GetMinute(),
	}
}

func GetSimulationDate() SimulationDate {
	date := SimulationDate{GetDay(), GetHour(), GetMinute()}
	return date
}

func (sd SimulationDate) IsAfter(test SimulationDate) bool {
	return sd.day > test.day || sd.hour > test.hour || sd.minute > test.minute
}

func (sd SimulationDate) ToString() string {
	return fmt.Sprintf("Day: %2.0f, Hour: %2.0f, Minute: %2.2f", sd.day, sd.hour, sd.minute)
}