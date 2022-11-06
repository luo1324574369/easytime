package easytime

import (
	"reflect"
	"testing"
	"time"
)

func TestMondayCurrentWeek(t *testing.T) {
	tests := []struct {
		name  string
		today time.Time
		want  time.Time
	}{
		{
			"",
			time.Date(2022, 11, 06, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 10, 31, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"",
			time.Date(2022, 11, 07, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 11, 07, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"",
			time.Date(2022, 11, 9, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 11, 07, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"",
			time.Date(2022, 11, 13, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 11, 07, 0, 0, 0, 0, time.Now().Location()),
		},
	}

	for _, tt := range tests {
		DefaultTime = tt.today
		t.Run(tt.name, func(t *testing.T) {
			if got := MondayCurrentWeek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MondayCurrentWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSundayCurrentWeek(t *testing.T) {
	tests := []struct {
		name  string
		today time.Time
		want  time.Time
	}{
		{
			"",
			time.Date(2022, 11, 06, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 11, 06, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"",
			time.Date(2022, 11, 07, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 11, 13, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"",
			time.Date(2022, 11, 9, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 11, 13, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"",
			time.Date(2022, 11, 13, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 11, 13, 0, 0, 0, 0, time.Now().Location()),
		},
	}

	for _, tt := range tests {
		DefaultTime = tt.today
		t.Run(tt.name, func(t *testing.T) {
			if got := SundayCurrentWeek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SundayCurrentWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstDayCurrentMonth(t *testing.T) {
	tests := []struct {
		name  string
		today time.Time
		want  time.Time
	}{
		{
			"",
			time.Date(2022, 11, 9, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 11, 01, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"",
			time.Date(2022, 11, 30, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 11, 01, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"",
			time.Date(2022, 11, 01, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 11, 01, 0, 0, 0, 0, time.Now().Location()),
		},
	}

	for _, tt := range tests {
		DefaultTime = tt.today
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstDayCurrentMonth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstDayCurrentMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastDayCurrentMonth(t *testing.T) {
	tests := []struct {
		name  string
		today time.Time
		want  time.Time
	}{
		{
			"",
			time.Date(2022, 11, 9, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 11, 30, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"",
			time.Date(2022, 11, 30, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 11, 30, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"",
			time.Date(2022, 11, 01, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 11, 30, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"",
			time.Date(2022, 12, 31, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 12, 31, 0, 0, 0, 0, time.Now().Location()),
		},
	}

	for _, tt := range tests {
		DefaultTime = tt.today
		t.Run(tt.name, func(t *testing.T) {
			if got := LastDayCurrentMonth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastDayCurrentMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstDayCurrentYear(t *testing.T) {
	tests := []struct {
		name  string
		today time.Time
		want  time.Time
	}{
		{
			"",
			time.Date(2022, 1, 1, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 1, 1, 0, 0, 0, 0, time.Now().Location()),
		},
		{
			"",
			time.Date(2022, 12, 31, 0, 0, 0, 0, time.Now().Location()),
			time.Date(2022, 1, 1, 0, 0, 0, 0, time.Now().Location()),
		},
	}

	for _, tt := range tests {
		DefaultTime = tt.today
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstDayCurrentYear(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstDayCurrentYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastDayCurrentYear(t *testing.T) {
	tests := []struct {
		name  string
		today time.Time
		want  time.Time
	}{
		{
			"",
			time.Date(2022, 1, 1, 1, 1, 1, 1, time.Now().Location()),
			time.Date(2022, 12, 31, 1, 1, 1, 1, time.Now().Location()),
		},
		{
			"",
			time.Date(2022, 12, 31, 1, 1, 1, 1, time.Now().Location()),
			time.Date(2022, 12, 31, 1, 1, 1, 1, time.Now().Location()),
		},
	}

	for _, tt := range tests {
		DefaultTime = tt.today
		t.Run(tt.name, func(t *testing.T) {
			if got := LastDayCurrentYear(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastDayCurrentYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
