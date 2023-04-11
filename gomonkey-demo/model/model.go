package model

import "fmt"

var Num = 10

// ReadLeaf 可导出函数，有返回值
func ReadLeaf(url string) (string, error) {
	output := fmt.Sprintf("%s, %s!", "Hello", "World")
	return output, nil
}

// ModifyMap 可导出函数，修改形参
func ModifyMap(mp map[string]int, v int) {
	mp["a"] = v
}

// ModifyMapV2 可导出函数，修改形参
func ModifyMapV2(mp map[string]int, v int) {
	modifyMapV2(mp, v)
}

// modifyMapV2 不可导出函数，修改形参
func modifyMapV2(mp map[string]int, v int) {
	mp["a"] = v
}

var Marshal = func(v interface{}) ([]byte, error) {
	return nil, nil
}

type PrivateMethodStruct struct {
}

func (s *PrivateMethodStruct) ok() bool {
	return s != nil
}

func (s *PrivateMethodStruct) Happy() string {
	if s.ok() {
		return "happy"
	}
	return "unhappy"
}

func (s PrivateMethodStruct) haveEaten() bool {
	return false
}

func (s PrivateMethodStruct) AreYouHungry() string {
	if s.haveEaten() {
		return "I am full"
	}

	return "I am hungry"
}
