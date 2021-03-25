package scenario

import (
	"fmt"
	"simplebattle/racer"
)

func GoSleepRunner() {
	fmt.Println("Start Racer and wait some millisecond")
	racer.RacerSleepForResult()
	fmt.Println("===================================")
}

func GoChannelRunner() {
	fmt.Println("Start Racer and wait until someone win")
	racer.RacerUseChannel()
	fmt.Println("===================================")
}
