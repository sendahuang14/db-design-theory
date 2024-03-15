package main

func (r Relation) isBCNF() bool {
	sks := r.findSuperKeys()

	for _, fd := range r.F {
		isSrcSK := false
		for sk := range sks.Iter() {
			if fd.src.Equal(sk) {
				isSrcSK = true
				break
			}
		}

		if !isSrcSK {
			return false
		}
	}

	return true
}

func (r Relation) is3NF() bool {
	sks := r.findSuperKeys()
	ps := r.findPrimes()

	for _, fd := range r.F {
		isSrcSK := false
		isPrime := true

		for _, d := range fd.des.ToSlice() {
			if !ps.Contains(d) {
				isPrime = false
				break
			}
		}

		for sk := range sks.Iter() {
			if fd.src.Equal(sk) {
				isSrcSK = true
				break
			}
		}

		if !isSrcSK && !isPrime {
			return false
		}
	}

	return true
}
