package cpu

import (
	"runtime"
	"sync"
	"time"
)

// UserTime 用于在用户空间消费CPU资源
// userPect 设置总CPU所消费的大致百分比，如20，表示20%
// runDuration 设置持续的时间
func UserTime(userPect int, runDuration time.Duration) {
	numCPU := runtime.NumCPU()
	prefixNum := runtime.GOMAXPROCS(numCPU)

	var wg sync.WaitGroup
	for i := 0; i <= numCPU; i++ {
		wg.Add(1)
		go func(userPect int, runDuration time.Duration) {
			defer wg.Done()
			runCPUUserTime(userPect, runDuration)
		}(userPect, runDuration)
	}
	wg.Wait()

	if prefixNum >= 1 {
		runtime.GOMAXPROCS(prefixNum)
	}
}

func runCPUUserTime(userPect int, runDuration time.Duration) {
	idleDuration := 100 * time.Millisecond
	busyDuration := idleDuration

	if userPect != 0 {
		busyDuration = time.Duration(userPect) * time.Millisecond
	}

	runTime := time.Now()

	for {
		startTime := time.Now()

		for {
			taskTime := time.Now()

			if taskTime.Sub(startTime) >= busyDuration {
				break
			}
		}

		if startTime.Sub(runTime) >= runDuration {
			break
		}

		time.Sleep(idleDuration)
	}
}
