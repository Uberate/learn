package main

func main() {
	var nilCh chan struct{} = nil
	<- nilCh
}

// STDOUT
//
// fatal error: all goroutines are asleep - deadlock!
//
// goroutine 1 [chan receive (nil chan)]:
// main.main()
//         /learn/codes/go/channel/readnil/read_nil_channel.go:5 +0x24
// exit status 2