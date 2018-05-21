package main

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "Database" can be found in github.com/Ragnar-BY/mock_testing
*/
import (
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	testify_assert "github.com/stretchr/testify/assert"
)

//DatabaseMock implements github.com/Ragnar-BY/mock_testing.Database
type DatabaseMock struct {
	t minimock.Tester

	ReadFunc       func(p string) (r string, r1 error)
	ReadCounter    uint64
	ReadPreCounter uint64
	ReadMock       mDatabaseMockRead

	WriteFunc       func(p string, p1 string) (r error)
	WriteCounter    uint64
	WritePreCounter uint64
	WriteMock       mDatabaseMockWrite
}

//NewDatabaseMock returns a mock for github.com/Ragnar-BY/mock_testing.Database
func NewDatabaseMock(t minimock.Tester) *DatabaseMock {
	m := &DatabaseMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ReadMock = mDatabaseMockRead{mock: m}
	m.WriteMock = mDatabaseMockWrite{mock: m}

	return m
}

type mDatabaseMockRead struct {
	mock             *DatabaseMock
	mockExpectations *DatabaseMockReadParams
}

//DatabaseMockReadParams represents input parameters of the Database.Read
type DatabaseMockReadParams struct {
	p string
}

//Expect sets up expected params for the Database.Read
func (m *mDatabaseMockRead) Expect(p string) *mDatabaseMockRead {
	m.mockExpectations = &DatabaseMockReadParams{p}
	return m
}

//Return sets up a mock for Database.Read to return Return's arguments
func (m *mDatabaseMockRead) Return(r string, r1 error) *DatabaseMock {
	m.mock.ReadFunc = func(p string) (string, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Database.Read method
func (m *mDatabaseMockRead) Set(f func(p string) (r string, r1 error)) *DatabaseMock {
	m.mock.ReadFunc = f
	return m.mock
}

//Read implements github.com/Ragnar-BY/mock_testing.Database interface
func (m *DatabaseMock) Read(p string) (r string, r1 error) {
	atomic.AddUint64(&m.ReadPreCounter, 1)
	defer atomic.AddUint64(&m.ReadCounter, 1)

	if m.ReadMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ReadMock.mockExpectations, DatabaseMockReadParams{p},
			"Database.Read got unexpected parameters")

		if m.ReadFunc == nil {

			m.t.Fatal("No results are set for the DatabaseMock.Read")

			return
		}
	}

	if m.ReadFunc == nil {
		m.t.Fatal("Unexpected call to DatabaseMock.Read")
		return
	}

	return m.ReadFunc(p)
}

//ReadMinimockCounter returns a count of DatabaseMock.ReadFunc invocations
func (m *DatabaseMock) ReadMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ReadCounter)
}

//ReadMinimockPreCounter returns the value of DatabaseMock.Read invocations
func (m *DatabaseMock) ReadMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.ReadPreCounter)
}

type mDatabaseMockWrite struct {
	mock             *DatabaseMock
	mockExpectations *DatabaseMockWriteParams
}

//DatabaseMockWriteParams represents input parameters of the Database.Write
type DatabaseMockWriteParams struct {
	p  string
	p1 string
}

//Expect sets up expected params for the Database.Write
func (m *mDatabaseMockWrite) Expect(p string, p1 string) *mDatabaseMockWrite {
	m.mockExpectations = &DatabaseMockWriteParams{p, p1}
	return m
}

//Return sets up a mock for Database.Write to return Return's arguments
func (m *mDatabaseMockWrite) Return(r error) *DatabaseMock {
	m.mock.WriteFunc = func(p string, p1 string) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Database.Write method
func (m *mDatabaseMockWrite) Set(f func(p string, p1 string) (r error)) *DatabaseMock {
	m.mock.WriteFunc = f
	return m.mock
}

//Write implements github.com/Ragnar-BY/mock_testing.Database interface
func (m *DatabaseMock) Write(p string, p1 string) (r error) {
	atomic.AddUint64(&m.WritePreCounter, 1)
	defer atomic.AddUint64(&m.WriteCounter, 1)

	if m.WriteMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.WriteMock.mockExpectations, DatabaseMockWriteParams{p, p1},
			"Database.Write got unexpected parameters")

		if m.WriteFunc == nil {

			m.t.Fatal("No results are set for the DatabaseMock.Write")

			return
		}
	}

	if m.WriteFunc == nil {
		m.t.Fatal("Unexpected call to DatabaseMock.Write")
		return
	}

	return m.WriteFunc(p, p1)
}

//WriteMinimockCounter returns a count of DatabaseMock.WriteFunc invocations
func (m *DatabaseMock) WriteMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.WriteCounter)
}

//WriteMinimockPreCounter returns the value of DatabaseMock.Write invocations
func (m *DatabaseMock) WriteMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.WritePreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *DatabaseMock) ValidateCallCounters() {

	if m.ReadFunc != nil && atomic.LoadUint64(&m.ReadCounter) == 0 {
		m.t.Fatal("Expected call to DatabaseMock.Read")
	}

	if m.WriteFunc != nil && atomic.LoadUint64(&m.WriteCounter) == 0 {
		m.t.Fatal("Expected call to DatabaseMock.Write")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *DatabaseMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *DatabaseMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *DatabaseMock) MinimockFinish() {

	if m.ReadFunc != nil && atomic.LoadUint64(&m.ReadCounter) == 0 {
		m.t.Fatal("Expected call to DatabaseMock.Read")
	}

	if m.WriteFunc != nil && atomic.LoadUint64(&m.WriteCounter) == 0 {
		m.t.Fatal("Expected call to DatabaseMock.Write")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *DatabaseMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *DatabaseMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.ReadFunc == nil || atomic.LoadUint64(&m.ReadCounter) > 0)
		ok = ok && (m.WriteFunc == nil || atomic.LoadUint64(&m.WriteCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.ReadFunc != nil && atomic.LoadUint64(&m.ReadCounter) == 0 {
				m.t.Error("Expected call to DatabaseMock.Read")
			}

			if m.WriteFunc != nil && atomic.LoadUint64(&m.WriteCounter) == 0 {
				m.t.Error("Expected call to DatabaseMock.Write")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

//AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
//it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *DatabaseMock) AllMocksCalled() bool {

	if m.ReadFunc != nil && atomic.LoadUint64(&m.ReadCounter) == 0 {
		return false
	}

	if m.WriteFunc != nil && atomic.LoadUint64(&m.WriteCounter) == 0 {
		return false
	}

	return true
}
