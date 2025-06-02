package main

func main() {
	var nilCh chan struct{} = nil
	nilCh <- struct{}{}
}

// STDOUT
//
// fatal error: all goroutines are asleep - deadlock!
//
// goroutine 1 [chan send (nil chan)]:
// main.main()
//         /Users/uberate/workspace/learn/codes/go/channel/writenil/write_nil_channel.go:5 +0x24
// exit status 2