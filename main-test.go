package main

import "testing"

func Testcalculate(t *testing.T){
	if func1(2,4)!=6 {
       t.Error("Expected value is 6")
	}
}