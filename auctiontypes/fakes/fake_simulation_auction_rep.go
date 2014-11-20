// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/auction/auctiontypes"
)

type FakeSimulationCellRep struct {
	StateStub        func() (auctiontypes.CellState, error)
	stateMutex       sync.RWMutex
	stateArgsForCall []struct{}
	stateReturns     struct {
		result1 auctiontypes.CellState
		result2 error
	}
	PerformStub        func(auctiontypes.Work) (auctiontypes.Work, error)
	performMutex       sync.RWMutex
	performArgsForCall []struct {
		arg1 auctiontypes.Work
	}
	performReturns struct {
		result1 auctiontypes.Work
		result2 error
	}
	ResetStub        func() error
	resetMutex       sync.RWMutex
	resetArgsForCall []struct{}
	resetReturns     struct {
		result1 error
	}
}

func (fake *FakeSimulationCellRep) State() (auctiontypes.CellState, error) {
	fake.stateMutex.Lock()
	fake.stateArgsForCall = append(fake.stateArgsForCall, struct{}{})
	fake.stateMutex.Unlock()
	if fake.StateStub != nil {
		return fake.StateStub()
	} else {
		return fake.stateReturns.result1, fake.stateReturns.result2
	}
}

func (fake *FakeSimulationCellRep) StateCallCount() int {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	return len(fake.stateArgsForCall)
}

func (fake *FakeSimulationCellRep) StateReturns(result1 auctiontypes.CellState, result2 error) {
	fake.StateStub = nil
	fake.stateReturns = struct {
		result1 auctiontypes.CellState
		result2 error
	}{result1, result2}
}

func (fake *FakeSimulationCellRep) Perform(arg1 auctiontypes.Work) (auctiontypes.Work, error) {
	fake.performMutex.Lock()
	fake.performArgsForCall = append(fake.performArgsForCall, struct {
		arg1 auctiontypes.Work
	}{arg1})
	fake.performMutex.Unlock()
	if fake.PerformStub != nil {
		return fake.PerformStub(arg1)
	} else {
		return fake.performReturns.result1, fake.performReturns.result2
	}
}

func (fake *FakeSimulationCellRep) PerformCallCount() int {
	fake.performMutex.RLock()
	defer fake.performMutex.RUnlock()
	return len(fake.performArgsForCall)
}

func (fake *FakeSimulationCellRep) PerformArgsForCall(i int) auctiontypes.Work {
	fake.performMutex.RLock()
	defer fake.performMutex.RUnlock()
	return fake.performArgsForCall[i].arg1
}

func (fake *FakeSimulationCellRep) PerformReturns(result1 auctiontypes.Work, result2 error) {
	fake.PerformStub = nil
	fake.performReturns = struct {
		result1 auctiontypes.Work
		result2 error
	}{result1, result2}
}

func (fake *FakeSimulationCellRep) Reset() error {
	fake.resetMutex.Lock()
	fake.resetArgsForCall = append(fake.resetArgsForCall, struct{}{})
	fake.resetMutex.Unlock()
	if fake.ResetStub != nil {
		return fake.ResetStub()
	} else {
		return fake.resetReturns.result1
	}
}

func (fake *FakeSimulationCellRep) ResetCallCount() int {
	fake.resetMutex.RLock()
	defer fake.resetMutex.RUnlock()
	return len(fake.resetArgsForCall)
}

func (fake *FakeSimulationCellRep) ResetReturns(result1 error) {
	fake.ResetStub = nil
	fake.resetReturns = struct {
		result1 error
	}{result1}
}

var _ auctiontypes.SimulationCellRep = new(FakeSimulationCellRep)