package go_lock

import (
    "sync"
    "testing"
)

/*
   goos: darwin
   goarch: amd64
   pkg: lock-go
   cpu: Intel(R) Core(TM) i5-8500 CPU @ 3.00GHz
   BenchmarkMutexLock
   BenchmarkMutexLock-6       	28080441	        37.69 ns/op
   BenchmarkRecursiveLock
   BenchmarkRecursiveLock-6   	20344105	        59.34 ns/op
*/
func BenchmarkMutexLock(b *testing.B) {
    m := sync.Mutex{}
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            m.Lock()
            //nolint:staticcheck
            m.Unlock()
        }
    })
}

func BenchmarkRecursiveLock(b *testing.B) {
    m := NewRecursiveLock()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            m.Lock()
            //nolint:staticcheck
            m.Unlock()
        }
    })
}

func TestRecursiveLock(t *testing.T) {
    var (
        rl = NewRecursiveLock()
        wg = sync.WaitGroup{}
    )
    wg.Add(2)

    count, times := 0, 10000000
    for i := 0; i < 2; i++ {
        go func() {
            defer wg.Done()
            for i := 0; i < times; i++ {
                rl.Lock()
                count++
                //nolint:staticcheck
                rl.Unlock()
            }
        }()
    }
    wg.Wait()
    if times*2 != count {
        t.Errorf("lock times error: %d", count)
    }
    t.Logf("lock times : %d", count)
}
