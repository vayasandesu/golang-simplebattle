package racer

import (
	"fmt"
	"math/rand"
	"time"
)

func racing(name string, from int, to int) {
	start := time.Now()
	for i := from; i <= to; i++ {
		wait := rand.Intn(5)
		time.Sleep(time.Duration(wait) * time.Millisecond)
	}
	fmt.Println(name, " run complete use time", time.Since(start))
}

func racingWithChannel(name string, from int, to int, reportResult chan<- string) {
	start := time.Now()

	for i := from; i <= to; i++ {
		wait := rand.Intn(5)
		time.Sleep(time.Duration(wait) * time.Millisecond)
	}

	reportResult <- fmt.Sprintln(name, " complete in ", time.Since(start))
}

func RacerSleepForResult() {
	waitTime := 7
	go racing("a", 0, 3)
	go racing("b", 0, 5)

	fmt.Printf("Going to sleep and wait %d ms\n", waitTime)
	time.Sleep(time.Duration(waitTime) * time.Millisecond)
	fmt.Println("I Wake up for end it, may someone not done yet")
}

func RacerUseChannel() {
	channel := make(chan string)

	go racingWithChannel("a", 0, 20, channel)
	go racingWithChannel("b", 0, 20, channel)

	fmt.Println("I will waiting for winner")
	result := <-channel
	fmt.Println("winner is : ", result)
}
