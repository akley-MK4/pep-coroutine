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
	coGroup1 define.CoGroup = iota + 1
	coGroup2
	coGroup3
	coGroup4
)

func runExample1() (retErr error) {
	var handle define.CoroutineHandle = func(coID define.CoId, args ...interface{}) bool {
		//var tErr *int
		//*tErr = 1
		logger.GetLoggerInstance().DebugF("Coroutine %v is running, args: %v", coID, args)
		return true
	}

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

	for i := 0; i < 10; i++ {
		if err := implement.CreateAndStartStatelessCoroutine(coGroup2, handle, 4, 5, 6); err != nil {
			retErr = err
			return
		}
	}

	time.Sleep(time.Second * 5)

	return
}

func outputStats() {
	stats := implement.FetchStats()
	d, _ := json.Marshal(stats)
	fmt.Printf("Stats: \n%v\n", string(d))

}
