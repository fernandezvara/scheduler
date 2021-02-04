# scheduler

[![Build status](https://img.shields.io/github/workflow/status/fernandezvara/scheduler/ci?style=for-the-badge)](https://github.com/fernandezvara/scheduler/actions?workflow=ci)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](https://pkg.go.dev/github.com/fernandezvara/scheduler)
[![Coveralls github](https://img.shields.io/coveralls/github/fernandezvara/scheduler?style=for-the-badge)](https://coveralls.io/github/fernandezvara/scheduler)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE)

Job scheduling made easy.

Scheduler allows you to schedule recurrent jobs with an easy-to-read syntax.

Inspired by the article **[Rethinking Cron](http://adam.heroku.com/past/2010/4/13/rethinking_cron/)** and the **[schedule](https://github.com/dbader/schedule)** python module.

## How to use?

```go
package main

import (
    "fmt"
    "runtime"
    "time"

    "github.com/fernandezvara/scheduler"
)

func main() {
    
    job := func() {
        t := time.Now()
        fmt.Println("Time's up! @", t.UTC())
    }

    // Run every 2 seconds but not now.
    scheduler.Every(2).Seconds().NotImmediately().Run(job)
      
    // Run now and every X.
    scheduler.Every(5).Minutes().Run(job)
    scheduler.Every().Day().Run(job)
    
    // Run at a point on time
    scheduler.Every().Monday().At("08:30").Run(job)
      
    // Keep the program from not exiting.
    runtime.Goexit()
    
}
```

## How it works?

By specifying the chain of calls, a `Job` struct is instantiated and a goroutine is starts observing the `Job`.

The goroutine will be on pause until:

* The next run scheduled is due. This will cause to execute the job.
* The `SkipWait` channel is activated. This will cause to execute the job.
* The `Quit` channel is activated. This will cause to finish the job.

## Not immediate recurrent jobs

By default, the behaviour of the recurrent jobs (Every(N) seconds, minutes, hours) is to start executing the job right away and then wait the required amount of time. By calling specifically `.NotImmediately()` you can override that behaviour and not execute it directly when the function `Run()` is called.

```go
scheduler.Every(5).Minutes().NotImmediately().Run(job)
```

## On demand execution

Any scheduled job can be triggered on demand between executions if required by using `RunNow()`. If the job requires args use `RunNowWithArgs(args []string)`.

```go
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

    // trigger manually alfter 15 seconds
    time.Sleep(15 * time.Second)
    job.RunNow(fn)

    runtime.Goexit()

}
```

```go
> go run main.go
2021-02-04 17:00:00.863400481 +0100 CET m=+0.000078489  running
2021-02-04 17:00:15.86348832 +0100 CET m=+15.000166348  running
2021-02-04 17:01:00.863541725 +0100 CET m=+60.000219723  running
2021-02-04 17:02:00.863630401 +0100 CET m=+120.000308419  running
```

## License

Distributed under MIT license. See `LICENSE` for more information.
