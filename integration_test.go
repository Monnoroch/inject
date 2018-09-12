package inject

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type IntegrationTests struct {
	suite.Suite
}

func (self *IntegrationTests) TestEmptyModule() {
	type emptyModule struct{}
	_, err := InjectorOf(emptyModule{})
	self.Nil(err)
}

type testModule struct{}

func (self testModule) ProvideValue() (int, Annotation1) {
	return testValue, Annotation1{}
}

func (self *IntegrationTests) TestSingleProvider() {
	injector, err := InjectorOf(testModule{})
	self.Require().Nil(err)
	value := injector.MustGet(new(int), Annotation1{}).(int)
	self.Equal(testValue, value)
}

type testModuleRecursiveBase struct{}

func (self testModuleRecursiveBase) ProvideValue1() (int, Annotation1) {
	return testValue, Annotation1{}
}

type testModuleRecursive struct{}

func (self testModuleRecursive) ProvideValue2() (int, Annotation2) {
	return testValue + 1, Annotation2{}
}

func (self testModuleRecursive) ProvideSum(
	value1 int, _ Annotation1,
	value2 int, _ Annotation2,
) (int, Annotation3) {
	return value1 + value2, Annotation3{}
}

func (self *IntegrationTests) TestRecursiveProviders() {
	injector, err := InjectorOf(
		testModuleRecursiveBase{},
		testModuleRecursive{},
	)
	self.Require().Nil(err)
	value := injector.MustGet(new(int), Annotation3{}).(int)
	self.Equal(testValue*2+1, value)
}

type testModuleNotCached struct {
	calls int
}

func (self *testModuleNotCached) ProvideValue() (int, Annotation1) {
	self.calls += 1
	return testValue, Annotation1{}
}

type testModuleCached struct {
	calls int
}

func (self *testModuleCached) ProvideCachedValue() (int, Annotation2) {
	self.calls += 1
	return testValue, Annotation2{}
}

func (self *IntegrationTests) TestCache() {
	notCachedModule := &testModuleNotCached{}
	cachedModule := &testModuleCached{}

	injector, err := InjectorOf(notCachedModule, cachedModule)
	self.Require().Nil(err)

	_ = injector.MustGet(new(int), Annotation1{}).(int)
	_ = injector.MustGet(new(int), Annotation1{}).(int)
	self.Equal(2, notCachedModule.calls)

	_ = injector.MustGet(new(int), Annotation2{}).(int)
	_ = injector.MustGet(new(int), Annotation2{}).(int)
	self.Equal(1, cachedModule.calls)
}

type testFlagModule struct {
	flag  bool
	calls int
}

func (self *testFlagModule) ProvideFlag() (bool, Annotation2) {
	return self.flag, Annotation2{}
}

func (self *testFlagModule) ProvideValue() (int, Annotation2) {
	self.calls += 1
	return testValue * 2, Annotation2{}
}

type testLazyModule struct{}

func (self testLazyModule) ProvideValueLazy(
	value1 func() int, _ Annotation1,
	value2 func() int, _ Annotation2,
	flag bool, _ Annotation2,
) (int, Annotation3) {
	if flag {
		return value2() + 1, Annotation3{}
	} else {
		return value1() + 1, Annotation3{}
	}
}

func (self *IntegrationTests) TestLazy() {
	module1 := &testModuleNotCached{}
	module2 := &testFlagModule{false, 0}
	injector, err := InjectorOf(module1, module2, testLazyModule{})
	self.Require().Nil(err)

	value := injector.MustGet(new(int), Annotation3{}).(int)
	self.Equal(testValue+1, value)
	self.Equal(1, module1.calls)
	self.Equal(0, module2.calls)

	module1 = &testModuleNotCached{}
	module2 = &testFlagModule{true, 0}
	injector, err = InjectorOf(module1, module2, testLazyModule{})
	self.Require().Nil(err)

	value = injector.MustGet(new(int), Annotation3{}).(int)
	self.Equal(testValue*2+1, value)
	self.Equal(0, module1.calls)
	self.Equal(1, module2.calls)
}

type testValuesModule struct{}

func (self testValuesModule) ProvideValue1() (int, Annotation1) {
	return testValue, Annotation1{}
}

func (self testValuesModule) ProvideValue2() (int, Annotation2) {
	return testValue + 1, Annotation2{}
}

type testInjectedAnnotation struct{}

type testDynamicAnnotationModule struct {
	annotation Annotation
}

func (self testDynamicAnnotationModule) ProvideDouble(
	value int, _ testInjectedAnnotation,
) (int64, testInjectedAnnotation) {
	return int64(value) * 2, testInjectedAnnotation{}
}

func (self testDynamicAnnotationModule) ProvideAnnotation() (Annotation, testInjectedAnnotation) {
	return self.annotation, testInjectedAnnotation{}
}

type testSumModule struct{}

func (self testSumModule) ProvideSum(
	value1 int64, _ Annotation1,
	value2 int64, _ Annotation2,
) (int64, Annotation3) {
	return value1 + value2, Annotation3{}
}

func (self *IntegrationTests) TestDynamicAnnotations() {
	injector, err := InjectorOf(
		testValuesModule{},
		testDynamicAnnotationModule{Annotation1{}},
		testDynamicAnnotationModule{Annotation2{}},
		testSumModule{},
	)
	self.Require().Nil(err)

	value1 := injector.MustGet(new(int64), Annotation1{}).(int64)
	value2 := injector.MustGet(new(int64), Annotation2{}).(int64)
	value3 := injector.MustGet(new(int64), Annotation3{}).(int64)
	self.Equal(int64(testValue*2), value1)
	self.Equal(int64((testValue+1)*2), value2)
	self.Equal(int64(value1+value2), value3)
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(IntegrationTests))
}