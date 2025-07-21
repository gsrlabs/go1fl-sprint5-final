package spentenergy

import (
	"errors"
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.

	minWeight = 2.0   // Минимальный допустимый вес (кг)
	maxWeight = 635.0 // Максимальный допустимый вес (кг)
	minHeight = 0.50  // Минимальный допустимый рост (м)
	maxHeight = 2.75  // Максимальный допустимый рост (м)
)

// WalkingSpentCalories рассчитывает количество потраченных калорий при ходьбе
// Использует RunningSpentCalories с понижающим коэффициентом
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	caloriesBurned, err := RunningSpentCalories(steps, weight, height, duration)

	if err != nil {
		return 0, err
	}

	walkingCalories := caloriesBurned * walkingCaloriesCoefficient

	return walkingCalories, nil
}

// RunningSpentCalories рассчитывает количество потраченных калорий при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("неверное значение шагов")
	}

	if !CheckWeight(weight) {
		return 0, errors.New("неверное значение веса")
	}

	if !CheckHeight(height) {
		return 0, errors.New("неверное значение роста")
	}

	if duration <= 0 {
		return 0, errors.New("неверное значение продолжительности:")
	}

	speed := MeanSpeed(steps, height, duration)
	caloriesBurned := (weight * speed * duration.Minutes()) / minInH

	return caloriesBurned, nil
}

// MeanSpeed рассчитывает среднюю скорость в км/ч
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		fmt.Println("неверное значение продолжительности:", duration)
		return 0
	}

	distance := Distance(steps, height)
	speed := distance / duration.Hours()

	return speed
}

// Distance рассчитывает пройденную дистанцию в километрах
func Distance(steps int, height float64) float64 {
	if !CheckHeight(height) {
		fmt.Println("неверное значение роста:", height)
		return 0
	}

	if steps <= 0 {
		fmt.Println("неверное значение шагов:", steps)
		return 0
	}

	stepLength := height * stepLengthCoefficient
	distance := (float64(steps) * stepLength) / mInKm
	return distance
}

// CheckWeight проверяет корректность веса
func CheckWeight(weight float64) bool {
	if weight < minWeight {
		return false
	}
	if weight > maxWeight {
		return false
	}
	return true
}

// CheckHeight проверяет корректность роста
func CheckHeight(height float64) bool {
	if height < minHeight {
		return false
	}
	if height > maxHeight {
		return false
	}
	return true
}