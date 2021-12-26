package shellSort

import (
	"fmt"
	"interviewPractice/basicSort/sortData"
	"reflect"
	"testing"
)

func TestShellSort(t *testing.T) {
	if !reflect.DeepEqual(shellSort(sortData.Data), sortData.Expect) {
		fmt.Printf("%+v\n", sortData.Data)
		t.Errorf("this two are not the same")
	}
}
