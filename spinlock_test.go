package lock_go

import (
    "runtime"
    "sync"
    "sync/atomic"
    "testing"
)

/*
   goos: darwin
   goarch: amd64
   pkg: lock-go
   cpu: Intel(R) Core(TM) i5-8500 CPU @ 3.00GHz
   BenchmarkMutex
   BenchmarkMutex-6           	28435448	        38.55 ns/op
   BenchmarkBasicSpinLock
   BenchmarkBasicSpinLock-6   	53480954	        22.88 ns/op
   BenchmarkSpinLock
   BenchmarkSpinLock-6        	83418048	        13.76 ns/op
   PASS
*/
type basicSpinLock uint32

func (sl *basicSpinLock) Lock() {
    for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
        runtime.Gosched()
    }
}

func (sl *basicSpinLock) Unlock() {
    atomic.StoreUint32((*uint32)(sl), 0)
}

func NewBasicSpinLock() sync.Locker {
    return new(basicSpinLock)
}

func BenchmarkMutex(b *testing.B) {
    m := sync.Mutex{}
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            m.Lock()
            //nolint:staticcheck
            m.Unlock()
        }
    })
}

func BenchmarkBasicSpinLock(b *testing.B) {
    spin := NewBasicSpinLock()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            spin.Lock()
            //nolint:staticcheck
            spin.Unlock()
        }
    })
}

func BenchmarkSpinLock(b *testing.B) {
    spin := NewSpinLock()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            spin.Lock()
            //nolint:staticcheck
            spin.Unlock()
        }
    })
}
