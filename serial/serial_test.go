package serial

import (
	"fmt"
	"testing"
	"time"
)

func TestSerial(t *testing.T) {
	ch := make(chan string)
	go ReadArduino(ch)
	time.Sleep(2 * time.Second)

	//	for i := 0; i < 100; i++ {
	//		if i < 0 {
	//			continue
	//		}
	//

	for {
		data := <-ch
		fmt.Printf("%v\n", data)
	}

}
