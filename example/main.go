package main

import (
	"fmt"
)

func main() {
	//defer func() {
	//	time.Sleep(time.Second * 1)
	//	outputStats()
	//}()

	//loggerInst := newExampleLogger("[example]")
	//if err := pep_coroutine.InitializeLibrary(nil, coGroup1, coGroup2, coGroup3); err != nil {
	//	logger.GetLoggerInstance().WarningF("Failed to initialize library, %v", err)
	//	return
	//}

	if err := runExample1(); err != nil {
		fmt.Printf("Failed to run the example, %v\n", err)
	}
}
