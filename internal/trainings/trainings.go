package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	pd "github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)
	
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	pd.Personal
}

// Parse разбирает строку с данными о тренировке
// Формат данных: "steps,activityType,duration" (например: "5000,Бег,1h30m")
func (t *Training) Parse(datastring string) (err error) {
	splitData := strings.Split(datastring, ",")

	if len(splitData) != 3 {
		return errors.New("неправильное количество параметров")
	}

	t.Steps, err = strconv.Atoi(splitData[0])
	if err != nil {
		return  err
	}

	if t.Steps <= 0 {
		return  errors.New("неверное значение шагов")
	}

	t.Duration, err = time.ParseDuration(splitData[2])
	if err != nil {
		return err
	}

	if t.Duration <= 0 {
		return errors.New("неверная продолжительность - ноль")
	}

	t.TrainingType = splitData[1]

	return nil
	
}

// ActionInfo формирует отчет о тренировке
func (t Training) ActionInfo() (string, error) {

	distance := spentenergy.Distance(t.Steps, t.Height)

	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	caloriesBurned, err := calculateCalories(t.TrainingType, t.Steps, t.Weight, t.Height, t.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %w", err)
	}

	return formatTrainingInfo(t.TrainingType, t.Duration, distance, speed, caloriesBurned), nil
}

// calculateCalories рассчитывает количество потраченных калорий в зависимости от типа активности.
func calculateCalories(activityType string, steps int, weight, height float64, duration time.Duration) (float64, error) {
	switch activityType {
	case "Бег":
		return spentenergy.RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		return spentenergy.WalkingSpentCalories(steps, weight, height, duration)
	default:
		return 0, errors.New("неизвестный тип тренировки")
	}
}

// formatTrainingInfo форматирует информацию о тренировке в читаемую строку.
func formatTrainingInfo(activityType string, duration time.Duration, distance, speed, calories float64) string {
	return fmt.Sprintf(
		"Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n",
		activityType, duration.Hours(), distance, speed, calories,
	)
}
