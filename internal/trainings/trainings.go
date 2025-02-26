package trainings

import (
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
		return err
	}

	t.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return err
	}

	t.TrainingType = data[1]
	if (data[1]) != "Бег" && data[1] != "Ходьба" {
		return err
	}

	t.Duration, err = time.ParseDuration(data[2])
	if err != nil {
		return err
	}

	if t.Steps <= 0 || t.Duration.Minutes() <= 0 {
		return err
	}

	return nil
}

func (t Training) ActionInfo() string {
	if t.TrainingType == "Бег" {
		text := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps), spentenergy.MeanSpeed(t.Steps, t.Duration), spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration))
		return text
	} else if t.TrainingType == "Ходьба" {
		text := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps), spentenergy.MeanSpeed(t.Steps, t.Duration), spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration))
		return text
	}

	return "неизвестный тип тренировки"
}
