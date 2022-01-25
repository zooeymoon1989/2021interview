package main

func main() {
	var n = 3
	done := make(chan interface{})
	defer close(done)

	repeat := func(done <-chan interface{}, s string) <-chan string {
		stream := make(chan string)
		go func() {
			defer close(stream)
			for {
				select {
				case <-done:
					return
				case stream <- s:
				}
			}
		}()
		return stream
	}

	take := func(done <-chan interface{}, n int, stream <-chan string) <-chan string {
		takeStream := make(chan string)
		go func() {
			defer close(takeStream)
			for i := 0; i < n; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-stream:
				}
			}
		}()
		return takeStream
	}

	jiagong := func(done <-chan interface{}, stream <-chan string, s string) <-chan string {
		jiaGongstream := make(chan string)
		go func() {
			defer close(jiaGongstream)
			for {
				select {
				case <-done:
					return
				case jiaGongstream <- <-stream + s:
				}
			}
		}()
		return jiaGongstream
	}

	for a := range take(done, n, jiagong(done, jiagong(done, repeat(done, "a"), "b"), "c")) {
		print(a)
	}
}
