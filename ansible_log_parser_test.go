package ansiblelogparser

import (
	"reflect"
	"testing"

	"github.com/yukihiko-shinoda/go-ansible-log-parser/_testlibraries"
)

func TestPickupNumberPlayRecap(t *testing.T) {
	message, err := _testlibraries.LoadMessage()
	if err != nil {
		t.Errorf("%v", err)
	}
	actual, err := PickupNumberPlayRecap(*message)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !reflect.DeepEqual(*actual, StructPlayRecap{
		Ok:          3,
		Changed:     1,
		Unreachable: 0,
		Failed:      0,
		Skipped:     0,
		Rescued:     0,
		Ignored:     0,
	}) {
		t.Errorf("%v", *actual)
	}
}

func TestPickupNumberPlayRecapNotMatch(t *testing.T) {
	message := ""
	actual, err := PickupNumberPlayRecap(message)
	if actual != nil {
		t.Errorf("%v", actual)
	}
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestParseInt(t *testing.T) {
	matches := []string{"", "3", "1", "0", "0", "0", "0", "0"}
	actual, err := parseInt(matches)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !reflect.DeepEqual(*actual, StructPlayRecap{3, 1, 0, 0, 0, 0, 0}) {
		t.Errorf("%v", *actual)
	}
}

func TestParseIntError(t *testing.T) {
	matches := []string{"", "3", "1", "0", "0", "0", "0", "a"}
	_, err := parseInt(matches)
	if err == nil {
		t.Errorf("Return value: err is not nil!")
	}
}

func TestPickUpChangedTasks(t *testing.T) {
	expected := "TASK [file] *********************************************************************************************************************************************************************************************************************"
	message, err := _testlibraries.LoadMessage()
	if err != nil {
		t.Errorf("%v", err)
	}
	latestTaskName := ""
	actual, newLatestTaskName := PickUpChangedTasks(*message, latestTaskName)
	if !reflect.DeepEqual(actual, []string{expected}) {
		t.Errorf("%v", actual)
	}
	if newLatestTaskName != expected {
		t.Errorf("%v", newLatestTaskName)
	}
}
func TestTrimRecap(t *testing.T) {
	message, err := _testlibraries.LoadMessage()
	expected := "PLAY RECAP **********************************************************************************************************************************************************************************************************************\n" +
		"localhost                  : ok=3    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   \n\n"
	if err != nil {
		t.Errorf("%v", err)
	}
	var tests = []struct {
		testname string
		message  string
		expected string
	}{
		{"normal", *message, expected},
		{"empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			actual := TrimRecap(tt.message)
			if actual != tt.expected {
				t.Errorf("%v", actual)
			}
		})
	}
}
