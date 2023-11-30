package define

type CoId uint64
type CoType uint8
type CoGroup uint16
type CoStatus uint8

type CoroutineHandle func(coID CoId, args ...interface{}) bool
