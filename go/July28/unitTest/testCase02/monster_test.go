package testCase02

import (
	"fmt"
	"testing"
)

func TestMonster_Store(t *testing.T) {
	monster := Monster{
		Name:  "tom",
		Age:   11,
		Skill: "fuck",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("error")
	}
	t.Logf("success")
}

func TestMonster_ReStore(t *testing.T) {
	var monster Monster
	res := monster.ReStore()
	fmt.Println(res)
	if !res {
		t.Fatalf("error")
	}
	t.Logf("success")

}