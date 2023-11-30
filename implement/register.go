package implement

import (
	"errors"
	"github.com/akley-MK4/pep-coroutine/define"
	"sync/atomic"
	"time"
)

type coroutineGroupInfo struct {
	baseStatsHandler baseGroupStatsHandler
}

var (
	regCoGroupMark uint32
	coGroupInfoMap = make(map[define.CoGroup]*coroutineGroupInfo)
)

func getCoroutineGroupInfoMap() map[define.CoGroup]*coroutineGroupInfo {
	return coGroupInfoMap
}

func RegisterCoroutineGroups(groups ...define.CoGroup) error {
	if !atomic.CompareAndSwapUint32(&regCoGroupMark, 0, 1) {
		return errors.New("repeated registration")
	}

	newGroupMap := make(map[define.CoGroup]*coroutineGroupInfo)
	for _, group := range groups {
		newGroupMap[group] = &coroutineGroupInfo{}
	}
	coGroupInfoMap = newGroupMap
	return nil
}

var (
	registerCoroutineTypesFunc = map[define.CoType]NewCoroutineFunc{
		define.TimerCoroutineType: func(coId define.CoId, coGroup define.CoGroup, interval time.Duration,
			handle define.CoroutineHandle, handleArgs ...interface{}) (ICoroutine, error) {
			
			co := &timerCoroutine{}
			co.baseCoroutine = newBaseCoroutine(coId, define.TimerCoroutineType, coGroup, interval, handle, handleArgs...)
			return co, nil
		},
	}
)
