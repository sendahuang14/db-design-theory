package main

import (
	"testing"
)

func data1() Relation {
	data1 := parseJson("data1.json")
	return getRelation(data1)
}

func data2() Relation {
	data2 := parseJson("data2.json")
	return getRelation(data2)
}

func Test_getClosure(t *testing.T) {
	r1 := data1()

	a := initAttrs()
	a.Add("A")
	cA := r1.getClosure(a)

	cACorrect := initAttrs()
	cACorrect.Append([]string{"A"}...)

	if !cA.Equal(cACorrect) {
		t.Error("getClosure has some bugs (cA)")
	}

	ab := initAttrs()
	ab.Append([]string{"A", "B"}...)
	cAB := r1.getClosure(ab)

	cABCorrect := initAttrs()
	cABCorrect.Append([]string{"A", "B", "E", "C"}...)

	if !cAB.Equal(cABCorrect) {
		t.Error("getClosure has some bugs (cAB)")
	}
}

func Test_allSubsets(t *testing.T) {
	ab := initAttrs()
	ab.Append([]string{"A", "B"}...)

	sAB := allSubsets(ab)

	if sAB.Cardinality() != 4 {
		t.Error("There should be 4 subsets for {A, B}")
	}

	abc := initAttrs()
	abc.Append([]string{"A", "B", "C"}...)

	sABC := allSubsets(abc)

	if sABC.Cardinality() != 8 {
		t.Error("There should be 8 subsets for {A, B, C}")
	}
}

func Test_findAllSuperKeys(t *testing.T) {
	r1 := data1()

	sk1 := r1.findAllSuperKeys()

	if sk1.Cardinality() != 8 {
		t.Error("There should be 8 super-keys in data1")
	}

	bd := initAttrs()
	bd.Append([]string{"B", "D"}...)

	for sk := range sk1.Iter() {
		if !sk.IsSuperset(bd) {
			t.Error("Every super-keys in data1 should include BD")
		}
	}
}

func Test_findAllKeys(t *testing.T) {
	r1 := data1()

	k1 := r1.findAllKeys()

	if k1.Cardinality() != 1 {
		t.Error("There should be 1 key in data1")
	}

	bd := initAttrs()
	bd.Append([]string{"B", "D"}...)

	for k := range k1.Iter() {
		if !k.Equal(bd) {
			t.Error("The key of data1 should be BD")
		}
	}
}
