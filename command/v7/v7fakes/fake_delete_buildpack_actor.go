// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	sync "sync"

	v7action "code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeDeleteBuildpackActor struct {
	DeleteBuildpackStub        func(string, string) (v7action.Warnings, error)
	deleteBuildpackMutex       sync.RWMutex
	deleteBuildpackArgsForCall []struct {
		arg1 string
		arg2 string
	}
	deleteBuildpackReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	deleteBuildpackReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeleteBuildpackActor) DeleteBuildpack(arg1 string, arg2 string) (v7action.Warnings, error) {
	fake.deleteBuildpackMutex.Lock()
	ret, specificReturn := fake.deleteBuildpackReturnsOnCall[len(fake.deleteBuildpackArgsForCall)]
	fake.deleteBuildpackArgsForCall = append(fake.deleteBuildpackArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DeleteBuildpack", []interface{}{arg1, arg2})
	fake.deleteBuildpackMutex.Unlock()
	if fake.DeleteBuildpackStub != nil {
		return fake.DeleteBuildpackStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteBuildpackReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDeleteBuildpackActor) DeleteBuildpackCallCount() int {
	fake.deleteBuildpackMutex.RLock()
	defer fake.deleteBuildpackMutex.RUnlock()
	return len(fake.deleteBuildpackArgsForCall)
}

func (fake *FakeDeleteBuildpackActor) DeleteBuildpackCalls(stub func(string, string) (v7action.Warnings, error)) {
	fake.deleteBuildpackMutex.Lock()
	defer fake.deleteBuildpackMutex.Unlock()
	fake.DeleteBuildpackStub = stub
}

func (fake *FakeDeleteBuildpackActor) DeleteBuildpackArgsForCall(i int) (string, string) {
	fake.deleteBuildpackMutex.RLock()
	defer fake.deleteBuildpackMutex.RUnlock()
	argsForCall := fake.deleteBuildpackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDeleteBuildpackActor) DeleteBuildpackReturns(result1 v7action.Warnings, result2 error) {
	fake.deleteBuildpackMutex.Lock()
	defer fake.deleteBuildpackMutex.Unlock()
	fake.DeleteBuildpackStub = nil
	fake.deleteBuildpackReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteBuildpackActor) DeleteBuildpackReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.deleteBuildpackMutex.Lock()
	defer fake.deleteBuildpackMutex.Unlock()
	fake.DeleteBuildpackStub = nil
	if fake.deleteBuildpackReturnsOnCall == nil {
		fake.deleteBuildpackReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.deleteBuildpackReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteBuildpackActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteBuildpackMutex.RLock()
	defer fake.deleteBuildpackMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeleteBuildpackActor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v7.DeleteBuildpackActor = new(FakeDeleteBuildpackActor)
