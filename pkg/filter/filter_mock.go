// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package filter

import (
	"sync"
)

// Ensure, that EnvLookupMock does implement EnvLookup.
// If this is not the case, regenerate this file with moq.
var _ EnvLookup = &EnvLookupMock{}

// EnvLookupMock is a mock implementation of EnvLookup.
//
// 	func TestSomethingThatUsesEnvLookup(t *testing.T) {
//
// 		// make and configure a mocked EnvLookup
// 		mockedEnvLookup := &EnvLookupMock{
// 			LookupEnvFunc: func(s string) (string, bool) {
// 				panic("mock out the LookupEnv method")
// 			},
// 		}
//
// 		// use mockedEnvLookup in code that requires EnvLookup
// 		// and then make assertions.
//
// 	}
type EnvLookupMock struct {
	// LookupEnvFunc mocks the LookupEnv method.
	LookupEnvFunc func(s string) (string, bool)

	// calls tracks calls to the methods.
	calls struct {
		// LookupEnv holds details about calls to the LookupEnv method.
		LookupEnv []struct {
			// S is the s argument value.
			S string
		}
	}
	lockLookupEnv sync.RWMutex
}

// LookupEnv calls LookupEnvFunc.
func (mock *EnvLookupMock) LookupEnv(s string) (string, bool) {
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockLookupEnv.Lock()
	mock.calls.LookupEnv = append(mock.calls.LookupEnv, callInfo)
	mock.lockLookupEnv.Unlock()
	if mock.LookupEnvFunc == nil {
		var (
			sOut string
			bOut bool
		)
		return sOut, bOut
	}
	return mock.LookupEnvFunc(s)
}

// LookupEnvCalls gets all the calls that were made to LookupEnv.
// Check the length with:
//     len(mockedEnvLookup.LookupEnvCalls())
func (mock *EnvLookupMock) LookupEnvCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockLookupEnv.RLock()
	calls = mock.calls.LookupEnv
	mock.lockLookupEnv.RUnlock()
	return calls
}
