package calc_distance

import (
	"strings"
	"strconv"
)

func calcDistance(directions []string) int {
	total := 0
	for _, s := range directions {
		total += parseDistance(s)
	}
	return total
}

func parseDistance(str string) int {
	distance := 0
	for _, elem := range strings.Fields(str) {
		if strings.HasSuffix(elem, "km") {
			kilometers, err := strconv.ParseFloat(elem[:len(elem) - 2], 64)
			if err == nil {
				distance += int(1000 * kilometers)
			}
		} else if strings.HasSuffix(elem, "m") {
			meters, err := strconv.Atoi(elem[:len(elem) - 1])
			if err == nil {
				distance += meters
			}
		}
	}
	return distance
}