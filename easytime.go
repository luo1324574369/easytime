package easytime

import "time"

var DefaultTime time.Time = time.Now()

const (
	Week = 7
)

func MondayCurrentWeek() time.Time {
	weekDay := DefaultTime.Weekday()
	if weekDay == time.Sunday {
		return DefaultTime.AddDate(0, 0, -6)
	}
	return DefaultTime.AddDate(0, 0, int(1-DefaultTime.Weekday()))
}

func SundayCurrentWeek() time.Time {
	weekDay := DefaultTime.Weekday()
	if weekDay == time.Sunday {
		return DefaultTime
	}
	return DefaultTime.AddDate(0, 0, int(7-weekDay))
}

func MonDayLastWeek() time.Time {
	DefaultTime = DefaultTime.AddDate(0, 0, -Week)
	return MondayCurrentWeek()
}

func SundayLastWeek() time.Time {
	DefaultTime = DefaultTime.AddDate(0, 0, -Week)
	return SundayCurrentWeek()
}

func MondayNextWeek() time.Time {
	DefaultTime = DefaultTime.AddDate(0, 0, Week)
	return MondayCurrentWeek()
}

func SundayNextWeek() time.Time {
	DefaultTime = DefaultTime.AddDate(0, 0, Week)
	return SundayCurrentWeek()
}

func FirstDayCurrentMonth() time.Time {
	return time.Date(DefaultTime.Year(), DefaultTime.Month(), 1, DefaultTime.Hour(), DefaultTime.Minute(), DefaultTime.Second(), DefaultTime.Nanosecond(), DefaultTime.Location())
}

func LastDayCurrentMonth() time.Time {
	return time.Date(DefaultTime.Year(), DefaultTime.Month()+1, 1, DefaultTime.Hour(), DefaultTime.Minute(), DefaultTime.Second(), DefaultTime.Nanosecond(), DefaultTime.Location()).AddDate(0, 0, -1)
}

func FirstDayCurrentYear() time.Time {
	return time.Date(DefaultTime.Year(), 1, 1, DefaultTime.Hour(), DefaultTime.Minute(), DefaultTime.Second(), DefaultTime.Nanosecond(), DefaultTime.Location())
}

func LastDayCurrentYear() time.Time {
	return time.Date(DefaultTime.Year()+1, 1, 1, DefaultTime.Hour(), DefaultTime.Minute(), DefaultTime.Second(), DefaultTime.Nanosecond(), DefaultTime.Location()).AddDate(0, 0, -1)
}
