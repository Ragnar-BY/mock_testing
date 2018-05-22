// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package main

import (
	"sync"
)

var (
	lockDatabaseMockRead  sync.RWMutex
	lockDatabaseMockWrite sync.RWMutex
)

// DatabaseMock is a mock implementation of Database.
//
//     func TestSomethingThatUsesDatabase(t *testing.T) {
//
//         // make and configure a mocked Database
//         mockedDatabase := &DatabaseMock{
//             ReadFunc: func(key string) (string, error) {
// 	               panic("TODO: mock out the Read method")
//             },
//             WriteFunc: func(key string, value string) error {
// 	               panic("TODO: mock out the Write method")
//             },
//         }
//
//         // TODO: use mockedDatabase in code that requires Database
//         //       and then make assertions.
//
//     }
type DatabaseMock struct {
	// ReadFunc mocks the Read method.
	ReadFunc func(key string) (string, error)

	// WriteFunc mocks the Write method.
	WriteFunc func(key string, value string) error

	// calls tracks calls to the methods.
	calls struct {
		// Read holds details about calls to the Read method.
		Read []struct {
			// Key is the key argument value.
			Key string
		}
		// Write holds details about calls to the Write method.
		Write []struct {
			// Key is the key argument value.
			Key string
			// Value is the value argument value.
			Value string
		}
	}
}

// Read calls ReadFunc.
func (mock *DatabaseMock) Read(key string) (string, error) {
	if mock.ReadFunc == nil {
		panic("moq: DatabaseMock.ReadFunc is nil but Database.Read was just called")
	}
	callInfo := struct {
		Key string
	}{
		Key: key,
	}
	lockDatabaseMockRead.Lock()
	mock.calls.Read = append(mock.calls.Read, callInfo)
	lockDatabaseMockRead.Unlock()
	return mock.ReadFunc(key)
}

// ReadCalls gets all the calls that were made to Read.
// Check the length with:
//     len(mockedDatabase.ReadCalls())
func (mock *DatabaseMock) ReadCalls() []struct {
	Key string
} {
	var calls []struct {
		Key string
	}
	lockDatabaseMockRead.RLock()
	calls = mock.calls.Read
	lockDatabaseMockRead.RUnlock()
	return calls
}

// Write calls WriteFunc.
func (mock *DatabaseMock) Write(key string, value string) error {
	if mock.WriteFunc == nil {
		panic("moq: DatabaseMock.WriteFunc is nil but Database.Write was just called")
	}
	callInfo := struct {
		Key   string
		Value string
	}{
		Key:   key,
		Value: value,
	}
	lockDatabaseMockWrite.Lock()
	mock.calls.Write = append(mock.calls.Write, callInfo)
	lockDatabaseMockWrite.Unlock()
	return mock.WriteFunc(key, value)
}

// WriteCalls gets all the calls that were made to Write.
// Check the length with:
//     len(mockedDatabase.WriteCalls())
func (mock *DatabaseMock) WriteCalls() []struct {
	Key   string
	Value string
} {
	var calls []struct {
		Key   string
		Value string
	}
	lockDatabaseMockWrite.RLock()
	calls = mock.calls.Write
	lockDatabaseMockWrite.RUnlock()
	return calls
}