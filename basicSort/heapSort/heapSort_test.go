package heapSort

import (
	"fmt"
	"interviewPractice/basicSort/sortData"
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	if !reflect.DeepEqual(heapSort(sortData.Data), sortData.Expect) {
		fmt.Printf("%+v\n", sortData.Data)
		t.Errorf("this two are not the same")
	}
}
