package main

import (
	"testing"
)
func TestStore(t *testing.T) {
	var monster = Monster{
		Name: "牛魔王",
		Age: 500,
		Skill: 200,
	}
	err := monster.Store(monster)
	if err != nil {
		t.Fatalf("Store执行结果错误 err= %v",err)
	}
	t.Logf("Store执行结果正确")
}
func TestReStore(t *testing.T) {
	var monster Monster
	err := monster.ReStore(monster)
	if err != nil {
		t.Fatalf("ReStore执行结果错误 err= %v",err)
	}
	t.Logf("ReStore执行结果正确")
}
