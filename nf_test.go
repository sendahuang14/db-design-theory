package main

import "testing"

func Test_isBCNF(t *testing.T) {
	r1 := data1()
	if r1.isBCNF() { // r1 is not BCNF
		t.Error("data1 is not BCNF")
	}

	r2 := data2()
	if r2.isBCNF() { // r1 is not BCNF
		t.Error("data2 is not BCNF")
	}
}

func Test_is3NF(t *testing.T) {
	r1 := data1()
	if r1.is3NF() { // r1 is not BCNF
		t.Error("data1 is not 3NF")
	}

	r2 := data2()
	if r2.is3NF() { // r1 is not BCNF
		t.Error("data2 is not 3NF")
	}
}
