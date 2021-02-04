package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/fernandezvara/scheduler"
)

func main() {

	fn := func() {
		fmt.Println(time.Now(), " running")
	}

	job, err := scheduler.Every(1).Minutes().Run(fn)
	if err != nil {
		panic(err)
	}

	time.Sleep(15 * time.Second)

	job.RunNow(fn)

	runtime.Goexit()

}
