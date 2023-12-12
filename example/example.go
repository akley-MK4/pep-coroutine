package main

import (
	"encoding/json"
	"fmt"
	"github.com/akley-MK4/pep-coroutine/define"
	"github.com/akley-MK4/pep-coroutine/implement"
	"github.com/akley-MK4/pep-coroutine/logger"
	"time"
)

const (
	coGroup1 = "g1"
	coGroup2 = "g2"
	coGroup3 = "g3"
	coGroup4 = "g4"
)

func runExample1() (retErr error) {
	var handle define.CoroutineHandle = func(coID define.CoId, args ...interface{}) bool {
		//var tErr *int
		//*tErr = 1
		logger.GetLoggerInstance().DebugF("Coroutine %v is running, args: %v", coID, args)
		return true
	}

	implement.AddCoroutineGroupInfo(coGroup2)
	for i := 0; i < 1; i++ {
		if err := implement.CreateAndStartStatelessCoroutine(coGroup2, func(coID define.CoId, args ...interface{}) bool {
			time.Sleep(time.Microsecond * 1)
			return false
		}, 4, 5, 6); err != nil {
			retErr = err
			return
		}
	}
	time.Sleep(5)
	outputStats()

	time.Sleep(5)

	co, createCoErr := implement.CreateCoroutine(define.CoType(0), coGroup1, time.Second*2, handle, 1, 2, 3)
	if createCoErr != nil {
		retErr = createCoErr
		return
	}

	if err := implement.StartCoroutine(co); err != nil {
		retErr = err
		return
	}

	time.Sleep(time.Second * 1)
	outputStats()

	time.Sleep(time.Second * 10)
	if err := implement.CloseCoroutine(co); err != nil {
		//retErr = err
		//return
	}

	time.Sleep(time.Second * 2)
	//time.Sleep(time.Second * 1)
	//return

	co3, createCo3Err := implement.CreateCoroutine(define.CoType(0), coGroup3, time.Second*2, handle, 1, 2, 3)
	if createCo3Err != nil {
		retErr = createCo3Err
		return
	}

	if err := implement.StartCoroutine(co3); err != nil {
		retErr = err
		return
	}

	time.Sleep(time.Second * 5)

	return
}

func outputStats() {
	stats := implement.FetchStats()
	d, _ := json.Marshal(stats)
	fmt.Printf("Stats: \n%v\n", string(d))

}
