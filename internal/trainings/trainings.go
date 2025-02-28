package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mary-ppv/sprint5/internal/personaldata"
	"github.com/mary-ppv/sprint5/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return errors.New("length should be 3")
	}

	t.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return err
	}

	t.TrainingType = data[1]

	t.Duration, err = time.ParseDuration(data[2])
	if err != nil {
		return err
	}

	if t.Steps <= 0 || t.Duration.Minutes() <= 0 {
		return errors.New("length should be positive")
	}

	return nil
}

func (t Training) ActionInfo() string {
	distance := spentenergy.Distance(t.Steps)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)
	caloriesRunning := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
	caloriesWalking := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	
	switch t.TrainingType {
	case "Бег":
		text := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, caloriesRunning)
		return text
	case "Ходьба":
		text := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, caloriesWalking)
		return text
	default:
		return "неизвестный тип тренировки"
	}
}
