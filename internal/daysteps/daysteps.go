package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mary-ppv/sprint5/internal/personaldata"
	"github.com/mary-ppv/sprint5/internal/spentenergy"
)

const (
	StepLength = 0.65
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return err
	}

	ds.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return err
	}

	ds.Duration, err = time.ParseDuration(data[1])
	if err != nil {
		return err
	}

	if ds.Steps <= 0 || ds.Duration <= 0 {
		return err
	}

	return nil
}

func (ds DaySteps) ActionInfo() string {

	text := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, spentenergy.Distance(ds.Steps), spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration))

	return text
}
