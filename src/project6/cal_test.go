package main

import (
	"testing"
)
func TestAddUpper(t *testing.T) {
	res := addUpper(1,2)
	if res != 3 {
		t.Fatalf("addUpper执行结果错误 期望值为 %v 实际结果为%v",3,res)
	}
	t.Logf("addUpper执行结果正确")
}
func TestAddCount(t *testing.T) {
	res := addCount(10)
	if res != 55 {
		t.Fatalf("addCount执行结果错误 期望值为 %v 实际结果为%v",55,res)
	}
	t.Logf("addUpper执行结果正确")
}
