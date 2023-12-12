package pep_coroutine

import (
	"errors"
	"github.com/akley-MK4/pep-coroutine/define"
	"github.com/akley-MK4/pep-coroutine/implement"
	"github.com/akley-MK4/pep-coroutine/logger"
	"sync/atomic"
)

var (
	initMarkTag uint32
)

func InitializeLibrary(loggerInst logger.ILogger, groups ...define.CoGroup) (retErr error) {
	if !atomic.CompareAndSwapUint32(&initMarkTag, 0, 1) {
		return errors.New("repeated initialization")
	}

	defer func() {
		if retErr != nil {
			initMarkTag = 0
		}
	}()

	if loggerInst != nil {
		if err := logger.SetLoggerInstance(loggerInst); err != nil {
			retErr = errors.New("unable to set logger instance")
			return
		}
	}

	for _, g := range groups {
		if err := implement.AddCoroutineGroupInfo(g); err != nil {
			retErr = err
			return
		}
	}

	return
}
