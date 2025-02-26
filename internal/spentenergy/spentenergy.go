package spentenergy

import "time"

const (
	lenStep   = 0.65
	mInKm     = 1000
	minInH    = 60
	kmhInMsec = 0.278
	cmInM     = 100
	speed     = 1.39
)

const (
	walkingCaloriesWeightMultiplier = 0.035
	walkingSpeedHeightMultiplier    = 0.029
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	if steps < 0 || weight < 0 || height <= 0 || duration.Hours() <= 0 {
		return 0
	}

	meanSpeed := MeanSpeed(steps, duration)

	return ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed*meanSpeed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH
}

const (
	runningCaloriesMeanSpeedMultiplier = 18.0
	runningCaloriesMeanSpeedShift      = 20.0
)

func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	if steps <= 0 || weight <= 0 || duration.Hours() <= 0 {
		return 0
	}

	meanSpeed := MeanSpeed(steps, duration)
	return ((runningCaloriesMeanSpeedMultiplier * meanSpeed) - runningCaloriesMeanSpeedShift) * weight
}

func MeanSpeed(steps int, duration time.Duration) float64 {
	if steps <= 0 || duration.Hours() <= 0 {
		return 0
	}

	distance := Distance(steps)

	return distance / duration.Hours()
}

func Distance(steps int) float64 {
	if steps <= 0 {
		return 0
	}
	return float64(steps) * lenStep / float64(mInKm)
}
