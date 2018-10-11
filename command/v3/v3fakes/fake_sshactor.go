// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeSSHActor struct {
	GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexStub        func(appName string, spaceGUID string, processType string, processIndex uint) (v3action.SSHAuthentication, v3action.Warnings, error)
	getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex       sync.RWMutex
	getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall []struct {
		appName      string
		spaceGUID    string
		processType  string
		processIndex uint
	}
	getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturns struct {
		result1 v3action.SSHAuthentication
		result2 v3action.Warnings
		result3 error
	}
	getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall map[int]struct {
		result1 v3action.SSHAuthentication
		result2 v3action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSSHActor) GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndex(appName string, spaceGUID string, processType string, processIndex uint) (v3action.SSHAuthentication, v3action.Warnings, error) {
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.Lock()
	ret, specificReturn := fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall[len(fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall)]
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall = append(fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall, struct {
		appName      string
		spaceGUID    string
		processType  string
		processIndex uint
	}{appName, spaceGUID, processType, processIndex})
	fake.recordInvocation("GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndex", []interface{}{appName, spaceGUID, processType, processIndex})
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.Unlock()
	if fake.GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexStub != nil {
		return fake.GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexStub(appName, spaceGUID, processType, processIndex)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturns.result1, fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturns.result2, fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturns.result3
}

func (fake *FakeSSHActor) GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexCallCount() int {
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.RLock()
	defer fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.RUnlock()
	return len(fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall)
}

func (fake *FakeSSHActor) GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall(i int) (string, string, string, uint) {
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.RLock()
	defer fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.RUnlock()
	return fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall[i].appName, fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall[i].spaceGUID, fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall[i].processType, fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall[i].processIndex
}

func (fake *FakeSSHActor) GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturns(result1 v3action.SSHAuthentication, result2 v3action.Warnings, result3 error) {
	fake.GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexStub = nil
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturns = struct {
		result1 v3action.SSHAuthentication
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSSHActor) GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall(i int, result1 v3action.SSHAuthentication, result2 v3action.Warnings, result3 error) {
	fake.GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexStub = nil
	if fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall == nil {
		fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall = make(map[int]struct {
			result1 v3action.SSHAuthentication
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall[i] = struct {
		result1 v3action.SSHAuthentication
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSSHActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.RLock()
	defer fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSSHActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.SSHActor = new(FakeSSHActor)
