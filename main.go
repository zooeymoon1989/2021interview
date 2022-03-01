package main

func main() {
	//var n = 3
	//done := make(chan interface{})
	//defer close(done)
	//
	//repeat := func(done <-chan interface{}, s string) <-chan string {
	//	stream := make(chan string)
	//	go func() {
	//		defer close(stream)
	//		for {
	//			select {
	//			case <-done:
	//				return
	//			case stream <- s:
	//			}
	//		}
	//	}()
	//	return stream
	//}
	//
	//take := func(done <-chan interface{}, n int, stream <-chan string) <-chan string {
	//	takeStream := make(chan string)
	//	go func() {
	//		defer close(takeStream)
	//		for i := 0; i < n; i++ {
	//			select {
	//			case <-done:
	//				return
	//			case takeStream <- <-stream:
	//			}
	//		}
	//	}()
	//	return takeStream
	//}
	//
	//process := func(done <-chan interface{}, stream <-chan string, s string) <-chan string {
	//	processStream := make(chan string)
	//	go func() {
	//		defer close(processStream)
	//		for {
	//			select {
	//			case <-done:
	//				return
	//			case processStream <- <-stream + s:
	//			}
	//		}
	//	}()
	//	return processStream
	//}
	//
	//for a := range take(done, n, process(done, process(done, repeat(done, "a"), "b"), "c")) {
	//	print(a)
	//}

	//const constSize = 10
	//var varSize = 10
	//var s1 []int32
	//s2 := make([]int32, varSize)
	//s3 := make([]int32, constSize)
	//s4 := make([]int32, varSize, varSize)
	//s5 := make([]int32, constSize, constSize)
	//s6 := make([]int32, constSize, varSize)
	//s7 := make([]int32, varSize, constSize)
	//_ = s1
	//_ = s2
	//_ = s3
	//_ = s4
	//_ = s5
	//_ = s6+-96*
	//_ = s7

}
