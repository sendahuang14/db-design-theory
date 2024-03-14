package main

import (
	"encoding/json"
	"os"

	mapset "github.com/deckarep/golang-set/v2"
)

type Attrs mapset.Set[string]

type Relation struct {
	attr Attrs
	F    []FD
}

type FD struct {
	src Attrs
	des Attrs
}

func parseJson(filename string) Data {
	byteResult, _ := os.ReadFile(filename)

	var data Data

	json.Unmarshal(byteResult, &data)

	return data
}

func getRelation(data Data) Relation {
	r := Relation{
		attr: mapset.NewSet[string](),
		F:    nil,
	}

	r.attr.Append(data.Attr...)

	for _, fd := range data.F {
		var src Attrs = mapset.NewSet[string]()
		src.Append(fd.Src...)

		var des Attrs = mapset.NewSet[string]()
		des.Append(fd.Des...)

		r.F = append(r.F, FD{src: src, des: des})
	}

	return r
}

func (r Relation) getClosure(attrs Attrs) Attrs {
	var c Attrs = mapset.NewSet[string]()
	c.Append(attrs.ToSlice()...)

	for {
		copy := c

		for _, fd := range r.F {
			if fd.src.IsSubset(c) {
				c = c.Union(fd.des)
			}
		}

		if copy.Equal(c) {
			break
		}
	}

	return c
}

func allSubsets(attrs Attrs) mapset.Set[Attrs] {
	s := mapset.NewSet[Attrs]()
	n := attrs.Cardinality()
	attrSlice := attrs.ToSlice()

	for i := 0; i < (1 << n); i++ {
		var subset Attrs = mapset.NewSet[string]()

		for j, a := range attrSlice {
			if (i & (1 << j)) != 0 {
				subset.Add(a)
			}
		}

		if !s.Contains(subset) {
			s.Add(subset)
		}
	}

	return s
}

func (r Relation) findAllSuperKeys() mapset.Set[Attrs] {
	subsets := allSubsets(r.attr)

	var emptySet Attrs = mapset.NewSet[string]()
	subsets.Remove(emptySet)

	superKeys := mapset.NewSet[Attrs]()

	for elem := range subsets.Iter() {
		if r.getClosure(elem).Equal(r.attr) {
			superKeys.Add(elem)
		}
	}

	return superKeys
}

func (r Relation) findAllKeys() mapset.Set[Attrs] {
	superKeys := r.findAllSuperKeys()

	keys := mapset.NewSet[Attrs]()

	for elem := range superKeys.Iter() {
		isKey := true
		for other := range superKeys.Iter() {
			if elem.IsProperSuperset(other) {
				isKey = false
				break
			}
		}

		if isKey {
			keys.Add(elem)
		}
	}

	return keys
}
