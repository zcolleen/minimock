package tests

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestContextAccepterMock_AnyContext(t *testing.T) {
	tester := NewTesterMock(t)

	var mockCalled bool
	tester.ErrorfMock.Set(func(s string, args ...interface{}) {
		assert.Equal(t, "ContextAccepterMock.AcceptContext got unexpected parameters, want: %#v, got: %#v%s\n", s)

		mockCalled = true
	}).CleanupMock.Return()

	mock := NewContextAccepterMock(tester).
		AcceptContextMock.Expect(context.Background()).Return()

	mock.AcceptContext(context.TODO())

	assert.True(t, mockCalled)
}

func TestContextAccepterMock_TodoContextMatchesAnycontext(t *testing.T) {
	tester := NewTesterMock(t)
	tester.CleanupMock.Return()

	mock := NewContextAccepterMock(tester).
		AcceptContextMock.Expect(minimock.AnyContext).Return()

	mock.AcceptContext(context.TODO())
}

func TestContextAccepterMock_WhenThenMatchAnycontext(t *testing.T) {
	tester := NewTesterMock(t)
	tester.CleanupMock.Return()

	mock := NewContextAccepterMock(tester).
		AcceptContextWithOtherArgsMock.When(minimock.AnyContext, 1).Then(42, nil)

	result, err := mock.AcceptContextWithOtherArgs(context.TODO(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 42, result)
}

func TestContextAccepterMock_DiffWithoutAnyContext(t *testing.T) {
	tester := NewTesterMock(t)
	tester.CleanupMock.Return()

	tester.ErrorfMock.
		Expect("ContextAccepterMock.AcceptContextWithOtherArgs got unexpected parameters, want: %#v, got: %#v%s\n",
			ContextAccepterMockAcceptContextWithOtherArgsParams{
				ctx: minimock.AnyContext,
				i1:  24,
			},
			ContextAccepterMockAcceptContextWithOtherArgsParams{
				ctx: context.Background(),
				i1:  123,
			},
			"\n\nDiff:\n--- Expected params\n+++ Actual params\n@@ -4,3 +4,3 @@\n  },\n- i1: (int) 24\n+ i1: (int) 123\n }\n").
		Return()

	mock := NewContextAccepterMock(tester).
		AcceptContextWithOtherArgsMock.Expect(minimock.AnyContext, 24).Return(1, nil)

	_, _ = mock.AcceptContextWithOtherArgs(context.Background(), 123)
}

func TestContextAccepterMock_DiffInStructArgWithoutAnyContext(t *testing.T) {
	tester := NewTesterMock(t)
	tester.CleanupMock.Return()

	tester.ErrorfMock.
		Expect("ContextAccepterMock.AcceptContextWithStructArgs got unexpected parameters, want: %#v, got: %#v%s\n",
			ContextAccepterMockAcceptContextWithStructArgsParams{
				ctx: minimock.AnyContext,
				s1: structArg{
					a: 124,
					b: "abcd",
				},
			},
			ContextAccepterMockAcceptContextWithStructArgsParams{
				ctx: context.Background(),
				s1: structArg{
					a: 123,
					b: "abcd",
				},
			},
			"\n\nDiff:\n--- Expected params\n+++ Actual params\n@@ -5,3 +5,3 @@\n  s1: (tests.structArg) {\n-  a: (int) 124,\n+  a: (int) 123,\n   b: (string) (len=4) \"abcd\"\n").
		Return()

	mock := NewContextAccepterMock(tester).
		AcceptContextWithStructArgsMock.Expect(minimock.AnyContext, structArg{
		a: 124,
		b: "abcd",
	}).
		Return(1, nil)

	_, _ = mock.AcceptContextWithStructArgs(context.Background(), structArg{
		a: 123,
		b: "abcd",
	})
}

func TestContextAccepterMock_TimesSuccess(t *testing.T) {
	tester := NewTesterMock(t)
	tester.CleanupMock.Return()

	mock := NewContextAccepterMock(tester).
		AcceptContextWithStructArgsMock.Times(2).Expect(minimock.AnyContext, structArg{
		a: 124,
		b: "abcd",
	}).
		Return(1, nil).
		AcceptContextMock.Return()

	_, _ = mock.AcceptContextWithStructArgs(context.Background(), structArg{
		a: 124,
		b: "abcd",
	})
	_, _ = mock.AcceptContextWithStructArgs(context.Background(), structArg{
		a: 124,
		b: "abcd",
	})

	mock.AcceptContext(context.TODO())

	// explicitly call MinimockFinish here to imitate call of t.Cleanup(m.MinimockFinish)
	// as we mocked Cleanup call
	mock.MinimockFinish()
}

func TestContextAccepterMock_TimesFailure(t *testing.T) {
	tester := NewTesterMock(t)
	tester.CleanupMock.Return().
		ErrorfMock.Expect("Expected %d calls to ContextAccepterMock.AcceptContextWithStructArgs but found %d calls", uint64(1), uint64(2)).
		Return()

	// Expected 1 calls to ContextAccepterMock.AcceptContextWithStructArgs but found 2 calls
	mock := NewContextAccepterMock(tester).
		AcceptContextWithStructArgsMock.Times(1).Expect(minimock.AnyContext, structArg{
		a: 124,
		b: "abcd",
	}).
		Return(1, nil).
		AcceptContextMock.
		Times(1).Return()

	_, _ = mock.AcceptContextWithStructArgs(context.Background(), structArg{
		a: 124,
		b: "abcd",
	})
	_, _ = mock.AcceptContextWithStructArgs(context.Background(), structArg{
		a: 124,
		b: "abcd",
	})

	mock.AcceptContext(context.TODO())

	// explicitly call MinimockFinish here to imitate call of t.Cleanup(m.MinimockFinish)
	// as we mocked Cleanup call
	mock.MinimockFinish()
}

func TestContextAccepterMock_TimesZero(t *testing.T) {
	tester := NewTesterMock(t)
	tester.CleanupMock.Return().
		FatalfMock.Expect("Times of ContextAccepterMock.AcceptContextWithStructArgs mock can not be zero").
		Return()

	_ = NewContextAccepterMock(tester).
		AcceptContextWithStructArgsMock.Times(0).
		Return(1, nil)
}

func TestContextAccepterMock_ExpectedCall(t *testing.T) {
	tester := NewTesterMock(t)
	tester.CleanupMock.Times(1).Return().
		ErrorMock.Expect("Expected call to ContextAccepterMock.AcceptContext").Times(1).
		Return()

	mock := NewContextAccepterMock(tester).AcceptContextMock.Return()

	// explicitly call MinimockFinish here to imitate call of t.Cleanup(m.MinimockFinish)
	// as we mocked Cleanup call
	mock.MinimockFinish()
}
