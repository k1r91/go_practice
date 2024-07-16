package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

// начало решения

// Task описывает задачу, выполненную в определенный день
type Task struct {
	Date  time.Time
	Dur   time.Duration
	Title string
}

// ParsePage разбирает страницу журнала
// и возвращает задачи, выполненные за день
func ParsePage(src string) ([]Task, error) {
	lines := strings.Split(src, "\n")
	date, err := parseDate(lines[0])
	if err != nil {
		return []Task{}, err
	}
	tasks, err := parseTasks(date, lines[1:])
	if err != nil {
		return []Task{}, err
	}
	sortTasks(tasks)
	return tasks, nil
}

// parseDate разбирает дату в формате дд.мм.гггг
func parseDate(src string) (time.Time, error) {
	t, err := time.Parse("02.01.2006", src)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// parseTasks разбирает задачи из записей журнала
func parseTasks(date time.Time, lines []string) ([]Task, error) {
	tasksMap := make(map[string]time.Duration)
	var tasks []Task
	for i, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return []Task{}, errors.New("incorrect duration")
		}
		start, err := time.Parse("15:04 ", parts[0])
		if err != nil {
			return []Task{}, fmt.Errorf("incorrect start time at line %d", i)
		}
		endParts := strings.Fields(parts[1])
		if len(endParts) < 2 {
			return []Task{}, errors.New("incorrect duration")
		}
		end, err := time.Parse("15:04", endParts[0])
		taskName := strings.Join(endParts[1:], " ")
		if err != nil {
			return []Task{}, fmt.Errorf("incorrect end time at line %d", i)
		}
		if start.After(end) || start.Equal(end) {
			return []Task{}, fmt.Errorf("start time must be lower than end time")
		}
		tasksMap[taskName] += end.Sub(start)
	}
	for taskName, duration := range tasksMap {
		tasks = append(tasks, Task{date, duration, taskName})
	}
	return tasks, nil
}

// sortTasks упорядочивает задачи по убыванию длительности
func sortTasks(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Dur > tasks[j].Dur
	})
}

// конец решения
// ::footer

func main() {
	page := `15.04.2022
8:00 - 8:30 Завтрак
8:30 - 9:30 Оглаживание кота
9:30 - 10:00 Интернеты
10:00 - 14:00 Напряженная работа
14:00 - 14:45 Обед
14:45 - 15:00 Оглаживание кота
15:00 - 19:00 Напряженная работа
19:00 - 19:30 Интернеты
19:30 - 22:30 Безудержное веселье
22:30 - 23:00 Оглаживание кота`

	entries, err := ParsePage(page)
	if err != nil {
		panic(err)
	}
	fmt.Println("Мои достижения за", entries[0].Date.Format("2006-01-02"))
	for _, entry := range entries {
		fmt.Printf("- %v: %v\n", entry.Title, entry.Dur)
	}

	// ожидаемый результат
	/*
		Мои достижения за 2022-04-15
		- Напряженная работа: 8h0m0s
		- Безудержное веселье: 3h0m0s
		- Оглаживание кота: 1h45m0s
		- Интернеты: 1h0m0s
		- Обед: 45m0s
		- Завтрак: 30m0s
	*/
}