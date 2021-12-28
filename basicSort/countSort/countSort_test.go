package countSort

import (
	"fmt"
	"interviewPractice/basicSort/sortData"
	"reflect"
	"testing"
)

func TestCountSort(t *testing.T) {
	if !reflect.DeepEqual(countSort(sortData.Data), sortData.Expect) {
		fmt.Printf("%+v\n", sortData.Data)
		t.Errorf("this two are not the same")
	}
}
