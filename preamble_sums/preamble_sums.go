package preamble_sums

type FoundSum struct {
	M int
	N int
	LowestI int
}

// ns: a list of integers
// pL: preamble length
// returns a list of numbers not preceded within preamble length by two distinct numbers that sum to it
func ValidatePreambleSums(ns []int, pL int) []int {
	sums := map[int][]FoundSum{}
	invalid := make([]int, 0)
	for i, n := range ns {
		//populate found sums
		for j := i - pL; j < i; j++ {
			if j < 0 { j = 0 }
			m := ns[j]
			if n == m { continue }
			sum := n + m
			fS := FoundSum{M: m, N: n, LowestI: j}
			_, exists := sums[sum]
			if !exists {
				sums[sum] = []FoundSum{fS}
			} else {
				sums[sum] = append(sums[sum], fS)
			}
		}
		//check if current num has been found
		if i < pL {
			continue
		}
		fSs, exists := sums[n]
		if !exists {
			invalid = append(invalid, i)
			continue
		}
		valid := false
		for _, s := range fSs {
			if s.LowestI >= i - pL {
				valid = true
				continue
			}
		}
		if !valid {
			invalid = append(invalid, i)
		}
	}
	return invalid
}
