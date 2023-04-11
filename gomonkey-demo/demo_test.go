package main

import (
	"gomonkeydemo/model"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

// gomonkey 用法示例
// 场景：打桩可导出函数的返回值
func TestApplyFuncReturn(t *testing.T) {
	a := assert.New(t)
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	info := "hello cpp"
	patches = patches.ApplyFuncReturn(model.ReadLeaf, info, nil)

	ret, err := model.ReadLeaf("")
	a.Nil(err)
	a.Equal(info, ret)
}

// 场景：打桩可导出函数的返回值，多次调用返回不同
func TestApplyFuncSeq(t *testing.T) {
	a := assert.New(t)
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	info1 := "hello cpp"
	info2 := "hello golang"
	info3 := "hello gomonkey"
	outputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{info1, nil}},
		{Values: gomonkey.Params{info2, nil}},
		{Values: gomonkey.Params{info3, nil}},
	}
	patches = patches.ApplyFuncSeq(model.ReadLeaf, outputs)

	output, err := model.ReadLeaf("")
	a.Nil(err)
	a.Equal(info1, output)
	output, err = model.ReadLeaf("")
	a.Nil(err)
	a.Equal(info2, output)
	output, err = model.ReadLeaf("")
	a.Nil(err)
	a.Equal(info3, output)
}

// 场景：替换可导出函数的实现，用于修改形参
func TestApplyFunc(t *testing.T) {
	a := assert.New(t)
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches = patches.ApplyFunc(model.ModifyMap,
		func(mp map[string]int, v int) {
			mp["a"] = 200
		})

	mp := map[string]int{}
	model.ModifyMap(mp, 2)
	a.Equal(200, mp["a"])
}

// 场景：替换函数变量的返回值
func TestApplyFuncVarReturn(t *testing.T) {
	a := assert.New(t)
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	info := "hello cpp"
	patches = patches.ApplyFuncVarReturn(&model.Marshal, []byte(info), nil)

	ret, err := model.Marshal("")
	a.Nil(err)
	a.Equal([]byte(info), ret)
}

// 场景：替换全局变量
func TestApplyGlobalVar(t *testing.T) {
	a := assert.New(t)
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches = patches.ApplyGlobalVar(&model.Num, 20)
	a.Equal(20, model.Num)
}

// 场景：替换私有方法
func TestApplyPrivateMethod(t *testing.T) {
	a := assert.New(t)
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	s := &model.PrivateMethodStruct{}
	patches = patches.ApplyPrivateMethod(s, "haveEaten",
		func(_ model.PrivateMethodStruct) bool {
			return true
		})
	ret := s.AreYouHungry()
	a.Equal("I am hungry", ret)
}

// 暂时不支持的场景：替换不可导出结构的私有方法
// 实测不可导出结构的公有方法已经可以替换
func TestApplyPrivateStructPrivateMethod(t *testing.T) {
	//a := assert.New(t)
	//patches := gomonkey.NewPatches()
	//defer patches.Reset()
	//
	//h := model.NewHorse()
	//patches = patches.ApplyMethodFunc(h, "runImpl",
	//	func() error {
	//		return fmt.Errorf("err horse1")
	//	})
	//err := h.Run()
	//a.NotNil(err)
}
