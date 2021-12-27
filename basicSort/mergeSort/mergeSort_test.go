package mergeSort

import (
	"fmt"
	"interviewPractice/basicSort/sortData"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	if !reflect.DeepEqual(mergeSort(sortData.Data), sortData.Expect) {
		fmt.Printf("%+v\n", sortData.Data)
		t.Errorf("this two are not the same")
	}
}
