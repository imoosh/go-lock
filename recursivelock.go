package go_lock

import (
    "fmt"
    "github.com/petermattis/goid"
    "sync"
    "sync/atomic"
)

type recursiveLock struct {
    sync.Mutex
    owner     int64
    recursion int32
}

func NewRecursiveLock() sync.Locker {
    return new(recursiveLock)
}

func (rl *recursiveLock) Lock() {
    // 若当前goroutine已获得锁，立即返回。实现锁的可重入性
    gid := goid.Get()
    if atomic.LoadInt64(&rl.owner) == gid {
        rl.recursion++
        return
    }

    // 第一次获得锁，标记当前goroutine获得锁
    rl.Mutex.Lock()
    atomic.StoreInt64(&rl.owner, gid)
    rl.recursion = 1
}

func (rl *recursiveLock) Unlock() {
    // 若锁属于其他goroutine，错误使用
    gid := goid.Get()
    if atomic.LoadInt64(&rl.owner) != gid {
        panic(fmt.Sprintf("wrong the owner(%d): %d", rl.owner, gid))
    }

    // 同一个goroutine，释放锁次数必须与获得锁次数相同
    rl.recursion--
    if rl.recursion != 0 {
        return
    }

    // 释放锁
    atomic.StoreInt64(&rl.owner, -1)
    rl.Mutex.Unlock()
}
