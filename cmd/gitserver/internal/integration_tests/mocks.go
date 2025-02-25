// Code generated by go-mockgen 1.3.7; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package inttests

import (
	"sync"

	internal "github.com/sourcegraph/sourcegraph/cmd/gitserver/internal"
	common "github.com/sourcegraph/sourcegraph/cmd/gitserver/internal/common"
)

// MockRepositoryLock is a mock implementation of the RepositoryLock
// interface (from the package
// github.com/sourcegraph/sourcegraph/cmd/gitserver/internal) used for unit
// testing.
type MockRepositoryLock struct {
	// ReleaseFunc is an instance of a mock function object controlling the
	// behavior of the method Release.
	ReleaseFunc *RepositoryLockReleaseFunc
	// SetStatusFunc is an instance of a mock function object controlling
	// the behavior of the method SetStatus.
	SetStatusFunc *RepositoryLockSetStatusFunc
}

// NewMockRepositoryLock creates a new mock of the RepositoryLock interface.
// All methods return zero values for all results, unless overwritten.
func NewMockRepositoryLock() *MockRepositoryLock {
	return &MockRepositoryLock{
		ReleaseFunc: &RepositoryLockReleaseFunc{
			defaultHook: func() {
				return
			},
		},
		SetStatusFunc: &RepositoryLockSetStatusFunc{
			defaultHook: func(string) {
				return
			},
		},
	}
}

// NewStrictMockRepositoryLock creates a new mock of the RepositoryLock
// interface. All methods panic on invocation, unless overwritten.
func NewStrictMockRepositoryLock() *MockRepositoryLock {
	return &MockRepositoryLock{
		ReleaseFunc: &RepositoryLockReleaseFunc{
			defaultHook: func() {
				panic("unexpected invocation of MockRepositoryLock.Release")
			},
		},
		SetStatusFunc: &RepositoryLockSetStatusFunc{
			defaultHook: func(string) {
				panic("unexpected invocation of MockRepositoryLock.SetStatus")
			},
		},
	}
}

// NewMockRepositoryLockFrom creates a new mock of the MockRepositoryLock
// interface. All methods delegate to the given implementation, unless
// overwritten.
func NewMockRepositoryLockFrom(i internal.RepositoryLock) *MockRepositoryLock {
	return &MockRepositoryLock{
		ReleaseFunc: &RepositoryLockReleaseFunc{
			defaultHook: i.Release,
		},
		SetStatusFunc: &RepositoryLockSetStatusFunc{
			defaultHook: i.SetStatus,
		},
	}
}

// RepositoryLockReleaseFunc describes the behavior when the Release method
// of the parent MockRepositoryLock instance is invoked.
type RepositoryLockReleaseFunc struct {
	defaultHook func()
	hooks       []func()
	history     []RepositoryLockReleaseFuncCall
	mutex       sync.Mutex
}

// Release delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockRepositoryLock) Release() {
	m.ReleaseFunc.nextHook()()
	m.ReleaseFunc.appendCall(RepositoryLockReleaseFuncCall{})
	return
}

// SetDefaultHook sets function that is called when the Release method of
// the parent MockRepositoryLock instance is invoked and the hook queue is
// empty.
func (f *RepositoryLockReleaseFunc) SetDefaultHook(hook func()) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Release method of the parent MockRepositoryLock instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *RepositoryLockReleaseFunc) PushHook(hook func()) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *RepositoryLockReleaseFunc) SetDefaultReturn() {
	f.SetDefaultHook(func() {
		return
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *RepositoryLockReleaseFunc) PushReturn() {
	f.PushHook(func() {
		return
	})
}

func (f *RepositoryLockReleaseFunc) nextHook() func() {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *RepositoryLockReleaseFunc) appendCall(r0 RepositoryLockReleaseFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of RepositoryLockReleaseFuncCall objects
// describing the invocations of this function.
func (f *RepositoryLockReleaseFunc) History() []RepositoryLockReleaseFuncCall {
	f.mutex.Lock()
	history := make([]RepositoryLockReleaseFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// RepositoryLockReleaseFuncCall is an object that describes an invocation
// of method Release on an instance of MockRepositoryLock.
type RepositoryLockReleaseFuncCall struct{}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c RepositoryLockReleaseFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c RepositoryLockReleaseFuncCall) Results() []interface{} {
	return []interface{}{}
}

// RepositoryLockSetStatusFunc describes the behavior when the SetStatus
// method of the parent MockRepositoryLock instance is invoked.
type RepositoryLockSetStatusFunc struct {
	defaultHook func(string)
	hooks       []func(string)
	history     []RepositoryLockSetStatusFuncCall
	mutex       sync.Mutex
}

// SetStatus delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockRepositoryLock) SetStatus(v0 string) {
	m.SetStatusFunc.nextHook()(v0)
	m.SetStatusFunc.appendCall(RepositoryLockSetStatusFuncCall{v0})
	return
}

// SetDefaultHook sets function that is called when the SetStatus method of
// the parent MockRepositoryLock instance is invoked and the hook queue is
// empty.
func (f *RepositoryLockSetStatusFunc) SetDefaultHook(hook func(string)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// SetStatus method of the parent MockRepositoryLock instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *RepositoryLockSetStatusFunc) PushHook(hook func(string)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *RepositoryLockSetStatusFunc) SetDefaultReturn() {
	f.SetDefaultHook(func(string) {
		return
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *RepositoryLockSetStatusFunc) PushReturn() {
	f.PushHook(func(string) {
		return
	})
}

func (f *RepositoryLockSetStatusFunc) nextHook() func(string) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *RepositoryLockSetStatusFunc) appendCall(r0 RepositoryLockSetStatusFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of RepositoryLockSetStatusFuncCall objects
// describing the invocations of this function.
func (f *RepositoryLockSetStatusFunc) History() []RepositoryLockSetStatusFuncCall {
	f.mutex.Lock()
	history := make([]RepositoryLockSetStatusFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// RepositoryLockSetStatusFuncCall is an object that describes an invocation
// of method SetStatus on an instance of MockRepositoryLock.
type RepositoryLockSetStatusFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 string
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c RepositoryLockSetStatusFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c RepositoryLockSetStatusFuncCall) Results() []interface{} {
	return []interface{}{}
}

// MockRepositoryLocker is a mock implementation of the RepositoryLocker
// interface (from the package
// github.com/sourcegraph/sourcegraph/cmd/gitserver/internal) used for unit
// testing.
type MockRepositoryLocker struct {
	// AllStatusesFunc is an instance of a mock function object controlling
	// the behavior of the method AllStatuses.
	AllStatusesFunc *RepositoryLockerAllStatusesFunc
	// StatusFunc is an instance of a mock function object controlling the
	// behavior of the method Status.
	StatusFunc *RepositoryLockerStatusFunc
	// TryAcquireFunc is an instance of a mock function object controlling
	// the behavior of the method TryAcquire.
	TryAcquireFunc *RepositoryLockerTryAcquireFunc
}

// NewMockRepositoryLocker creates a new mock of the RepositoryLocker
// interface. All methods return zero values for all results, unless
// overwritten.
func NewMockRepositoryLocker() *MockRepositoryLocker {
	return &MockRepositoryLocker{
		AllStatusesFunc: &RepositoryLockerAllStatusesFunc{
			defaultHook: func() (r0 map[common.GitDir]string) {
				return
			},
		},
		StatusFunc: &RepositoryLockerStatusFunc{
			defaultHook: func(common.GitDir) (r0 string, r1 bool) {
				return
			},
		},
		TryAcquireFunc: &RepositoryLockerTryAcquireFunc{
			defaultHook: func(common.GitDir, string) (r0 internal.RepositoryLock, r1 bool) {
				return
			},
		},
	}
}

// NewStrictMockRepositoryLocker creates a new mock of the RepositoryLocker
// interface. All methods panic on invocation, unless overwritten.
func NewStrictMockRepositoryLocker() *MockRepositoryLocker {
	return &MockRepositoryLocker{
		AllStatusesFunc: &RepositoryLockerAllStatusesFunc{
			defaultHook: func() map[common.GitDir]string {
				panic("unexpected invocation of MockRepositoryLocker.AllStatuses")
			},
		},
		StatusFunc: &RepositoryLockerStatusFunc{
			defaultHook: func(common.GitDir) (string, bool) {
				panic("unexpected invocation of MockRepositoryLocker.Status")
			},
		},
		TryAcquireFunc: &RepositoryLockerTryAcquireFunc{
			defaultHook: func(common.GitDir, string) (internal.RepositoryLock, bool) {
				panic("unexpected invocation of MockRepositoryLocker.TryAcquire")
			},
		},
	}
}

// NewMockRepositoryLockerFrom creates a new mock of the
// MockRepositoryLocker interface. All methods delegate to the given
// implementation, unless overwritten.
func NewMockRepositoryLockerFrom(i internal.RepositoryLocker) *MockRepositoryLocker {
	return &MockRepositoryLocker{
		AllStatusesFunc: &RepositoryLockerAllStatusesFunc{
			defaultHook: i.AllStatuses,
		},
		StatusFunc: &RepositoryLockerStatusFunc{
			defaultHook: i.Status,
		},
		TryAcquireFunc: &RepositoryLockerTryAcquireFunc{
			defaultHook: i.TryAcquire,
		},
	}
}

// RepositoryLockerAllStatusesFunc describes the behavior when the
// AllStatuses method of the parent MockRepositoryLocker instance is
// invoked.
type RepositoryLockerAllStatusesFunc struct {
	defaultHook func() map[common.GitDir]string
	hooks       []func() map[common.GitDir]string
	history     []RepositoryLockerAllStatusesFuncCall
	mutex       sync.Mutex
}

// AllStatuses delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockRepositoryLocker) AllStatuses() map[common.GitDir]string {
	r0 := m.AllStatusesFunc.nextHook()()
	m.AllStatusesFunc.appendCall(RepositoryLockerAllStatusesFuncCall{r0})
	return r0
}

// SetDefaultHook sets function that is called when the AllStatuses method
// of the parent MockRepositoryLocker instance is invoked and the hook queue
// is empty.
func (f *RepositoryLockerAllStatusesFunc) SetDefaultHook(hook func() map[common.GitDir]string) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// AllStatuses method of the parent MockRepositoryLocker instance invokes
// the hook at the front of the queue and discards it. After the queue is
// empty, the default hook function is invoked for any future action.
func (f *RepositoryLockerAllStatusesFunc) PushHook(hook func() map[common.GitDir]string) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *RepositoryLockerAllStatusesFunc) SetDefaultReturn(r0 map[common.GitDir]string) {
	f.SetDefaultHook(func() map[common.GitDir]string {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *RepositoryLockerAllStatusesFunc) PushReturn(r0 map[common.GitDir]string) {
	f.PushHook(func() map[common.GitDir]string {
		return r0
	})
}

func (f *RepositoryLockerAllStatusesFunc) nextHook() func() map[common.GitDir]string {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *RepositoryLockerAllStatusesFunc) appendCall(r0 RepositoryLockerAllStatusesFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of RepositoryLockerAllStatusesFuncCall objects
// describing the invocations of this function.
func (f *RepositoryLockerAllStatusesFunc) History() []RepositoryLockerAllStatusesFuncCall {
	f.mutex.Lock()
	history := make([]RepositoryLockerAllStatusesFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// RepositoryLockerAllStatusesFuncCall is an object that describes an
// invocation of method AllStatuses on an instance of MockRepositoryLocker.
type RepositoryLockerAllStatusesFuncCall struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 map[common.GitDir]string
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c RepositoryLockerAllStatusesFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c RepositoryLockerAllStatusesFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// RepositoryLockerStatusFunc describes the behavior when the Status method
// of the parent MockRepositoryLocker instance is invoked.
type RepositoryLockerStatusFunc struct {
	defaultHook func(common.GitDir) (string, bool)
	hooks       []func(common.GitDir) (string, bool)
	history     []RepositoryLockerStatusFuncCall
	mutex       sync.Mutex
}

// Status delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockRepositoryLocker) Status(v0 common.GitDir) (string, bool) {
	r0, r1 := m.StatusFunc.nextHook()(v0)
	m.StatusFunc.appendCall(RepositoryLockerStatusFuncCall{v0, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Status method of the
// parent MockRepositoryLocker instance is invoked and the hook queue is
// empty.
func (f *RepositoryLockerStatusFunc) SetDefaultHook(hook func(common.GitDir) (string, bool)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Status method of the parent MockRepositoryLocker instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *RepositoryLockerStatusFunc) PushHook(hook func(common.GitDir) (string, bool)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *RepositoryLockerStatusFunc) SetDefaultReturn(r0 string, r1 bool) {
	f.SetDefaultHook(func(common.GitDir) (string, bool) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *RepositoryLockerStatusFunc) PushReturn(r0 string, r1 bool) {
	f.PushHook(func(common.GitDir) (string, bool) {
		return r0, r1
	})
}

func (f *RepositoryLockerStatusFunc) nextHook() func(common.GitDir) (string, bool) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *RepositoryLockerStatusFunc) appendCall(r0 RepositoryLockerStatusFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of RepositoryLockerStatusFuncCall objects
// describing the invocations of this function.
func (f *RepositoryLockerStatusFunc) History() []RepositoryLockerStatusFuncCall {
	f.mutex.Lock()
	history := make([]RepositoryLockerStatusFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// RepositoryLockerStatusFuncCall is an object that describes an invocation
// of method Status on an instance of MockRepositoryLocker.
type RepositoryLockerStatusFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 common.GitDir
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 string
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 bool
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c RepositoryLockerStatusFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c RepositoryLockerStatusFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// RepositoryLockerTryAcquireFunc describes the behavior when the TryAcquire
// method of the parent MockRepositoryLocker instance is invoked.
type RepositoryLockerTryAcquireFunc struct {
	defaultHook func(common.GitDir, string) (internal.RepositoryLock, bool)
	hooks       []func(common.GitDir, string) (internal.RepositoryLock, bool)
	history     []RepositoryLockerTryAcquireFuncCall
	mutex       sync.Mutex
}

// TryAcquire delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockRepositoryLocker) TryAcquire(v0 common.GitDir, v1 string) (internal.RepositoryLock, bool) {
	r0, r1 := m.TryAcquireFunc.nextHook()(v0, v1)
	m.TryAcquireFunc.appendCall(RepositoryLockerTryAcquireFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the TryAcquire method of
// the parent MockRepositoryLocker instance is invoked and the hook queue is
// empty.
func (f *RepositoryLockerTryAcquireFunc) SetDefaultHook(hook func(common.GitDir, string) (internal.RepositoryLock, bool)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// TryAcquire method of the parent MockRepositoryLocker instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *RepositoryLockerTryAcquireFunc) PushHook(hook func(common.GitDir, string) (internal.RepositoryLock, bool)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *RepositoryLockerTryAcquireFunc) SetDefaultReturn(r0 internal.RepositoryLock, r1 bool) {
	f.SetDefaultHook(func(common.GitDir, string) (internal.RepositoryLock, bool) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *RepositoryLockerTryAcquireFunc) PushReturn(r0 internal.RepositoryLock, r1 bool) {
	f.PushHook(func(common.GitDir, string) (internal.RepositoryLock, bool) {
		return r0, r1
	})
}

func (f *RepositoryLockerTryAcquireFunc) nextHook() func(common.GitDir, string) (internal.RepositoryLock, bool) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *RepositoryLockerTryAcquireFunc) appendCall(r0 RepositoryLockerTryAcquireFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of RepositoryLockerTryAcquireFuncCall objects
// describing the invocations of this function.
func (f *RepositoryLockerTryAcquireFunc) History() []RepositoryLockerTryAcquireFuncCall {
	f.mutex.Lock()
	history := make([]RepositoryLockerTryAcquireFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// RepositoryLockerTryAcquireFuncCall is an object that describes an
// invocation of method TryAcquire on an instance of MockRepositoryLocker.
type RepositoryLockerTryAcquireFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 common.GitDir
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 internal.RepositoryLock
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 bool
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c RepositoryLockerTryAcquireFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c RepositoryLockerTryAcquireFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}
