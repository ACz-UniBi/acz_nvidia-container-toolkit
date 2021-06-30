// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package config

import (
	"sync"
)

// Ensure, that configUpdaterMock does implement configUpdater.
// If this is not the case, regenerate this file with moq.
var _ configUpdater = &configUpdaterMock{}

// configUpdaterMock is a mock implementation of configUpdater.
//
// 	func TestSomethingThatUsesconfigUpdater(t *testing.T) {
//
// 		// make and configure a mocked configUpdater
// 		mockedconfigUpdater := &configUpdaterMock{
// 			UpdateFunc: func(config *Config) error {
// 				panic("mock out the Update method")
// 			},
// 		}
//
// 		// use mockedconfigUpdater in code that requires configUpdater
// 		// and then make assertions.
//
// 	}
type configUpdaterMock struct {
	// UpdateFunc mocks the Update method.
	UpdateFunc func(config *Config) error

	// calls tracks calls to the methods.
	calls struct {
		// Update holds details about calls to the Update method.
		Update []struct {
			// Config is the config argument value.
			Config *Config
		}
	}
	lockUpdate sync.RWMutex
}

// Update calls UpdateFunc.
func (mock *configUpdaterMock) Update(config *Config) error {
	callInfo := struct {
		Config *Config
	}{
		Config: config,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	if mock.UpdateFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.UpdateFunc(config)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedconfigUpdater.UpdateCalls())
func (mock *configUpdaterMock) UpdateCalls() []struct {
	Config *Config
} {
	var calls []struct {
		Config *Config
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
