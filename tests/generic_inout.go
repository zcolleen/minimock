// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package tests

//go:generate minimock -i github.com/gojuno/minimock/v3/tests.genericInout -o generic_inout.go -n GenericInoutMock -p tests

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// GenericInoutMock implements genericInout
type GenericInoutMock[T any] struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcName          func(t1 T) (t2 T)
	inspectFuncName   func(t1 T)
	afterNameCounter  uint64
	beforeNameCounter uint64
	NameMock          mGenericInoutMockName[T]
}

// NewGenericInoutMock returns a mock for genericInout
func NewGenericInoutMock[T any](t minimock.Tester) *GenericInoutMock[T] {
	m := &GenericInoutMock[T]{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.NameMock = mGenericInoutMockName[T]{mock: m}
	m.NameMock.callArgs = []*GenericInoutMockNameParams[T]{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mGenericInoutMockName[T any] struct {
	mock               *GenericInoutMock[T]
	defaultExpectation *GenericInoutMockNameExpectation[T]
	expectations       []*GenericInoutMockNameExpectation[T]

	callArgs []*GenericInoutMockNameParams[T]
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// GenericInoutMockNameExpectation specifies expectation struct of the genericInout.Name
type GenericInoutMockNameExpectation[T any] struct {
	mock      *GenericInoutMock[T]
	params    *GenericInoutMockNameParams[T]
	paramPtrs *GenericInoutMockNameParamPtrs[T]
	results   *GenericInoutMockNameResults[T]
	Counter   uint64
}

// GenericInoutMockNameParams contains parameters of the genericInout.Name
type GenericInoutMockNameParams[T any] struct {
	t1 T
}

// GenericInoutMockNameParamPtrs contains pointers to parameters of the genericInout.Name
type GenericInoutMockNameParamPtrs[T any] struct {
	t1 *T
}

// GenericInoutMockNameResults contains results of the genericInout.Name
type GenericInoutMockNameResults[T any] struct {
	t2 T
}

// Expect sets up expected params for genericInout.Name
func (mmName *mGenericInoutMockName[T]) Expect(t1 T) *mGenericInoutMockName[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericInoutMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericInoutMockNameExpectation[T]{}
	}

	if mmName.defaultExpectation.paramPtrs != nil {
		mmName.mock.t.Fatalf("GenericInoutMock.Name mock is already set by ExpectParams functions")
	}

	mmName.defaultExpectation.params = &GenericInoutMockNameParams[T]{t1}
	for _, e := range mmName.expectations {
		if minimock.Equal(e.params, mmName.defaultExpectation.params) {
			mmName.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmName.defaultExpectation.params)
		}
	}

	return mmName
}

// ExpectT1Param1 sets up expected param t1 for genericInout.Name
func (mmName *mGenericInoutMockName[T]) ExpectT1Param1(t1 T) *mGenericInoutMockName[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericInoutMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericInoutMockNameExpectation[T]{}
	}

	if mmName.defaultExpectation.params != nil {
		mmName.mock.t.Fatalf("GenericInoutMock.Name mock is already set by Expect")
	}

	if mmName.defaultExpectation.paramPtrs == nil {
		mmName.defaultExpectation.paramPtrs = &GenericInoutMockNameParamPtrs[T]{}
	}
	mmName.defaultExpectation.paramPtrs.t1 = &t1

	return mmName
}

// Inspect accepts an inspector function that has same arguments as the genericInout.Name
func (mmName *mGenericInoutMockName[T]) Inspect(f func(t1 T)) *mGenericInoutMockName[T] {
	if mmName.mock.inspectFuncName != nil {
		mmName.mock.t.Fatalf("Inspect function is already set for GenericInoutMock.Name")
	}

	mmName.mock.inspectFuncName = f

	return mmName
}

// Return sets up results that will be returned by genericInout.Name
func (mmName *mGenericInoutMockName[T]) Return(t2 T) *GenericInoutMock[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericInoutMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericInoutMockNameExpectation[T]{mock: mmName.mock}
	}
	mmName.defaultExpectation.results = &GenericInoutMockNameResults[T]{t2}
	return mmName.mock
}

// Set uses given function f to mock the genericInout.Name method
func (mmName *mGenericInoutMockName[T]) Set(f func(t1 T) (t2 T)) *GenericInoutMock[T] {
	if mmName.defaultExpectation != nil {
		mmName.mock.t.Fatalf("Default expectation is already set for the genericInout.Name method")
	}

	if len(mmName.expectations) > 0 {
		mmName.mock.t.Fatalf("Some expectations are already set for the genericInout.Name method")
	}

	mmName.mock.funcName = f
	return mmName.mock
}

// When sets expectation for the genericInout.Name which will trigger the result defined by the following
// Then helper
func (mmName *mGenericInoutMockName[T]) When(t1 T) *GenericInoutMockNameExpectation[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericInoutMock.Name mock is already set by Set")
	}

	expectation := &GenericInoutMockNameExpectation[T]{
		mock:   mmName.mock,
		params: &GenericInoutMockNameParams[T]{t1},
	}
	mmName.expectations = append(mmName.expectations, expectation)
	return expectation
}

// Then sets up genericInout.Name return parameters for the expectation previously defined by the When method
func (e *GenericInoutMockNameExpectation[T]) Then(t2 T) *GenericInoutMock[T] {
	e.results = &GenericInoutMockNameResults[T]{t2}
	return e.mock
}

func (mmName *mGenericInoutMockName[T]) Times(n uint64) *mGenericInoutMockName[T] {
	if n == 0 {
		mmName.mock.t.Fatalf("Times of GenericInoutMock.Name mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmName.expectedInvocations, n)
	return mmName
}

func (mmName *mGenericInoutMockName[T]) invocationsDone() bool {
	if len(mmName.expectations) == 0 && mmName.defaultExpectation == nil && mmName.mock.funcName == nil {
		// does not need to check invocations if no expectations, defaultExpectation or funcName set
		return true
	}

	// if expectations were set we check total invocations
	// if default expectation was set then invocations count should be greater than zero
	// if func was set then invocations count should be greater than zero
	totalInvocations := mm_atomic.LoadUint64(&mmName.mock.afterNameCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmName.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Name implements genericInout
func (mmName *GenericInoutMock[T]) Name(t1 T) (t2 T) {
	mm_atomic.AddUint64(&mmName.beforeNameCounter, 1)
	defer mm_atomic.AddUint64(&mmName.afterNameCounter, 1)

	if mmName.inspectFuncName != nil {
		mmName.inspectFuncName(t1)
	}

	mm_params := GenericInoutMockNameParams[T]{t1}

	// Record call args
	mmName.NameMock.mutex.Lock()
	mmName.NameMock.callArgs = append(mmName.NameMock.callArgs, &mm_params)
	mmName.NameMock.mutex.Unlock()

	for _, e := range mmName.NameMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.t2
		}
	}

	if mmName.NameMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmName.NameMock.defaultExpectation.Counter, 1)
		mm_want := mmName.NameMock.defaultExpectation.params
		mm_want_ptrs := mmName.NameMock.defaultExpectation.paramPtrs

		mm_got := GenericInoutMockNameParams[T]{t1}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.t1 != nil && !minimock.Equal(*mm_want_ptrs.t1, mm_got.t1) {
				mmName.t.Errorf("GenericInoutMock.Name got unexpected parameter t1, want: %#v, got: %#v%s\n", *mm_want_ptrs.t1, mm_got.t1, minimock.Diff(*mm_want_ptrs.t1, mm_got.t1))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmName.t.Errorf("GenericInoutMock.Name got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmName.NameMock.defaultExpectation.results
		if mm_results == nil {
			mmName.t.Fatal("No results are set for the GenericInoutMock.Name")
		}
		return (*mm_results).t2
	}
	if mmName.funcName != nil {
		return mmName.funcName(t1)
	}
	mmName.t.Fatalf("Unexpected call to GenericInoutMock.Name. %v", t1)
	return
}

// NameAfterCounter returns a count of finished GenericInoutMock.Name invocations
func (mmName *GenericInoutMock[T]) NameAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.afterNameCounter)
}

// NameBeforeCounter returns a count of GenericInoutMock.Name invocations
func (mmName *GenericInoutMock[T]) NameBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.beforeNameCounter)
}

// Calls returns a list of arguments used in each call to GenericInoutMock.Name.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmName *mGenericInoutMockName[T]) Calls() []*GenericInoutMockNameParams[T] {
	mmName.mutex.RLock()

	argCopy := make([]*GenericInoutMockNameParams[T], len(mmName.callArgs))
	copy(argCopy, mmName.callArgs)

	mmName.mutex.RUnlock()

	return argCopy
}

// MinimockNameDone returns true if the count of the Name invocations corresponds
// the number of defined expectations
func (m *GenericInoutMock[T]) MinimockNameDone() bool {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.NameMock.invocationsDone()
}

// MinimockNameInspect logs each unmet expectation
func (m *GenericInoutMock[T]) MinimockNameInspect() {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to GenericInoutMock.Name with params: %#v", *e.params)
		}
	}

	afterNameCounter := mm_atomic.LoadUint64(&m.afterNameCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.NameMock.defaultExpectation != nil && afterNameCounter < 1 {
		if m.NameMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to GenericInoutMock.Name")
		} else {
			m.t.Errorf("Expected call to GenericInoutMock.Name with params: %#v", *m.NameMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcName != nil && afterNameCounter < 1 {
		m.t.Error("Expected call to GenericInoutMock.Name")
	}

	if !m.NameMock.invocationsDone() && afterNameCounter > 0 {
		m.t.Errorf("Expected %d calls to GenericInoutMock.Name but found %d calls",
			mm_atomic.LoadUint64(&m.NameMock.expectedInvocations), afterNameCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *GenericInoutMock[T]) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockNameInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *GenericInoutMock[T]) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *GenericInoutMock[T]) minimockDone() bool {
	done := true
	return done &&
		m.MinimockNameDone()
}
