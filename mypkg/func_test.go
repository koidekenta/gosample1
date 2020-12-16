package mypkg

import "testing"

func TestAdd(t *testing.T){
  if Add(5, 5) != 10 {
		t.Fatal("failed test")
  }
}