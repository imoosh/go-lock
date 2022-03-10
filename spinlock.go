package lock_go

import (
    "runtime"
    "sync"
    "sync/atomic"
)

type spinLock uint32

const maxBackoff = 16

func NewSpinLock() sync.Locker {
    return new(spinLock)
}

func (sl *spinLock) Lock() {
    backoff := 1
    for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
        for i := 0; i < maxBackoff; i++ {
            runtime.Gosched() // 让出时间片
        }
        if backoff < maxBackoff {
            backoff = backoff << 1
        }
    }
}

func (sl *spinLock) Unlock() {
    atomic.StoreUint32((*uint32)(sl), 0)
}
