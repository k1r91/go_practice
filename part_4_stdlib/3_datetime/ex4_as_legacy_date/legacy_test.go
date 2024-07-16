package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

// начало решения

// asLegacyDate преобразует время в легаси-дату
func asLegacyDate(t time.Time) string {
	unixStartTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	dur := t.Sub(unixStartTime)
	nseconds := dur.Nanoseconds()
	seconds := nseconds / 1000000000
	nanoseconds := nseconds % 1000000000
	var nanosecondsStr string
	if nanoseconds == 0 {
		nanosecondsStr = "0"
	} else {
		nanosecondsStr = fmt.Sprintf("%09d", nanoseconds)
		nanosecondsStr = strings.TrimRight(nanosecondsStr, "0")
	}
	return fmt.Sprintf("%d.%s", seconds, nanosecondsStr)
}

// parseLegacyDate преобразует легаси-дату во время.
// Возвращает ошибку, если легаси-дата некорректная.
func parseLegacyDate(d string) (time.Time, error) {
	parts := strings.Split(d, ".")
	if len(parts) != 2 {
		return time.Time{}, fmt.Errorf("incorrect legacy date format")
	}
	seconds, err := strconv.Atoi(parts[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("incorrect seconds: got %v", seconds)
	}
	nanoseconds, err := strconv.Atoi(parts[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("incorrect nanoseconds: got %v", nanoseconds)
	}
	secondsDuration := time.Duration(seconds * int(time.Second))
	for i := 0; i < 9 - len(parts[1]); i ++ {
		nanoseconds *= 10
	}
	nanosecondsDuration := time.Duration(nanoseconds * int(time.Nanosecond))
	unixStartTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	result := unixStartTime.Add(secondsDuration).Add(nanosecondsDuration)
	return result, nil
}

// конец решения

func Test_asLegacyDate(t *testing.T) {
	samples := map[time.Time]string{
		time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC): "3600.123456789",
		time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC):         "3600.0",
		time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC):         "0.0",
		time.Date(2022, 5, 24, 14, 45, 22, 951205999, time.UTC): "1653403522.951205999",
		time.Date(2022, 5, 24, 14, 45, 22, 951000000, time.UTC): "1653403522.951",
		time.Date(1970, 1, 1, 1, 0, 0, 1, time.UTC): "3600.000000001",
	}
	for src, want := range samples {
		got := asLegacyDate(src)
		if got != want {
			t.Fatalf("%v: got %v, want %v", src, got, want)
		}
	}
}

func Test_parseLegacyDate(t *testing.T) {
	samples := map[string]time.Time{
		"3600.123456789": time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC),
		"3600.0":         time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC),
		"0.0":            time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		"1.123456789":    time.Date(1970, 1, 1, 0, 0, 1, 123456789, time.UTC),
	}
	for src, want := range samples {
		got, err := parseLegacyDate(src)
		if err != nil {
			t.Fatalf("%v: unexpected error", src)
		}
		if !got.Equal(want) {
			t.Fatalf("%v: got %v, want %v", src, got, want)
		}
	}
}
