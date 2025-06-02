package main

import (
	"fmt"
)

func main() {
	var nilCh chan struct{} = nil
	select {
	case <-nilCh:
		// never reach here
		fmt.Println("receive from nilCh")
	default:
		fmt.Println("default")
	}
}

// STDOUT:
//
// default
