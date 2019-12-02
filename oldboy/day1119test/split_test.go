package day1119test

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("abcd", "bc")         // 程序输出的结果
	want := []string{"a", "d"}         // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}
