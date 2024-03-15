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
		src := initAttrs()
		src.Append(fd.Src...)

		des := initAttrs()
		des.Append(fd.Des...)

		r.F = append(r.F, FD{src: src, des: des})
	}

	return r
}

func initAttrs() Attrs {
	var attrs Attrs = mapset.NewSet[string]()
	return attrs
}

func (r Relation) getClosure(attrs Attrs) Attrs {
	c := initAttrs()
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
		subset := initAttrs()

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

	emptySet := initAttrs()
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
