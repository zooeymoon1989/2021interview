package quickSort

import (
	"fmt"
	"interviewPractice/basicSort/sortData"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	if !reflect.DeepEqual(quickSort(sortData.Data), sortData.Expect) {
		fmt.Printf("%+v\n", sortData.Data)
		t.Errorf("this two are not the same")
	}
	fmt.Printf("%+v\n", quickSort(sortData.Data))
}
