package test

import (
	"fmt"
	"testing"
)
func TestAddUpper(t *testing.T) {
	res := addUpper(10,2)
	fmt.Println(res)
	t.Logf("pass")
	fmt.Println(res)
}
func TestSub(t *testing.T) {
	res := sub(2,1)
	fmt.Println(res)
}