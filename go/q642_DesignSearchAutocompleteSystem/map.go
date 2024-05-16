package q642_DesignSearchAutocompleteSystem

import (
	"sort"
)

type AutocompleteSystem struct {
	m map[string]int
	h string
	q map[string]int
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	m := make(map[string]int)
	for i, s := range sentences {
		m[s] = times[i]
	}

	return AutocompleteSystem{
		m,
		"",
		m,
	}
}

func (a *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		a.m[string(a.h)]++
		a.h = ""
		a.q = a.m
		return []string{}
	}

	a.h += string(c)

	sli := []pair{}

	for s, t := range a.q {
		if s[:len(a.h)] == a.h {
			sli = append(sli, pair{
				s,
				t,
			})
		} else {
			delete(a.q, s)
		}
	}

	sort.Slice(sli, func(i, j int) bool {
		a := sli[i]
		b := sli[j]
		return a.t>b.t || (a.t==b.t && a.s<b.s)
	})

	ans := []string{}

	for i := 0; i < 3; i++ {
        if len(sli) < i+1 {
            break
        }
        
		ans = append(ans, sli[i].s)
	}

	return ans
}

type pair struct {
	s string
	t int
}
