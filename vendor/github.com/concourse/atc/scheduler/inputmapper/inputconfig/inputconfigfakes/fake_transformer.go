// Code generated by counterfeiter. DO NOT EDIT.
package inputconfigfakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db/algorithm"
	"github.com/concourse/atc/scheduler/inputmapper/inputconfig"
)

type FakeTransformer struct {
	TransformInputConfigsStub        func(db *algorithm.VersionsDB, jobName string, inputs []atc.JobInput) (algorithm.InputConfigs, error)
	transformInputConfigsMutex       sync.RWMutex
	transformInputConfigsArgsForCall []struct {
		db      *algorithm.VersionsDB
		jobName string
		inputs  []atc.JobInput
	}
	transformInputConfigsReturns struct {
		result1 algorithm.InputConfigs
		result2 error
	}
	transformInputConfigsReturnsOnCall map[int]struct {
		result1 algorithm.InputConfigs
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTransformer) TransformInputConfigs(db *algorithm.VersionsDB, jobName string, inputs []atc.JobInput) (algorithm.InputConfigs, error) {
	var inputsCopy []atc.JobInput
	if inputs != nil {
		inputsCopy = make([]atc.JobInput, len(inputs))
		copy(inputsCopy, inputs)
	}
	fake.transformInputConfigsMutex.Lock()
	ret, specificReturn := fake.transformInputConfigsReturnsOnCall[len(fake.transformInputConfigsArgsForCall)]
	fake.transformInputConfigsArgsForCall = append(fake.transformInputConfigsArgsForCall, struct {
		db      *algorithm.VersionsDB
		jobName string
		inputs  []atc.JobInput
	}{db, jobName, inputsCopy})
	fake.recordInvocation("TransformInputConfigs", []interface{}{db, jobName, inputsCopy})
	fake.transformInputConfigsMutex.Unlock()
	if fake.TransformInputConfigsStub != nil {
		return fake.TransformInputConfigsStub(db, jobName, inputs)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.transformInputConfigsReturns.result1, fake.transformInputConfigsReturns.result2
}

func (fake *FakeTransformer) TransformInputConfigsCallCount() int {
	fake.transformInputConfigsMutex.RLock()
	defer fake.transformInputConfigsMutex.RUnlock()
	return len(fake.transformInputConfigsArgsForCall)
}

func (fake *FakeTransformer) TransformInputConfigsArgsForCall(i int) (*algorithm.VersionsDB, string, []atc.JobInput) {
	fake.transformInputConfigsMutex.RLock()
	defer fake.transformInputConfigsMutex.RUnlock()
	return fake.transformInputConfigsArgsForCall[i].db, fake.transformInputConfigsArgsForCall[i].jobName, fake.transformInputConfigsArgsForCall[i].inputs
}

func (fake *FakeTransformer) TransformInputConfigsReturns(result1 algorithm.InputConfigs, result2 error) {
	fake.TransformInputConfigsStub = nil
	fake.transformInputConfigsReturns = struct {
		result1 algorithm.InputConfigs
		result2 error
	}{result1, result2}
}

func (fake *FakeTransformer) TransformInputConfigsReturnsOnCall(i int, result1 algorithm.InputConfigs, result2 error) {
	fake.TransformInputConfigsStub = nil
	if fake.transformInputConfigsReturnsOnCall == nil {
		fake.transformInputConfigsReturnsOnCall = make(map[int]struct {
			result1 algorithm.InputConfigs
			result2 error
		})
	}
	fake.transformInputConfigsReturnsOnCall[i] = struct {
		result1 algorithm.InputConfigs
		result2 error
	}{result1, result2}
}

func (fake *FakeTransformer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.transformInputConfigsMutex.RLock()
	defer fake.transformInputConfigsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTransformer) recordInvocation(key string, args []interface{}) {
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

var _ inputconfig.Transformer = new(FakeTransformer)
