# 实现常见的几种锁
## 自旋锁
### 基本原理
使用`CAS`技术实现`sync.Locker`接口。
### 调用接口
```go
// 创建锁
func NewSpinLock() sync.Locker
// 申请锁
func (sl *spinLock) Lock()
// 释放锁
func (sl *spinLock) Unlock() 
```

## 可重入锁
### 基本原理
在`sync.Mutex`基础上，增加计数器和拥有者信息。
### 调用接口
```go
// 创建锁
func NewRecursiveLock() sync.Locker
// 申请锁
func (sl *recursiveLock) Lock()
// 释放锁
func (sl *recursiveLock) Unlock() 
```
