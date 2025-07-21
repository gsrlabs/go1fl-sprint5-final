package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	pd "github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps int
	Duration time.Duration
	pd.Personal
}

// Parse разбирает строку данных на количество шагов и продолжительность активности.
// Формат данных: "steps,duration" (например: "678,0h50m")
func (ds *DaySteps) Parse(datastring string) (err error) {
	splitData := strings.Split(datastring, ",")

	if len(splitData) != 2 {
		return errors.New("неправильное количество параметров")
	}

	ds.Steps, err = strconv.Atoi(splitData[0])
	if err != nil {
		return  err
	}

	if ds.Steps <= 0 {
		return  errors.New("неверное значение шагов")
	}

	ds.Duration, err = time.ParseDuration(splitData[1])
	if err != nil {
		return err
	}

	if ds.Duration <= 0 {
		return errors.New("неверная продолжительность - ноль")
	}

	return nil
}
// ActionInfo() формирует и возвращает строку с данными о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	dayActionInfo := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	)

	return dayActionInfo, nil
}
