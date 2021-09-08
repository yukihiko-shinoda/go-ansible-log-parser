package ansiblelogparser

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type StructPlayRecap struct {
	Ok          int
	Changed     int
	Unreachable int
	Failed      int
	Skipped     int
	Rescued     int
	Ignored     int
}

func PickupNumberPlayRecap(message string) (*StructPlayRecap, error) {
	messagePlayRecap := TrimRecap(message)
	re := regexp.MustCompile(`:\s+ok=(\d+)\s+changed=(\d+)\s+unreachable=(\d+)\s+failed=(\d+)\s+skipped=(\d+)\s+rescued=(\d+)\s+ignored=(\d+)`)
	matches := re.FindStringSubmatch(messagePlayRecap)
	if len(matches) == 0 {
		return nil, nil
	}
	return parseInt(matches)
}

func parseInt(matches []string) (*StructPlayRecap, error) {
	playRecap := StructPlayRecap{0, 0, 0, 0, 0, 0, 0}
	elem := reflect.ValueOf(&playRecap).Elem()
	cnt := elem.NumField()
	for index := 0; index < cnt; index++ {
		match := matches[index+1]
		integer, err := strconv.ParseInt(match, 10, 64)
		if err != nil {
			return &playRecap, err
		}
		elem.FieldByIndex([]int{index}).SetInt(integer)
	}
	return &playRecap, nil
}

func PickUpChangedTasks(message string, latestTaskName string) ([]string, string) {
	sliceLine := strings.Split(message, "\n")
	var sliceChangedTask []string
	for _, line := range sliceLine {
		if strings.Contains(line, "TASK [") || strings.Contains(line, "RUNNING HANDLER [") {
			latestTaskName = line
		}
		if strings.Contains(line, "changed:") {
			sliceChangedTask = append(sliceChangedTask, latestTaskName)
		}
	}
	return sliceChangedTask, latestTaskName
}

func TrimRecap(message string) string {
	sliceLine := strings.Split(message, "\n")
	last := len(sliceLine) - 1
	for index := range sliceLine {
		reverseIndex := last - index
		if strings.Contains(sliceLine[reverseIndex], "PLAY RECAP") {
			return strings.Join(sliceLine[reverseIndex:], "\n") + "\n"
		}
	}
	return ""
}
